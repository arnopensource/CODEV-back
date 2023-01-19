package schema

import "entgo.io/ent"

// Reservation holds the schema definition for the Reservation entity.
type Reservation struct {
	ent.Schema
}

// Fields of the Reservation.
func (Reservation) Fields() []ent.Field {
	return []ent.Field{
		field.Integer("quantity_students"),
		field.Time("horaire_res_initial"),
		field.Time("horaire_res_final"),
		field.Time("horaire_act"),
	}
}

// Edges of the Reservation.
func (Reservation) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `User`
		// and reference it to the "cars" edge (in User schema)
		// explicitly using the `Ref` method.
		edge.From("id", Profile.Type).
			Ref("id_reservation").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),
		edge.From("id", Salle.Type).
			Ref("id_salle").
			Unique(),
	}
}
