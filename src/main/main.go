package main

import (
	"fmt"
	"myapp/src/api/firebase/snippets"
	"myapp/src/db"
	"myapp/src/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	db.InitDB()
	snippets.InitializeAppDefault()
	
}

func main() {
	router := router.Run()
	fmt.Println(router)
}



