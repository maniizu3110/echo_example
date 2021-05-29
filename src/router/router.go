package router

import (
	api "myapp/src/api/groups"
	"myapp/src/api/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// create groups
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	// set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetAdminMiddlewares(adminGroup)
	middlewares.SetCookieMiddlewares(cookieGroup)
	middlewares.SetJwtMiddlewares(jwtGroup)

	// set main routes
	api.MainGroup(e)
	// set group routes
	api.AdminGroup(adminGroup)
	api.CookieGroup(cookieGroup)
	//jwt認証がないと通れないということは全てのurlがこれを通るということ
	api.JwtGroup(jwtGroup)

	return e
}
