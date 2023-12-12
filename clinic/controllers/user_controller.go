package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User interface {
	Register(c echo.Context) error
}

type UserImpl struct {
}

func NewUserController() User {
	return &UserImpl{}
}

func (u *UserImpl) Register(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Hello World",
	})
}
