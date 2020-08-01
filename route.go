package main

import (
	"github.com/kchauxs/api_simple/blouse"
	"github.com/kchauxs/api_simple/user"
	"github.com/labstack/echo"
)

func startRoutes(e *echo.Echo) {
	//e.POST("/blouses", blouse.Create, user.ValidateJWT)
	e.POST("/blouses", blouse.Create)
	e.POST("/users", user.Login)
	e.POST("/api/v1/users", user.Create)
	e.GET("/api/v1/users", user.GetAll)
	e.GET("/api/v1/users/:email", user.GetByEmail)
	e.GET("/api/v1/users-paginate", user.GetAllPaginate)
}
