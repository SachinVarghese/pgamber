package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// IncomeBracket holds the schema definition for the IncomeBracket entity.
type IncomeBracket struct {
	ent.Schema
}

// Fields of the IncomeBracket.
func (IncomeBracket) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("class").Values("lte_50K", "gt_50K").Default("lte_50K"),
	}
}

// Edges of the IncomeBracket.
func (IncomeBracket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("person", Individual.Type).Ref("bracket").Unique(),
	}
}
