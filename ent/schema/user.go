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
type User struct {
    ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Annotations(entgql.OrderField("CREATED_AT"), entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Time("updated_at").UpdateDefault(time.Now).Default(time.Now).Annotations(entgql.OrderField("UPDATED_AT"), entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("username").Unique().Annotations(entgql.OrderField("USERNAME")),
		field.String("password").Sensitive().Annotations(entgql.Skip(entgql.SkipType, entgql.SkipWhereInput)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("userGroups",UserGroup.Type),
		edge.To("messages",Message.Type),
	}
}

// Annotations of the .User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
		entgql.QueryField("users"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}