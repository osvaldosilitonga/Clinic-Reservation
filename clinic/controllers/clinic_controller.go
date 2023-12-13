package controllers

import (
	"clinic/helpers"
	"clinic/models/dto"
	"clinic/services"
	"clinic/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ClinicImpl struct {
	ClinicService services.Clinic
}

func NewClinicController(cs services.Clinic) Clinic {
	return &ClinicImpl{
		ClinicService: cs,
	}
}

func (cl *ClinicImpl) CreateClinic(c echo.Context) error {
	body := dto.ClinicReq{}
	if err := c.Bind(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}
	if err := c.Validate(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}

	clinic, status, err := cl.ClinicService.Create(c.Request().Context(), &body)
	if err != nil {
		return helpers.ErrorCheck(c, status, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiCreate, clinic)
}

func (cl *ClinicImpl) List(c echo.Context) error {
	clinics, status, err := cl.ClinicService.List(c.Request().Context())
	if err != nil {
		return helpers.ErrorCheck(c, status, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiOk, clinics)
}

func (cl *ClinicImpl) FindByID(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "invalid param id")
	}

	clinic, status, err := cl.ClinicService.FindByID(c.Request().Context(), id)
	if err != nil {
		return helpers.ErrorCheck(c, status, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiOk, clinic)
}

func (cl *ClinicImpl) Update(c echo.Context) error {

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "invalid param id")
	}

	body := dto.ClinicUpdateReq{}
	if err := c.Bind(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}
	if err := c.Validate(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}

	if body.Name == "" && body.Phone == "" && body.Address == "" && body.Slot == 0 {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "empty request body not allowed")
	}

	clinic, status, err := cl.ClinicService.Update(c.Request().Context(), &body, id)
	if err != nil {
		return helpers.ErrorCheck(c, status, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiUpdate, clinic)
}

func (cl *ClinicImpl) Delete(c echo.Context) error {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "invalid param id")
	}

	status, err := cl.ClinicService.Delete(c.Request().Context(), id)
	if err != nil {
		return helpers.ErrorCheck(c, status, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiDelete, nil)
}
