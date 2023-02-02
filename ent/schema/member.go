package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Member struct {
	ent.Schema
}

func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("event_id", "profile_id"),
	}
}

func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("profile_id", uuid.New()).StructTag(`json:"profileId"`),
		field.Int("event_id").StructTag(`json:"eventId"`),
		field.Bool("is_admin").StructTag(`json:"isAdmin"`).Default(false),
	}
}

func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).
			Unique().
			Required().
			Field("profile_id"),
		edge.To("event", Event.Type).
			Unique().
			Required().
			Field("event_id"),
	}
}
