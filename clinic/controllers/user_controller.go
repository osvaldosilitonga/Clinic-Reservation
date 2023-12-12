package controllers

import (
	"clinic/helpers"
	"clinic/models/dto"
	"clinic/services"
	"clinic/utils"

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
	body := dto.UserRegisterReq{}
	if err := c.Bind(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}
	if err := c.Validate(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}

	user, code, err := u.UserService.Register(c.Request().Context(), &body)
	if err != nil {
		return helpers.ErrorCheck(c, code, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiCreate, user)
}
