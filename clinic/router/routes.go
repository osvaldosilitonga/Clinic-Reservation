package router

import (
	"clinic/configs"
	"clinic/controllers"
	"clinic/repositories"
	"clinic/services"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	db := configs.InitDB()

	// Repositories
	patientRepository := repositories.NewPatientRepository(db)

	// Services
	userService := services.NewUserService(&services.UserConfig{
		PatientRepo: patientRepository,
	})

	// Controllers
	userController := controllers.NewUserController(&controllers.UserConfig{
		UserService: userService,
	})

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.POST("/register", userController.Register)
		v1.POST("/login", userController.Login)
		v1.GET("/logout", userController.Logout)
	}

}
