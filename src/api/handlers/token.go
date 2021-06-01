package handlers

import (
	"fmt"
	"myapp/src/firebaase-admin/snippets"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Token(c echo.Context) error {
	firebaseApp := snippets.InitializeAppDefault()
	idToken :="ozyA2M7V5zeZIXUoFY6esuygjKU2"
	fmt.Println(idToken)
	//渡してるtokenでエラー出てる
	token := snippets.VerifyIDToken(c.Request().Context(),firebaseApp,idToken)
	fmt.Println(token,"これが取れるらしい")
	return c.String(http.StatusOK, "とれたああああああああああああ!")
}
