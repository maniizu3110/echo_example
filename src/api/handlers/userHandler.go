package handlers

import (
	"fmt"
	"myapp/src/api/models"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

//UserHandler retreave /user
func UserHandler(g *echo.Group){
	g.POST("",create)
	g.GET("",getAll)
}


func create(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	db := c.Get("Tx").(*gorm.DB)
	db.Create(&user)

	return c.String(http.StatusOK, "Registed new user")
}
func getAll(c echo.Context) error {
	db := c.Get("Tx").(*gorm.DB)
	result := db.Find(&models.User{})
	fmt.Println(result)

	return c.String(http.StatusOK, "get Users")
}