// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/amr-mahmoud/mrktly/internal/infrastructure/repository/ent/predicate"
	"github.com/amr-mahmoud/mrktly/internal/infrastructure/repository/ent/professional"
)

// ProfessionalDelete is the builder for deleting a Professional entity.
type ProfessionalDelete struct {
	config
	hooks    []Hook
	mutation *ProfessionalMutation
}

// Where appends a list predicates to the ProfessionalDelete builder.
func (pd *ProfessionalDelete) Where(ps ...predicate.Professional) *ProfessionalDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *ProfessionalDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *ProfessionalDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *ProfessionalDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(professional.Table, sqlgraph.NewFieldSpec(professional.FieldID, field.TypeInt))
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// ProfessionalDeleteOne is the builder for deleting a single Professional entity.
type ProfessionalDeleteOne struct {
	pd *ProfessionalDelete
}

// Where appends a list predicates to the ProfessionalDelete builder.
func (pdo *ProfessionalDeleteOne) Where(ps ...predicate.Professional) *ProfessionalDeleteOne {
	pdo.pd.mutation.Where(ps...)
	return pdo
}

// Exec executes the deletion query.
func (pdo *ProfessionalDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{professional.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *ProfessionalDeleteOne) ExecX(ctx context.Context) {
	if err := pdo.Exec(ctx); err != nil {
		panic(err)
	}
}
