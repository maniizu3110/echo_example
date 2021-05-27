package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", yallo)
	e.POST("/add", addUser)
	e.GET("/user/:id", getQuery)
	e.Logger.Fatal(e.Start(":8080"))
}

func yallo(c echo.Context) error {
	return c.String(http.StatusOK, "yallo")
}
func getQuery(c echo.Context) error {
	// http://localhost:8080/user/3?name=leoで接続すると値を返す
	name := c.QueryParam("name")
	id := c.Param("id")
	dataType := c.Param("data")

	fmt.Println(id)
	if dataType == "strign" {
		return c.String(http.StatusOK, name)
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": name,
		})
	}
	return c.JSON(http.StatusBadRequest, map[string]string{"error": "internal server error"})
}

type User struct {
	Name string
}

func addUser(c echo.Context) error {
	user := User{}
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("failed loading the request body for add")
		return c.String(http.StatusInternalServerError, "")
	}
	err = json.Unmarshal(b, &user)
	log.Printf(err.Error())
	if err != nil {
		log.Printf("failed unmardshaling in add")
		return c.String(http.StatusInternalServerError, "")
	}
	log.Printf("this is your user:%v", user)
	return c.String(http.StatusOK, "we got your user")

}
