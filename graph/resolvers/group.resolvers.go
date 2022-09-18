package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"chat/ent"
	"chat/graph/generated"
	"context"
)

// CreateGroup is the resolver for the createGroup field.
func (r *mutationResolver) CreateGroup(ctx context.Context, input ent.CreateGroupInput) (*ent.Group, error) {
	return ent.FromContext(ctx).Group.Create().SetInput(input).Save(ctx)
}

// UpdateGroup is the resolver for the updateGroup field.
func (r *mutationResolver) UpdateGroup(ctx context.Context, id int, input ent.UpdateGroupInput) (*ent.Group, error) {
	return ent.FromContext(ctx).Group.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteGroup is the resolver for the deleteGroup field.
func (r *mutationResolver) DeleteGroup(ctx context.Context, id int) (*ent.Group, error) {
	entity, _ := r.Client.Group.Get(ctx, id)
	return entity, ent.FromContext(ctx).Group.DeleteOneID(id).Exec(ctx)
}

// Group is the resolver for the Group field.
func (r *queryResolver) Group(ctx context.Context, id int) (*ent.Group, error) {
	return r.Client.Group.Get(ctx, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
