package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Reservation holds the schema definition for the Reservation entity.
type Reservation struct {
	ent.Schema
}

func (Reservation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("profile_id", "salle_id"),
	}
}

// Fields of the Reservation.
func (Reservation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("profile_id", uuid.New()),
		field.Int("salle_id"),
		field.Int("quantity_students"),
		field.Time("horaire_res_initial"),
		field.Time("horaire_res_final"),
		field.Time("horaire_act"),
	}
}

// Edges of the Reservation.
func (Reservation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).
			Required().
			Unique().
			Field("profile_id"),
		edge.To("salle", Salle.Type).
			Required().
			Unique().
			Field("salle_id"),
	}
}
