package api

import (
	"github.com/labstack/echo/v4"
	"myapp/src/api/handlers"
)

func UserGroup(g *echo.Group) {
	g.POST("/create", handlers.CreateUser)
}