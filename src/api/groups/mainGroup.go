package api

import (
	"myapp/src/api/handlers"
	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	e.GET("/login", handlers.Login)
	e.GET("/cats/:data", handlers.GetCats)
	e.POST("/cats", handlers.AddCat)
	e.POST("/dogs", handlers.AddDog)
	e.POST("/hamsters", handlers.AddHamster)
}
