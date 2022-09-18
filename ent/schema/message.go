package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/contrib/entgql"
	"time"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)
type Message struct {
    ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Time("sent_at").Default(time.Now).Annotations(entgql.OrderField("SENT_AT"), entgql.Skip(entgql.SkipMutationCreateInput)),
		field.String("message").Annotations(entgql.OrderField("MESSAGE")),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group",Group.Type).Unique().Ref("messages"),
		edge.From("user",User.Type).Unique().Ref("messages"),
	}
}

// Annotations of the .Message.
func (Message) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "messages"},
		entgql.QueryField("messages"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}