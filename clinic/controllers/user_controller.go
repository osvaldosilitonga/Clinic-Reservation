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

// @Summary 	Patient Register
// @Description Register new patient
// @Tags 			Default
// @Accept 		json
// @Produce 	json
// @Param 		data body dto.UserRegisterReq true "Patient Data"
// @Success 	201 {object} dto.SwUserRegisterRes
// @Failure 	400 {object} dto.ErrWebResponse
// @Failure 	500 {object} dto.ErrWebResponse
// @Router 		/register [post]
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

// @Summary 	Patient Login
// @Description Login for patient
// @Tags 			Default
// @Accept 		json
// @Produce 	json
// @Param 		data body dto.UserLoginReq true "Patient Credentials"
// @Success 	200 {object} dto.SwLoginRes
// @Failure 	400 {object} dto.ErrWebResponse
// @Failure 	404 {object} dto.ErrWebResponse
// @Failure 	500 {object} dto.ErrWebResponse
// @Router 		/login [post]
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

// @Summary 	Logout the user
// @Description Logout the currently authenticated user and clears the authorization cookie
// @Tags        Default
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} dto.SwLoginRes
// @Failure 	500 {object} dto.ErrWebResponse
// @Router 		/logout [get]
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

// @Summary 	Employee Register
// @Description Register new employee
// @Tags 			Default
// @Accept 		json
// @Produce 	json
// @Param 		data body dto.EmployeeRegisterReq true "Employee Data"
// @Success 	201 {object} dto.SwUserRegisterRes
// @Failure 	400 {object} dto.ErrWebResponse
// @Failure 	500 {object} dto.ErrWebResponse
// @Router 		/employee/register [post]
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

// @Summary 	Employee Login
// @Description Login for employee
// @Tags 			Default
// @Accept 		json
// @Produce 	json
// @Param 		data body dto.EmployeeLoginReq true "Employee Credentials"
// @Success 	200 {object} dto.SwLoginRes
// @Failure 	400 {object} dto.ErrWebResponse
// @Failure 	404 {object} dto.ErrWebResponse
// @Failure 	500 {object} dto.ErrWebResponse
// @Router 		/employee/login [post]
func (u *UserImpl) EmployeeLogin(c echo.Context) error {
	body := dto.EmployeeLoginReq{}
	if err := c.Bind(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}
	if err := c.Validate(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}

	employee, code, err := u.EmployeeService.Login(c.Request().Context(), &body)
	if err != nil {
		return helpers.ErrorCheck(c, code, err.Error())
	}

	if err := helpers.SignNewJWT(c, int(employee.ID), employee.Username, employee.Role); err != nil {
		return err
	}

	return utils.SuccessMessage(c, &utils.ApiOk, "Login Success. Authorization token has been set to cookie")
}
