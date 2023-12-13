package controllers

import "github.com/labstack/echo/v4"

type User interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error

	EmployeeRegister(c echo.Context) error
	EmployeeLogin(c echo.Context) error
}

type Clinic interface {
	CreateClinic(c echo.Context) error
	List(c echo.Context) error
	FindByID(c echo.Context) error
}
