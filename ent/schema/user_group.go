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
type UserGroup struct {
    ent.Schema
}

// Fields of the UserGroup.
func (UserGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Annotations(entgql.OrderField("CREATED_AT"), entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

// Edges of the UserGroup.
func (UserGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group",Group.Type).Unique().Ref("userGroups"),
		edge.From("user",User.Type).Unique().Ref("userGroups"),
	}
}

// Annotations of the .UserGroup.
func (UserGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users_groups"},
		entgql.QueryField("usersGroups"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}