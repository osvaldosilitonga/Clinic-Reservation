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

// @Summary 	Create appointment (patient only)
// @Description Create a new appointment for the logged in patient. You will need an 'Authorization' cookie attached with this request.
// @Tags 		Appointment
// @Accept 		json
// @Produce 	json
// @Param 		appointmentRequest body dto.CreateAppointmentReq true "Appointment details"
// @Success 	201 {object} dto.SwCreateAppointmentRes
// @Failure 	400 {object} dto.ErrWebResponse
// @Failure 	401 {object} dto.ErrWebResponse
// @Failure 	409 {object} dto.ErrWebResponse
// @Failure 	500 {object} dto.ErrWebResponse
// @Router 		/appointment [post]
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

	if ad.Before(time.Now()) {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "appointment schedules can only be set at least one day before")
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

// @Summary 	Cancel appointment (patient only)
// @Description Cancel scheduled appointment for the logged in patient. You will need an 'Authorization' cookie attached with this request.
// @Tags 		Appointment
// @Accept 		json
// @Produce 	json
// @Param 			id path string true "Appointment ID"
// @Success 	200 {object} dto.SwCancelAppointmentRes
// @Failure 	400 {object} dto.ErrWebResponse
// @Failure 	401 {object} dto.ErrWebResponse
// @Failure 	403 {object} dto.ErrWebResponse
// @Failure 	404 {object} dto.ErrWebResponse
// @Failure 	409 {object} dto.ErrWebResponse
// @Failure 	500 {object} dto.ErrWebResponse
// @Router 		/appointment/cancel/{id} [delete]
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

// @Summary 	Confirm appointment (admin only)
// @Description Confirm appointment and change status to success for the logged in admin. You will need an 'Authorization' cookie attached with this request.
// @Tags 		Appointment
// @Accept 		json
// @Produce 	json
// @Param 			id path string true "Appointment ID"
// @Param 		appointmentRequest body dto.ConfirmAppointmentReq true "Appointment details"
// @Success 	200 {object} dto.SwCancelAppointmentRes
// @Failure 	400 {object} dto.ErrWebResponse
// @Failure 	401 {object} dto.ErrWebResponse
// @Failure 	404 {object} dto.ErrWebResponse
// @Failure 	409 {object} dto.ErrWebResponse
// @Failure 	500 {object} dto.ErrWebResponse
// @Router 		/appointment/confirm/{id} [put]
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

// @Summary 	Find By ID (owner and admin)
// @Description Find appointment by appointment ID. You will need an 'Authorization' cookie attached with this request.
// @Tags 		Appointment
// @Accept 		json
// @Produce 	json
// @Param 			id path string true "Appointment ID"
// @Success 	200 {object} dto.SwCancelAppointmentRes
// @Failure 	400 {object} dto.ErrWebResponse
// @Failure 	401 {object} dto.ErrWebResponse
// @Failure 	403 {object} dto.ErrWebResponse
// @Failure 	404 {object} dto.ErrWebResponse
// @Failure 	500 {object} dto.ErrWebResponse
// @Router 		/appointment/{id} [get]
func (a *AppointmentImpl) FindByID(c echo.Context) error {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "invalid id")
	}

	res, status, err := a.AppointmentService.FindByID(c.Request().Context(), id)
	if err != nil {
		return helpers.ErrorCheck(c, status, err.Error())
	}

	if res.PatientEmail != claims.Email && claims.Role != "admin" {
		return utils.ErrorMessage(c, &utils.ApiForbidden, "you are not authorized to access this data")
	}

	return utils.SuccessMessage(c, &utils.ApiOk, res)
}

// @Summary 	Find By Email (owner and admin)
// @Description Find appointment by patient email. You will need an 'Authorization' cookie attached with this request.
// @Tags 		Appointment
// @Accept 		json
// @Produce 	json
// @Param 			email path string true "Patient Email"
// @Param        status    query     string  false  "Filter by appointment status"
// @Success 	200 {object} dto.SwFinByEmailAppointmentRes
// @Failure 	400 {object} dto.ErrWebResponse
// @Failure 	401 {object} dto.ErrWebResponse
// @Failure 	403 {object} dto.ErrWebResponse
// @Failure 	404 {object} dto.ErrWebResponse
// @Failure 	500 {object} dto.ErrWebResponse
// @Router 		/appointment/email/{email} [get]
func (a *AppointmentImpl) FindByEmail(c echo.Context) error {
	claims, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	paramEmail := c.Param("email")
	if paramEmail != claims.Email && claims.Role != "admin" {
		return utils.ErrorMessage(c, &utils.ApiForbidden, "you are not authorized to access this data")
	}

	filter := map[string]interface{}{}
	filter["patient_email"] = paramEmail

	queryStatus := c.QueryParam("status")
	if queryStatus != "" {
		filter["status"] = queryStatus
	}

	res, code, err := a.AppointmentService.FindByPatientEmail(c.Request().Context(), filter)
	if err != nil {
		return helpers.ErrorCheck(c, code, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiOk, res)
}
