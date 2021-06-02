package main

import (
	"myapp/src/db"
	"myapp/src/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	database :=db.OpenDB()
	defer database.Close()
	
	e := router.New(database)
	e.Start(":8080")
}



