package helpers

import (
	"clinic/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorCheck(c echo.Context, code int, msg string) error {
	switch code {
	case http.StatusBadRequest:
		return utils.ErrorMessage(c, &utils.ApiBadRequest, msg)
	case http.StatusNotFound:
		return utils.ErrorMessage(c, &utils.ApiNotFound, msg)
	case http.StatusForbidden:
		return utils.ErrorMessage(c, &utils.ApiForbidden, msg)
	case http.StatusUnauthorized:
		return utils.ErrorMessage(c, &utils.ApiUnauthorized, msg)
	case http.StatusConflict:
		return utils.ErrorMessage(c, &utils.ApiConflict, msg)
	default:
		return utils.ErrorMessage(c, &utils.ApiInternalServer, msg)
	}
}
