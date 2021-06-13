package main

import (
	"myapp/src/db"
	"myapp/src/router"
	"net"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	_, err := net.Listen("tcp", "127.0.0.1:8000")
    if err != nil {
        return
    }
	var database *gorm.DB
	database = db.InitDB()
	defer database.Close()
	
	route := router.Router(database)
	route.Logger.Fatal(route.Start(":8080"))
}



