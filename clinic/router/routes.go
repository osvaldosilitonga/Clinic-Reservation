package router

import (
	"clinic/configs"
	"clinic/controllers"
	"clinic/repositories"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	db := configs.InitDB()

	patientRepository := repositories.NewPatientRepository(db)

	userConfig := &controllers.UserConfig{
		PatientRepo: patientRepository,
	}

	v1 := e.Group("/api/v1")
	userController := controllers.NewUserController(userConfig)
	{
		v1.GET("/register", userController.Register)
	}
}
