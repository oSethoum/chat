package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"chat/ent"
	"context"
)

// CreateMessage is the resolver for the createMessage field.
func (r *mutationResolver) CreateMessage(ctx context.Context, input ent.CreateMessageInput) (*ent.Message, error) {
	message, err := ent.FromContext(ctx).Message.Create().SetInput(input).Save(ctx)

	// notify the observers in a another go routine
	go func() {
		for k := range r.ChatObservers {
			// notify only if the observer is in the group chat ( in this case the user )
			if r.ChatObservers[k].GroupId == *input.GroupID {
				r.ChatObservers[k].MessageChannel <- message
			}
		}
	}()

	return message, err
}

// UpdateMessage is the resolver for the updateMessage field.
func (r *mutationResolver) UpdateMessage(ctx context.Context, id int, input ent.UpdateMessageInput) (*ent.Message, error) {
	return ent.FromContext(ctx).Message.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteMessage is the resolver for the deleteMessage field.
func (r *mutationResolver) DeleteMessage(ctx context.Context, id int) (*ent.Message, error) {
	entity, _ := r.Client.Message.Get(ctx, id)
	return entity, ent.FromContext(ctx).Message.DeleteOneID(id).Exec(ctx)
}

// Message is the resolver for the Message field.
func (r *queryResolver) Message(ctx context.Context, id int) (*ent.Message, error) {
	return r.Client.Message.Get(ctx, id)
}
