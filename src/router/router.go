package router

import (
	"myapp/src/api/handlers"
	"myapp/src/api/middlewares"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func Router(database *gorm.DB) *echo.Echo {
	e := echo.New()
	middlewares.SetDBMiddleware(database)
	middlewares.CorsMiddlewares(e)
	
	//TODO:
	g := e.Group("/api/v1")
	middlewares.CheckAdminMiddlewares(g)


	handlers.UserHandler(g.Group("/user"))
	handlers.ScheduleHandler(g.Group("/schedule"))

	return e
}
