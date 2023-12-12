package utils

import (
	"clinic/models/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ApiOk = dto.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	ApiDelete = dto.WebResponse{
		Code:   http.StatusOK,
		Status: "Delete Success",
	}

	ApiCreate = dto.WebResponse{
		Code:   http.StatusCreated,
		Status: "Create Success",
	}

	ApiUpdate = dto.WebResponse{
		Code:   http.StatusOK,
		Status: "Update Success",
	}
)

func SuccessMessage(c echo.Context, apiSuccess *dto.WebResponse, data any) error {
	apiSuccess.Data = data

	return c.JSON(apiSuccess.Code, apiSuccess)
}
