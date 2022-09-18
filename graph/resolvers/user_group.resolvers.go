package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"chat/ent"
	"context"
)

// CreateUserGroup is the resolver for the createUserGroup field.
func (r *mutationResolver) CreateUserGroup(ctx context.Context, input ent.CreateUserGroupInput) (*ent.UserGroup, error) {
	return ent.FromContext(ctx).UserGroup.Create().SetInput(input).Save(ctx)
}

// UpdateUserGroup is the resolver for the updateUserGroup field.
func (r *mutationResolver) UpdateUserGroup(ctx context.Context, id int, input ent.UpdateUserGroupInput) (*ent.UserGroup, error) {
	return ent.FromContext(ctx).UserGroup.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteUserGroup is the resolver for the deleteUserGroup field.
func (r *mutationResolver) DeleteUserGroup(ctx context.Context, id int) (*ent.UserGroup, error) {
	entity, _ := r.Client.UserGroup.Get(ctx, id)
	return entity, ent.FromContext(ctx).UserGroup.DeleteOneID(id).Exec(ctx)
}

// UserGroup is the resolver for the UserGroup field.
func (r *queryResolver) UserGroup(ctx context.Context, id int) (*ent.UserGroup, error) {
	return r.Client.UserGroup.Get(ctx, id)
}
