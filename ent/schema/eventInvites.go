package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type EventInvite struct {
	ent.Schema
}

func (EventInvite) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("event_id", "profile_id"),
	}
}

func (EventInvite) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("profile_id", uuid.New()).StructTag(`json:"profileId"`),
		field.Int("event_id").StructTag(`json:"eventId"`),
		field.Time("since").Default(time.Now),
	}
}

func (EventInvite) Edges() []ent.Edge {
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
