package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type AvailableRoom struct {
	ent.Schema
}

func (AvailableRoom) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start"),
		field.Time("end"),
	}
}

func (AvailableRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rooms", Room.Type).Ref("availability").Unique(),
	}
}
