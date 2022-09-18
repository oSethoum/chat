package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"chat/ent"
	"chat/graph/generated"
	"context"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.Client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.Client.Noders(ctx, ids)
}

// Groups is the resolver for the Groups field.
func (r *queryResolver) Groups(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GroupOrder, where *ent.GroupWhereInput) (*ent.GroupConnection, error) {
	return r.Client.Group.Query().Paginate(ctx, after, first, before, last, ent.WithGroupOrder(orderBy), ent.WithGroupFilter(where.Filter))
}

// Messages is the resolver for the Messages field.
func (r *queryResolver) Messages(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MessageOrder, where *ent.MessageWhereInput) (*ent.MessageConnection, error) {
	return r.Client.Message.Query().Paginate(ctx, after, first, before, last, ent.WithMessageOrder(orderBy), ent.WithMessageFilter(where.Filter))
}

// Users is the resolver for the Users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.Client.User.Query().Paginate(ctx, after, first, before, last, ent.WithUserOrder(orderBy), ent.WithUserFilter(where.Filter))
}

// UsersGroups is the resolver for the UsersGroups field.
func (r *queryResolver) UsersGroups(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserGroupOrder, where *ent.UserGroupWhereInput) (*ent.UserGroupConnection, error) {
	return r.Client.UserGroup.Query().Paginate(ctx, after, first, before, last, ent.WithUserGroupOrder(orderBy), ent.WithUserGroupFilter(where.Filter))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
