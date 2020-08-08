package main

import (
	"github.com/kchauxs/api_simple/user"
	"github.com/labstack/echo"
)

func startRoutes(e *echo.Echo) {
	e.POST("/watches", watches.Create, user.ValidateJWT)
	e.GET("/watches", watches.GetAll)
	e.POST("/users", user.Login)
	e.POST("/api/v1/users", user.Create)
	e.GET("/api/v1/users", user.GetAll)
	e.GET("/api/v1/users/:email", user.GetByEmail)
	e.GET("/api/v1/users-paginate", user.GetAllPaginate)
}
