package router

import (
	"clinic/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {

	userController := controllers.NewUserController()

	v1 := e.Group("/api/v1")
	{
		v1.GET("/register", userController.Register)
	}
}
