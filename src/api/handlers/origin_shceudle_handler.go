package handlers

import (
	"fmt"
	"myapp/src/api/models"
	"myapp/src/db"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

//OriginScheduleHandler retreave /OriginSchedule
func OriginScheduleHandler(g *echo.Group) {
	g.POST("", createOriginScheduleHandler)
	g.GET("", getAllOriginSchedule)
	g.GET("/:id", getOriginSchedule)
	g.PUT("/:id", updateOriginSchedule)
	g.DELETE("/:id", deleteOriginSchedule)
}


func createOriginScheduleHandler(c echo.Context) (err error) {
	//TODO:共通化（db2回呼んでいる）
	db := db.InitDB()
	OriginSchedule := new(models.OriginSchedule)
	if err := c.Bind(OriginSchedule); err != nil {
		return err
	}
	db.Create(&OriginSchedule)
	return c.JSON(http.StatusOK, OriginSchedule)
}

func getAllOriginSchedule(c echo.Context) error {
	var OriginSchedules []models.OriginSchedule
	db := db.InitDB()
	result := db.Find(&OriginSchedules)
	fmt.Println(result)

	return c.JSON(http.StatusOK, result)
}

func getOriginSchedule(c echo.Context) (err error) {
	db := db.InitDB()
	data := &models.OriginSchedule{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return err
	}
	db.Where("id = ?", id).First(data)

	return c.JSON(http.StatusOK, data)
}

func updateOriginSchedule(c echo.Context) (err error) {
	db := db.InitDB()
	newData := models.OriginSchedule{}
	err = c.Bind(&newData)
	originData := &models.OriginSchedule{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	db.Where("id = ?", id).First(originData)
	if newData.ID != 0 && newData.ID != originData.ID {
		return c.JSON(http.StatusConflict, "urlに含まれるIDと編集データIDが一致していません")
	}
	if (newData.CreatedAt != time.Time{}) && (newData.CreatedAt != originData.CreatedAt) {
		return c.JSON(http.StatusConflict, "作成時間の編集はできません")
	}
	if (newData.UpdatedAt != time.Time{}) && (newData.UpdatedAt != originData.UpdatedAt) {
		return c.JSON(http.StatusConflict, "更新時間の編集はできません")
	}
	db.Save(&originData)

	return c.JSON(http.StatusOK, newData)
}

func deleteOriginSchedule(c echo.Context) (err error) {
	db := db.InitDB()
	data := &models.OriginSchedule{}
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	db.Where("id = ?", id).First(data)
	db.Delete(&data)
	return c.JSON(http.StatusOK,data)
}
