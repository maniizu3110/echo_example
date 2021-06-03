package main

import (
	"myapp/src/db"
	"myapp/src/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	database :=db.OpenDB()
	defer database.Close()
	
	route := router.Run(database)
	route.Logger.Fatal(route.Start(":8080"))
}



