// package handlers

// import (
// 	"fmt"
// 	"myapp/src/api/models"
// 	"myapp/src/db"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"github.com/labstack/echo/v4"
// )

// //UserHandler retreave /user
// func UserHandler(g *echo.Group) {
// 	g.POST("", createHandler)
// 	g.GET("", getAll)
// 	g.GET("/:id", get)
// 	g.PUT("/:id", update)
// 	g.DELETE("/:id", delete)
// }

// func createHandler(c echo.Context) (err error) {
// 	//TODO:共通化（db2回呼んでいる）
// 	db := db.InitDB()
// 	user := new(models.User)
// 	if err := c.Bind(user); err != nil {
// 		return err
// 	}
// 	db.Create(&user)
// 	return c.String(http.StatusOK, "Registed new user")
// }

// func getAll(c echo.Context) error {
// 	var users []models.User
// 	db := db.InitDB()
// 	result := db.Find(&users)
// 	fmt.Println(result)

// 	return c.JSON(http.StatusOK, result)
// }

// func get(c echo.Context) (err error) {
// 	db := db.InitDB()
// 	data := &models.User{}
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	if err != nil {
// 		return err
// 	}
// 	db.Where("id = ?", id).First(data)

// 	return c.JSON(http.StatusOK, data)
// }

// func update(c echo.Context) (err error) {
// 	db := db.InitDB()
// 	newData := models.User{}
// 	err = c.Bind(&newData)
// 	originData := &models.User{}
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	db.Where("id = ?", id).First(originData)
// 	fmt.Println("このログ出力したい", newData.ID)
// 	if newData.ID != 0 && newData.ID != originData.ID {
// 		return c.JSON(http.StatusConflict, "urlに含まれるIDと編集データIDが一致していません")
// 	}
// 	if (newData.CreatedAt != time.Time{}) && (newData.CreatedAt != originData.CreatedAt) {
// 		return c.JSON(http.StatusConflict, "作成時間の編集はできません")
// 	}
// 	if (newData.UpdatedAt != time.Time{}) && (newData.UpdatedAt != originData.UpdatedAt) {
// 		return c.JSON(http.StatusConflict, "更新時間の編集はできません")
// 	}
// 	if newData.Name != originData.Name {
// 		originData.Name = newData.Name
// 	}
// 	if newData.Email != originData.Email {
// 		originData.Email = newData.Email
// 	}
// 	db.Save(&originData)

// 	return c.JSON(http.StatusOK, newData)
// }

// func delete(c echo.Context) (err error) {
// 	db := db.InitDB()
// 	data := &models.User{}
// 	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
// 	db.Where("id = ?", id).First(data)
// 	db.Delete(&data)
// 	return c.JSON(http.StatusOK,data)
// }
