package main

import (
	"myapp/src/db"
	"myapp/src/router"
	"github.com/jinzhu/gorm"

	

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	var database *gorm.DB
	database = db.InitDB()
	defer database.Close()
	route := router.Run()
	route.Logger.Fatal(route.Start(":8080"))
}



