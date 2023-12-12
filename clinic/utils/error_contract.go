package utils

import (
	"clinic/models/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ApiBadRequest = dto.ErrWebResponse{
		Code:   http.StatusBadRequest,
		Status: "Bad Request",
	}

	ApiNotFound = dto.ErrWebResponse{
		Code:   http.StatusNotFound,
		Status: "not found",
	}

	ApiForbidden = dto.ErrWebResponse{
		Code:   http.StatusForbidden,
		Status: "forbidden",
	}

	ApiInternalServer = dto.ErrWebResponse{
		Code:   http.StatusInternalServerError,
		Status: "internal server error",
	}

	ApiUnauthorized = dto.ErrWebResponse{
		Code:   http.StatusUnauthorized,
		Status: "unauthorized",
	}
)

func ErrorMessage(c echo.Context, apiErr *dto.ErrWebResponse, detail any) error {
	apiErr.Detail = detail

	return echo.NewHTTPError(apiErr.Code, apiErr)
}
