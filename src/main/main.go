package main

import (
	"fmt"
	"myapp/src/router"
)

func main(){
	fmt.Println("Welcome to the webserver")
	e :=router.New()
	e.Start(":8080")
}