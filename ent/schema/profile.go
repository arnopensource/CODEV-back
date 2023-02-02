package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Profile struct {
	ent.Schema
}

func (Profile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{
			StructTag: `json:"-"`,
		},
	}
}

func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.String("firstname").Optional(),
		field.String("lastname").Optional(),
		field.String("phone").Optional(),
		field.String("email").Immutable(),
	}
}

func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("friends", Profile.Type).Through("friends_data", Friend.Type),
		edge.To("bookings", Room.Type).Through("bookings_data", Booking.Type),
	}
}
