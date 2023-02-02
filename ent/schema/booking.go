package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Booking struct {
	ent.Schema
}

func (Booking) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("profile_id", "room_id"),
	}
}

func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("profile_id", uuid.New()),
		field.Int("room_id").StructTag(`json:"roomId"`),
		field.Int("number_of_people").StructTag(`json:"numberOfPeople"`),
		field.Time("start"),
		field.Time("end"),
	}
}

func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).
			Required().
			Unique().
			Field("profile_id"),
		edge.To("room", Room.Type).
			Required().
			Unique().
			Field("room_id"),
	}
}
