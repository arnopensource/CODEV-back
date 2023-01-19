package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("start"),
		field.Time("end"),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return nil
}
