package middlewares

import (
	"clinic/helpers"
	"clinic/utils"

	"github.com/labstack/echo/v4"
)

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, err := helpers.GetClaims(c)
		if err != nil {
			return utils.ErrorMessage(c, &utils.ApiUnauthorized, "Please log in to access this page")
		}

		if claims.Role != "admin" {
			return utils.ErrorMessage(c, &utils.ApiForbidden, "You are not authorized to access this page")
		}

		return next(c)
	}
}
