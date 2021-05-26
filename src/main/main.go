package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", yallo)
	e.GET("user/:id", getQuery)
	e.Logger.Fatal(e.Start(":8080"))
}

func yallo(c echo.Context) error {
	return c.String(http.StatusOK, "yallo")
}
func getQuery(c echo.Context) error {
	// http://localhost:8080/user/3?name=leoで接続すると値を返す
	name := c.QueryParam("name")
	id := c.Param("id")
	fmt.Println(id)
	return c.JSON(http.StatusOK, map[string]string{
		"name": name,
	})
}
