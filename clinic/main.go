package main

import (
	"clinic/helpers"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Validator = &helpers.CustomValidator{Validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	PORT := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", PORT)))
}
