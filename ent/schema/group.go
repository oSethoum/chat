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
type Group struct {
    ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.Time("updated_at").UpdateDefault(time.Now).Default(time.Now).Annotations(entgql.OrderField("UPDATED_AT"), entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Time("created_at").Default(time.Now).Annotations(entgql.OrderField("CREATED_AT"), entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("name").Annotations(entgql.OrderField("NAME")),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("userGroups",UserGroup.Type),
		edge.To("messages",Message.Type),
	}
}

// Annotations of the .Group.
func (Group) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "groups"},
		entgql.QueryField("groups"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}