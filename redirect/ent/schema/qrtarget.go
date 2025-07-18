package schema

import (
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// QrTarget holds the schema definition for the QrTarget entity.
type QrTarget struct {
	ent.Schema
}

// Fields of the QrTarget.
func (QrTarget) Fields() []ent.Field {
	return []ent.Field{
		// UUID primary key
		field.UUID("id", uuid.New()).
			Default(uuid.New).
			Unique().
			StructTag(`json:"id"`),

		// Foreign key for target_type
		field.Int("target_type_id").
			Positive().
			StructTag(`json:"target_type_id"`),

		// Custom URL, nullable
		field.String("target").
			Optional().
			Nillable().
			StructTag(`json:"target"`),
	}
}

// Edges of the QrTarget.
func (QrTarget) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("targettype", TargetType.Type).
			Ref("qr_targets").
			Unique().
			Required().
			Field("target_type_id"),
	}
}
