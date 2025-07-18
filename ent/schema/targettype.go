package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TargetType holds the schema definition for the TargetType entity.
type TargetType struct {
	ent.Schema
}

func (TargetType) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Positive().
			Immutable().
			StructTag(`json:"id"`),

		field.String("type").
			NotEmpty().
			Unique().
			StructTag(`json:"type"`),
	}
}

func (TargetType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("qr_targets", QrTarget.Type),
	}
}
