package router

import (
	"myapp/src/api/handlers"
	"myapp/src/api/middlewares"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

//Run start main router
func Run(db *gorm.DB) *echo.Echo {
	e := echo.New()

	// middlewares.GcpConfig() //本番環境では必要
	middlewares.CorsMiddlewares(e)

	g := e.Group("/api/v1")

	middlewares.CheckAdminMiddlewares(g)

	handlers.UserHandler(g.Group("/user"))
	return e
}
