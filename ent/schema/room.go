package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Room struct {
	ent.Schema
}

func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("floor"),
		field.String("building"),
		field.Int("capacity"),
	}
}

func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bookings", Profile.Type).
			Ref("bookings").
			Through("bookings_data", Booking.Type),
		edge.To("availability", AvailableRoom.Type),
		edge.To("events", Event.Type),
	}
}
