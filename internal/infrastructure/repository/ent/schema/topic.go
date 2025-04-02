package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Topic holds the schema definition for the Topic entity.
type Topic struct {
	ent.Schema
}

// Fields of the Topic.
func (Topic) Fields() []ent.Field {
	return []ent.Field{
	
		field.String("name").Unique(),
		field.String("description").Optional(),
		field.String("arabic_name").Optional(),
		field.String("arabic_description").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),

	}
}

// Edges of the Topic.
func (Topic) Edges() []ent.Edge {
	return nil
}
