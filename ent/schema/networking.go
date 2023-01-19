package schema

import "entgo.io/ent"

// Networking holds the schema definition for the Networking entity.
type Networking struct {
	ent.Schema
}

// Fields of the Networking.
func (Networking) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date"),
		field.String("localisation"),
	}
}

// Edges of the Networking.
func (Networking) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `User`
		// and reference it to the "cars" edge (in User schema)
		// explicitly using the `Ref` method.
		edge.From("id", Profile.Type).
			Ref("id_networking").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),
	}

}
