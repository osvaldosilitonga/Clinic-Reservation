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
)

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

	PORT := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", PORT)))
}
