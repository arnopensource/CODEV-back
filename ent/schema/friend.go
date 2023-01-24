package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Friend struct {
	ent.Schema
}

func (Friend) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("profile_id", "friend_id"),
	}
}

func (Friend) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("profile_id", uuid.New()),
		field.UUID("friend_id", uuid.New()),
		field.Time("since"),
		field.Bool("accepted"),
	}
}

func (Friend) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", Profile.Type).
			Required().
			Unique().
			Field("profile_id"),
		edge.To("friend", Profile.Type).
			Required().
			Unique().
			Field("friend_id"),
	}
}
