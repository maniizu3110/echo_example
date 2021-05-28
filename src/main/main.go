package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(ServerHeader)
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")

	//正しいcookieを持っていれば/cookie/mainにアクセスできるが、持っていない時はcheckCookie内のunauthorizedが返される
	cookieGroup.Use(checkCookie)

	// g.Use(middleware.Logger()) create auto log
	adminGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${latency_human}` + "\n",
	}))
	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "leo" && password == "1234" {
			return true, nil
		}
		return false, nil
	}))

	adminGroup.GET("/main", mainAdmin)
	cookieGroup.GET("/main", mainCookie)
	e.GET("/", yallo)
	e.POST("/user/add", addUser)
	e.GET("/login", login)
	e.POST("/dog/add", addDog)
	e.POST("/cat/add", addCat)
	e.GET("/user/:data", getQuery)
	e.Logger.Fatal(e.Start(":8080"))
}

//grouping path
func mainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "you are on the secret cookie main page")
}
func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "you are on the secret admin main page")
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

//Native go1(high paformance)
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

//Native go2(high peformance)
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

//Using Echo(easy way)
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

//sqlに接続したと仮定してcookieを設定している
func login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")
	//check username and password against DB after hashing the password
	if username == "leo" && password == "1234" {
		cookie := &http.Cookie{}
		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)
		return c.String(http.StatusOK, "you were logged in")

	}
	return c.String(http.StatusUnauthorized, "your info was wrong")
}

/////////////////////////middlewares/////////////////////////////////
//creating custom header
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "BlueBot/1.0")
		c.Response().Header().Set("notrallyHeader", "test")
		return next(c)

	}
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			if strings.Contains(err.Error(),"named cookie not present"){
				return c.String(http.StatusUnauthorized,"you don't have any cookie")
			}
			log.Println(err)
			return err
		}
		//もし他に（セッションとか）認証情報があれば、承認するのでここに処理を入れることも多い
		if  cookie.Value == "some_string"{
			return next(c)
		}
		return c.String(http.StatusUnauthorized,"you don't have right cookie")

	}
}

