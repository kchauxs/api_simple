package main

import (
	"fmt"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	startRoutes(e)

	err := e.Start(":8080")
	if err != nil {
		fmt.Printf("No pude subir el servidor %v", err)
	}
}
