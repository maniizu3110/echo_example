package router

import (
	"myapp/src/api/handlers"
	"myapp/src/api/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/jinzhu/gorm"
)

//New start main router
func New(db *gorm.DB) *echo.Echo {
	e := echo.New()
	
	// middlewares.GcpConfig()　本番では必要だが一旦切っておく
	middlewares.CorsMiddlewares(e)

	g := e.Group("/api/v1")
	middlewares.CheckAdminMiddlewares(g)

	handlers.UserHandler(g.Group("/user"))
	return e
}
