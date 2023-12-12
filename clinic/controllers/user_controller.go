package controllers

import (
	"clinic/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserImpl struct {
	UserService services.User
}

type UserConfig struct {
	UserService services.User
}

func NewUserController(uc *UserConfig) User {
	return &UserImpl{
		UserService: uc.UserService,
	}
}

func (u *UserImpl) Register(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Hello World",
	})
}
