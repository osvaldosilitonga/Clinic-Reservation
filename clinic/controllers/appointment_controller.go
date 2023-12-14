package controllers

import (
	"clinic/helpers"
	"clinic/models/dto"
	"clinic/models/entity"
	"clinic/services"
	"clinic/utils"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type AppointmentImpl struct {
	AppointmentService services.Appointment
}

func NewAppointmentController(as services.Appointment) Appointment {
	return &AppointmentImpl{
		AppointmentService: as,
	}
}

func (a *AppointmentImpl) CreateAppointment(c echo.Context) error {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	body := dto.CreateAppointmentReq{}
	if err := c.Bind(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}
	if err := c.Validate(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}

	ad, err := time.Parse(time.DateOnly, body.AppointmentDate)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "invalid date format. please use YYYY-MM-DD")
	}

	appointment := &entity.Appointments{
		PatientEmail:    claims.Email,
		ClinicID:        body.ClinicID,
		AppointmentDate: ad.Unix(),
		Description:     body.Description,
	}

	res, status, err := a.AppointmentService.Create(c.Request().Context(), appointment)
	if err != nil {
		return helpers.ErrorCheck(c, status, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiCreate, res)
}

func (a *AppointmentImpl) CancelAppointment(c echo.Context) error {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "invalid id")
	}

	res, status, err := a.AppointmentService.Cancel(c.Request().Context(), id, claims.Email)
	if err != nil {
		return helpers.ErrorCheck(c, status, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiUpdate, res)
}

func (a *AppointmentImpl) ConfirmAppointment(c echo.Context) error {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	body := dto.ConfirmAppointmentReq{}
	if err := c.Bind(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "invalid request body")
	}
	if err := c.Validate(&body); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "invalid id")
	}

	res, status, err := a.AppointmentService.Confirm(c.Request().Context(), id, claims.Email, body.Price)
	if err != nil {
		return helpers.ErrorCheck(c, status, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiUpdate, res)
}