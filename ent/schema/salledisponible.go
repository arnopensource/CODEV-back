package schema

import "entgo.io/ent"

// SalleDisponible holds the schema definition for the SalleDisponible entity.
type SalleDisponible struct {
	ent.Schema
}

// Fields of the SalleDisponible.
func (SalleDisponible) Fields() []ent.Field {
	return []ent.Field{
		field.Date("date_disponible"),
		field.Time("horaire"),
	}
}

// Edges of the SalleDisponible.
func (SalleDisponible) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `User`
		// and reference it to the "cars" edge (in User schema)
		// explicitly using the `Ref` method.
		edge.From("id", Salle.Type).
			Ref("id_salle").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),
	}
}
