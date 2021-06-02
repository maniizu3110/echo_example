package handlers

import (
	"myapp/src/api/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/jinzhu/gorm"
)


//UserHandler retreave /user
func UserHandler(g *echo.Group){
	g.POST("",createUser)	
}


func createUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	db := c.Get("Tx").(*gorm.DB)
	db.Create(&user)

	return c.String(http.StatusOK, "Registed new user")
}