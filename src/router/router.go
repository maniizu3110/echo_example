package router

import (
	"myapp/src/api/handlers"
	"myapp/src/api/middlewares"

	"github.com/labstack/echo/v4"
)

func Router() *echo.Echo {
	e := echo.New()
	middlewares.CorsMiddlewares(e)
	
	//TODO:
	g := e.Group("/api/v1")
	middlewares.CheckAdminMiddlewares(g)


	handlers.UserHandler(g.Group("/user"))

	return e
}
