package auth

import (
	"chat/db"
	"chat/ent/user"
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ContextKey struct {
	Key string
}

type Respose struct {
	Errors interface{} `json:"errors"`
	Data   interface{} `json:"data"`
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	ctx := context.Background()
	var login LoginDTO
	c.Bind(&login)

	// check if the user is valid
	u, err := db.DB.User.Query().Where(user.UsernameEQ(login.Username)).First(ctx)

	if err != nil {
		return c.JSON(http.StatusBadRequest, Respose{Errors: map[string]string{"username": "User not found"}})
	}

	if u.Password != login.Password {
		return c.JSON(http.StatusBadRequest, Respose{Errors: map[string]string{"password": "Password don't match"}})

	}

	claims := Claims{
		UserID:   u.ID,
		Username: u.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Respose{Errors: "Token cannot be signed"})
	}

	return c.JSON(http.StatusOK, Respose{Data: map[string]interface{}{
		"token": t,
		"user":  u,
	}})
}

func ValuesToContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// extracting default context from echo
		ctx := c.Request().Context()

		if c.IsWebSocket() {
			// we usually get the socket client id in case of anonymuous chat ( doesn't require authentication )
			ctx = context.WithValue(ctx, ContextKey{Key: "socketClient"}, c.Request().Header.Get("Sec-Websocket-Key"))
		}

		c.SetRequest(c.Request().WithContext(ctx))

		next(c)
		return nil
	}
}

func Protected() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &Claims{},
		SigningKey: []byte("secret"),
	}

	return middleware.JWTWithConfig(config)
}
