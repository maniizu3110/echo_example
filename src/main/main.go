package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name string `json:"name"`
}
type Dog struct {
	Name string `json:"name"`
}
type Cat struct {
	Name string `json:"name"`
}

func main() {
	e := echo.New()
	e.GET("/", yallo)
	e.POST("user/add", addUser)
	e.POST("dog/add", addDog)
	e.POST("cat/add", addCat)
	e.GET("/user/:data", getQuery)
	e.Logger.Fatal(e.Start(":8080"))
}

func yallo(c echo.Context) error {
	return c.String(http.StatusOK, "yallo")
}

func getQuery(c echo.Context) error {
	// http://localhost:8080/user/json?name=leoで接続すると値を返す
	name := c.QueryParam("name")
	dataType := c.Param("data")

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

//Native go1
func addUser(c echo.Context) error {
	user := User{}
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("failed loading the request body for add: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Printf("failed unmarsharing: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	log.Printf("this is your user:%#v", user)
	return c.String(http.StatusOK, "we got your user")

}

//Native go2
func addDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("failed processing addDog request: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	log.Printf("this is your user:%#v", dog)
	return c.String(http.StatusOK, "we got your dog")
}

//Using Echo
func addCat(c echo.Context) error {
	cat := Cat{}
	err := c.Bind(&cat)
	if err != nil {
		log.Printf("Failed processing cat request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	log.Printf("this is your cat:%#v", cat)
	return c.String(http.StatusOK, "we get your cat")
}
