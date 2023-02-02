package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Event struct {
	ent.Schema
}

func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("activity"),
		field.Time("start"),
		field.Time("end"),
	}
}

func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profiles", Profile.Type).Through("members", Member.Type),
		edge.From("room", Room.Type).Ref("events").Unique().Required(),
		edge.To("invited", Profile.Type).Through("invites", EventInvite.Type),
	}
}
