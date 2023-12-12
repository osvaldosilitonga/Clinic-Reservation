package controllers

import (
	"clinic/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserImpl struct {
	PatientRepo repositories.Patient
}

type UserConfig struct {
	PatientRepo repositories.Patient
}

func NewUserController(ucc *UserConfig) User {
	return &UserImpl{
		PatientRepo: ucc.PatientRepo,
	}
}

func (u *UserImpl) Register(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Hello World",
	})
}
