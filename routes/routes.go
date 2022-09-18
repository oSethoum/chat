package routes

import (
	"chat/auth"
	"chat/db"
	"chat/handlers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.POST("/login", auth.Login)

	r := e.Group("/")
	// r.Use(handlers.Protected())
	r.Use(auth.ValuesToContext)
	r.GET("", handlers.PlaygroundHandler())
	r.Any("query", handlers.GraphqlHandlers(db.DB))

	// this one is used to split link for apollo client / urql
	//r.Any("subscriptions", handlers.GraphqlWsHandler(db.DB))
	r.GET("ws", handlers.PlaygroundWsHandler())
}
