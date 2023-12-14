package router

import (
	"clinic/configs"
	"clinic/controllers"
	"clinic/middlewares"
	"clinic/repositories"
	"clinic/services"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	db := configs.InitDB()

	// Repositories
	patientRepository := repositories.NewPatientRepository(db)
	employeeRepository := repositories.NewEmployeeRepository(db)
	clinicRepository := repositories.NewClinicRepository(db)
	appointmentRepository := repositories.NewAppointmentRepository(db)

	// Services
	userService := services.NewUserService(&services.UserConfig{
		PatientRepo: patientRepository,
	})
	employeeService := services.NewEmployeeService(employeeRepository)
	clinicService := services.NewClinicService(clinicRepository)
	appointmentService := services.NewAppointmentService(appointmentRepository, clinicRepository)

	// Controllers
	userController := controllers.NewUserController(&controllers.UserConfig{
		UserService:     userService,
		EmployeeService: employeeService,
	})
	clinicController := controllers.NewClinicController(clinicService)
	appointmentController := controllers.NewAppointmentController(appointmentService)

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.POST("/register", userController.Register)
		v1.POST("/login", userController.Login)
		v1.GET("/logout", userController.Logout)

		v1.POST("/employee/register", userController.EmployeeRegister)
		v1.POST("/employee/login", userController.EmployeeLogin)
	}

	clinic := v1.Group("/clinic")
	{
		clinic.POST("", clinicController.CreateClinic, middlewares.RequireAuth, middlewares.IsAdmin)
		clinic.GET("", clinicController.List)
		clinic.GET("/:id", clinicController.FindByID)
		clinic.PUT("/:id", clinicController.Update, middlewares.RequireAuth, middlewares.IsAdmin)
		clinic.DELETE("/:id", clinicController.Delete, middlewares.RequireAuth, middlewares.IsAdmin)
	}

	appointment := v1.Group("/appointment")
	{
		appointment.POST("", appointmentController.CreateAppointment, middlewares.RequireAuth, middlewares.IsPatient)
	}

}
