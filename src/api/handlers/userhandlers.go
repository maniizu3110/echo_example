package handlers

import (
	// "myapp/src/api/firebase/snippets"
	"fmt"
	"myapp/src/api/models"
	"myapp/src/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	fmt.Println("ログ表示されないんかな")
	db := db.InitDB()

	defer db.Close()

	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return err
	}
	//この操作まとめたい
	db.AutoMigrate(&models.User{})
	db.Create(&user)

	return c.String(http.StatusOK, "Registed new user")
}
