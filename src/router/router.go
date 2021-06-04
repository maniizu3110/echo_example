package router

import (
	api "myapp/src/api/groups"
	"myapp/src/api/middlewares"

	"github.com/labstack/echo/v4"
)

func Run() *echo.Echo {
	e := echo.New()

	g := e.Group("/api/v1")

	middlewares.CorsMiddlewares(e) //何もしてないやつ（グループの練習）
	middlewares.CheckAdminMiddlewares(g)

	api.MainGroup(e)

	return e
}
