package api

import (
	"myapp/src/api/handlers"
	"github.com/labstack/echo/v4"
)

func MainGroup(g *echo.Group) {
	g.Group("")
	g.POST("/user", handlers.CreateUser)
}
