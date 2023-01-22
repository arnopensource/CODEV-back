package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Profile struct {
	ent.Schema
}

func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("firstname"),
		field.String("lastname"),
		field.String("phone"),
	}
}

func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("friends", Profile.Type).Through("friends_data", Friend.Type),
		edge.To("bookings", Room.Type).Through("bookings_data", Booking.Type),
	}
}
