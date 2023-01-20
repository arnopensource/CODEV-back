package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Salle holds the schema definition for the Salle entity.
type Salle struct {
	ent.Schema
}

// Fields of the Salle.
func (Salle) Fields() []ent.Field {
	return []ent.Field{
		field.String("nom"),
		field.String("batiment"),
		field.String("etage"),
		field.String("num_salle"),
		field.Int("cap_max"),
	}
}

// Edges of the Salle.
func (Salle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profil_reservation", Profile.Type).
			Ref("salle_reservee").
			Through("reservations", Reservation.Type),
		edge.To("disponibilite", SalleDisponible.Type),
	}
}
