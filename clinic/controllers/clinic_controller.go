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

// @Summary 	Create clinic (admin only)
// @Description Create a new clinic for the logged in admin. You will need an 'Authorization' cookie attached with this request.
// @Tags 		Clinic
// @Accept 		json
// @Produce 	json
// @Param 		clinicRequest body dto.ClinicReq true "Clinic details"
// @Success 	201 {object} dto.SwCreateClinicRes
// @Failure 	400 {object} dto.ErrWebResponse
// @Failure 	401 {object} dto.ErrWebResponse
// @Failure 	500 {object} dto.ErrWebResponse
// @Router 		/clinic [post]
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

// @Summary 	List of clinic
// @Description List of all clinic.
// @Tags 		Clinic
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} dto.SwListClinicRes
// @Failure 	404 {object} dto.ErrWebResponse
// @Failure 	500 {object} dto.ErrWebResponse
// @Router 		/clinic [get]
func (cl *ClinicImpl) List(c echo.Context) error {
	clinics, status, err := cl.ClinicService.List(c.Request().Context())
	if err != nil {
		return helpers.ErrorCheck(c, status, err.Error())
	}

	if len(clinics) == 0 {
		return utils.ErrorMessage(c, &utils.ApiNotFound, "no clinic found")
	}

	return utils.SuccessMessage(c, &utils.ApiOk, clinics)
}

// @Summary 	Find By ID
// @Description Find clinic by clinic ID.
// @Tags 		Clinic
// @Accept 		json
// @Produce 	json
// @Param 			id path string true "Clinic ID"
// @Success 	200 {object} dto.SwCreateClinicRes
// @Failure 	400 {object} dto.ErrWebResponse
// @Failure 	404 {object} dto.ErrWebResponse
// @Failure 	500 {object} dto.ErrWebResponse
// @Router 		/clinic/{id} [get]
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

// @Summary 	Update By ID (admin only)
// @Description Update clinic data by clinic ID for the logged in admin. You will need an 'Authorization' cookie attached with this request.
// @Tags 		Clinic
// @Accept 		json
// @Produce 	json
// @Param 			id path string true "Clinic ID"
// @Param 		clinicRequest body dto.ClinicUpdateReq true "Clinic details"
// @Success 	200 {object} dto.SwCreateClinicRes
// @Failure 	400 {object} dto.ErrWebResponse
// @Failure 	401 {object} dto.ErrWebResponse
// @Failure 	404 {object} dto.ErrWebResponse
// @Failure 	500 {object} dto.ErrWebResponse
// @Router 		/clinic/{id} [put]
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

// @Summary 	Delete By ID (admin only)
// @Description Delete clinic data by clinic ID for the logged in admin. You will need an 'Authorization' cookie attached with this request.
// @Tags 		Clinic
// @Accept 		json
// @Produce 	json
// @Param 			id path string true "Clinic ID"
// @Success 	200 {object} dto.SwDeleteClinicRes
// @Failure 	400 {object} dto.ErrWebResponse
// @Failure 	401 {object} dto.ErrWebResponse
// @Failure 	404 {object} dto.ErrWebResponse
// @Failure 	500 {object} dto.ErrWebResponse
// @Router 		/clinic/{id} [delete]
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
