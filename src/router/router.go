package router

import (
	api "myapp/src/api/groups"
	"myapp/src/api/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// create groups
	//api/v1をグループ化しておきたい
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	// set all middlewares
	middlewares.SetMainMiddlewares(e) //何もしてないやつ（グループの練習）
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
