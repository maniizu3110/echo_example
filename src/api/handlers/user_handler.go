package handlers

import (
	"fmt"
	"myapp/src/api/models"
	"myapp/src/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//UserHandler retreave /user
func UserHandler(g *echo.Group) {
	g.POST("", createHandler)
	g.GET("", getAll)
	g.GET("/:id", get)
}

func createHandler(c echo.Context) (err error) {
	//TODO:共通化（db2回呼んでいる）
	db := db.InitDB()
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	db.Create(&user)
	return c.String(http.StatusOK, "Registed new user")
}

func getAll(c echo.Context) error {
	db := db.InitDB()
	result := db.Find(&models.User{})
	fmt.Println(result)

	return c.String(http.StatusOK, "get Users")
}

func get(c echo.Context) (err error ){
	db := db.InitDB()
	data:= &models.User{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	result := db.Where("id = ?",id).First(data)
	fmt.Println(result)

	return c.JSON(http.StatusOK,data)
}
