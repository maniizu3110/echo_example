package api

import (
	"myapp/src/api/handlers"
	"github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {
	e.POST("/user", handlers.CreateUser)
}
