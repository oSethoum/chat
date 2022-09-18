package resolvers

import (
	"chat/db"
	"chat/ent"
	"chat/graph/generated"
	"sync"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Chat struct {
	GroupId        int
	MessageChannel chan *ent.Message
}

type Resolver struct {
	Client              *ent.Client
	ChatObservers       map[string]*Chat
	ChatObserversMutext sync.Mutex
}

var schema *graphql.ExecutableSchema

func ExecutableSchema() graphql.ExecutableSchema {
	if schema == nil {
		schema = new(graphql.ExecutableSchema)
		*schema = generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{
			Client:              db.DB,
			ChatObservers:       map[string]*Chat{},
			ChatObserversMutext: sync.Mutex{},
		}})
	}

	return *schema
}
