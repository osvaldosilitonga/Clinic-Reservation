package main

import (
	"clinic/helpers"
	"clinic/initializers"
	"clinic/router"
	"fmt"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "clinic/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Zi.Care Clinic - API Documentation
// @version BETA
// @description Application for managing clinic data and appointments

// @contact.name Zi.Care
// @contact.url www.zicare.com
// @contact.email zicare@mail.com

// @host localhost:8080
// @BasePath /api/v1
// @license.name Apache 2.0

func init() {
	initializers.LoadEnvFile()
	initializers.RunDBMigrations(os.Getenv("MIGRATION_URL"), os.Getenv("DB_PG_STRING"))
}

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Validator = &helpers.CustomValidator{Validator: validator.New()}

	router.Routes(e)

	e.GET("/swagger/*.html", echoSwagger.WrapHandler)

	PORT := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", PORT)))
}
