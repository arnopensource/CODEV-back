package schema

import "entgo.io/ent"

// Salle holds the schema definition for the Salle entity.
type Salle struct {
	ent.Schema
}

// Fields of the Salle.
func (Salle) Fields() []ent.Field {
	return []ent.Field{
		field.Integer("id"),
		field.String("nom"),
		field.String("batiment"),
		field.String("etage"),
		field.String("num_salle"),
		field.String("cap_max"),
	}
}

// Edges of the Salle.
func (Salle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("id_salle", SalleDisponible.type),
		edge.To("id_reservation", Reservation.type),

	}
}
