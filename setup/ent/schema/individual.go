package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Individual holds the schema definition for the Individual entity.
type Individual struct {
	ent.Schema
}

// Fields of the Individual.
func (Individual) Fields() []ent.Field {
	return []ent.Field{
		field.Float("age").Positive(),
		field.Int("workclass").NonNegative(),
		field.Int("education").NonNegative(),
		field.Int("marital_status").NonNegative(),
		field.Int("occupation").NonNegative(),
		field.Int("relationship").NonNegative(),
		field.Int("race").NonNegative(),
		field.Int("sex").NonNegative(),
		field.Float("capital_gain"),
		field.Float("capital_loss"),
		field.Float("hours_per_week"),
		field.Int("country").NonNegative(),
	}
}

// Edges of the Individual.
func (Individual) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bracket", IncomeBracket.Type).Unique(),
	}
}
