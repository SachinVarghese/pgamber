package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Individual holds the schema definition for the Individual entity.
type Individual struct {
	ent.Schema
}

// Fields of the Individual.
func (Individual) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
	}
}

// Edges of the Individual.
func (Individual) Edges() []ent.Edge {
	return nil
}
