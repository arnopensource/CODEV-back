package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Networking holds the schema definition for the Networking entity.
type Networking struct {
	ent.Schema
}

func (Networking) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("profile_id", "friend_id"),
	}
}

// Fields of the Networking.
func (Networking) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("profile_id", uuid.New()),
		field.UUID("friend_id", uuid.New()),
		field.Time("created_at"),
		field.Bool("accepted"),
	}
}

// Edges of the Networking.
func (Networking) Edges() []ent.Edge {
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
