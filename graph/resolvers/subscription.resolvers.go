package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"chat/auth"
	"chat/ent"
	"chat/graph/generated"
	"context"
)

// MessageByGroup is the resolver for the messageByGroup field.
func (r *subscriptionResolver) MessageByGroup(ctx context.Context, groupdID int) (<-chan *ent.Message, error) {

	socketClient := ctx.Value(auth.ContextKey{Key: "socketClient"}).(string)

	println("Message resolver client: ", socketClient)

	cm := Chat{
		GroupId:        groupdID,
		MessageChannel: make(chan *ent.Message, 1),
	}
	r.ChatObserversMutext.Lock()
	r.ChatObservers[socketClient] = &cm
	r.ChatObserversMutext.Unlock()

	// remove observer from map when disconcected
	go func() {
		<-ctx.Done()
		r.ChatObserversMutext.Lock()
		delete(r.ChatObservers, socketClient)
		r.ChatObserversMutext.Unlock()
	}()

	return cm.MessageChannel, nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
