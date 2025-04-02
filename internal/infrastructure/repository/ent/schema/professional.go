package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Professional holds the schema definition for the Professional entity.
type Professional struct {
	ent.Schema
}

// Fields of the Professional.
func (Professional) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Optional(),
		
	}
}

// Edges of the Professional.
func (Professional) Edges() []ent.Edge {
	return nil
}
