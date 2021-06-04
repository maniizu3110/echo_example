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
	g.POST("",createHandler)
	g.GET("",getAll)
}


func createHandler(c echo.Context)(err error){
	fmt.Println("なんとかここまで来て欲しい")
	// user := new(models.User)
	// if err := c.Bind(user); err != nil {
	// 	return err
	// }
	// //ここでエラー出てるっぽい。
	// fmt.Println(c.Get("Tx").(*gorm.DB),"出力して欲しい")
	// db := c.Get("Tx").(*gorm.DB)
	// db.Create(&user)

	return c.String(http.StatusOK, "Registed new user")
}

func getAll(c echo.Context) error {
	//ここ
	db := c.Get("Tx").(*gorm.DB)
	result := db.Find(&models.User{})
	fmt.Println(result)

	return c.String(http.StatusOK, "get Users")
}