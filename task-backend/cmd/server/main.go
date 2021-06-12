package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shankernaik/task-backend/task-backend/routes"
)

const (
	//PORT 8080 is default port
	PORT = "8080"
)

func main() {

	//TODO Read ports and certificate from external config file

	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.Routes(e)
	bootserver(e)
}

func bootserver(e *echo.Echo) {

	fmt.Println("################################## ENDPOINT(S) #####################################")
	for _, route := range e.Routes() {
		fmt.Printf(route.Method)
		fmt.Printf(" ")
		fmt.Println(route.Path)
	}
	fmt.Println("################################## ENDPOINT(S) #####################################")

	e.Logger.Fatal(e.Start(":" + PORT))
}
