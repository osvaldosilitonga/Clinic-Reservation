package controllers

import (
	"clinic/helpers"
	"clinic/models/dto"
	"clinic/services"
	"clinic/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserImpl struct {
	UserService     services.User
	EmployeeService services.Employee
}

type UserConfig struct {
	UserService     services.User
	EmployeeService services.Employee
}

func NewUserController(uc *UserConfig) User {
	return &UserImpl{
		UserService:     uc.UserService,
		EmployeeService: uc.EmployeeService,
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

func (u *UserImpl) Login(c echo.Context) error {
	body := dto.UserLoginReq{}
	if err := c.Bind(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}
	if err := c.Validate(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}

	user, code, err := u.UserService.Login(c.Request().Context(), &body)
	if err != nil {
		return helpers.ErrorCheck(c, code, err.Error())
	}

	if err := helpers.SignNewJWT(c, int(user.ID), user.Email, user.Role); err != nil {
		return err
	}

	return utils.SuccessMessage(c, &utils.ApiOk, "Login Success. Authorization token has been set to cookie")
}

func (u *UserImpl) Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "Authorization"
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Value = ""
	cookie.SameSite = http.SameSiteLaxMode
	cookie.MaxAge = -1
	c.SetCookie(cookie)

	return utils.SuccessMessage(c, &utils.ApiOk, "Logout Success. Authorization token has been removed from cookie")
}

func (u *UserImpl) EmployeeRegister(c echo.Context) error {
	body := dto.EmployeeRegisterReq{}
	if err := c.Bind(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}
	if err := c.Validate(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}

	employee, code, err := u.EmployeeService.Register(c.Request().Context(), &body)
	if err != nil {
		return helpers.ErrorCheck(c, code, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiCreate, employee)
}
