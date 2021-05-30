package handlers

import (
	"myapp/src/api/models"
	"myapp/src/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

//AddDog decode and create dog
func AddDog(c echo.Context) error {

	db := db.OpenDB()

	defer db.Close()

	dog := new(models.Dog)

	if err := c.Bind(dog); err != nil{
		return err
	}

	db.Create(&dog)

	return c.String(http.StatusOK, "we got your dog!")
}
