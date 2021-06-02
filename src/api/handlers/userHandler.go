package handlers

import (
	"myapp/src/api/firebase/snippets"
	"myapp/src/api/models"
	"myapp/src/db"
	"net/http"

	"github.com/labstack/echo/v4"
)


func UserHandler(g *echo.Group ){
	g.POST("",createUser)
}


func createUser(c echo.Context) error {

	///////////この処理まとめておきたい////////
	err := snippets.AdminUser(c)
	if err != nil {
		return err
	}
	db := db.OpenDB()
	defer db.Close()
	db.AutoMigrate(&models.User{})
	///////////////////////////////////////

	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return err
	}
	db.Create(&user)

	return c.String(http.StatusOK, "Registed new user")
}