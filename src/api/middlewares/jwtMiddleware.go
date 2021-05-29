package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetJwtMiddlewares(g *echo.Group) {

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte("mySecret"),
		SigningMethod: "HS512",
	}))
}
