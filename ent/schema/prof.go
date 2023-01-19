package schema

import "entgo.io/ent"

// Prof holds the schema definition for the Prof entity.
type Prof struct {
	ent.Schema
}

// Fields of the Prof.
func (Prof) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("Disponible"),
		field.String("desc_act"),
		field.String("localisation"),
	}
}

// Edges of the Prof.
func (Prof) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `User`
		// and reference it to the "cars" edge (in User schema)
		// explicitly using the `Ref` method.
		edge.From("id", Profile.Type).
			Ref("id_prof").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),
	}
}
