package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SalleDisponible holds the schema definition for the SalleDisponible entity.
type SalleDisponible struct {
	ent.Schema
}

// Fields of the SalleDisponible.
func (SalleDisponible) Fields() []ent.Field {
	return []ent.Field{
		field.String("id_salle"),
		field.Time("start"),
		field.Time("end"),
	}
}

// Edges of the SalleDisponible.
func (SalleDisponible) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("salle", Salle.Type).Ref("disponibilite").Unique(),
	}
}
