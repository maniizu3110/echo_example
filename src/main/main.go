package main

import (
	"fmt"
	// "log"
	"myapp/src/db"
	"myapp/src/firebaase-admin/snippets"
	"myapp/src/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	db.InitDB()
	snippets.InitializeAppDefault()
	
}

func main() {
	fmt.Println("Welcome to the webserver")
	e := router.New()
	e.Start(":8080")
}



