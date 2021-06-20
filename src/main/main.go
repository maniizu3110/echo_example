package main

import (
	"myapp/src/api/api"
	"myapp/src/db"
	"myapp/src/router"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	api.GetDayCalendar()
	var database *gorm.DB
	database = db.InitDB()
	defer database.Close()
	
	route := router.Router(database)
	route.Logger.Fatal(route.Start(":8080"))
}



