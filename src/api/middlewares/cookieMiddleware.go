package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

//正しいcookieを持っていれば/cookie/mainにアクセスできるが、持っていない時はcheckCookie内のunauthorizedが返される
//ここでは仮でkey:"sessionID",value:"some_string"
func SetCookieMiddlewares(g *echo.Group){
	g.Use(checkCookie)
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "you dont have any cookie")
			}

			log.Println(err)
			return err
		}

		if cookie.Value == "some_string" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "you dont have the right cookie, cookie")
	}
}
