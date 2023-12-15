package dto

import "clinic/models/entity"

type SwLoginRes struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   string `json:"data"`
}

type SwUserRegisterRes struct {
	Code   int          `json:"code"`
	Status string       `json:"status"`
	Data   UserResponse `json:"data"`
}

type SwEmployeeRegisterRes struct {
	Code   int              `json:"code"`
	Status string           `json:"status"`
	Data   entity.Employees `json:"data"`
}

type SwCreateAppointmentRes struct {
	Code   int                  `json:"code"`
	Status string               `json:"status"`
	Data   CreateAppointmentRes `json:"data"`
}

type SwCancelAppointmentRes struct {
	Code   int            `json:"code"`
	Status string         `json:"status"`
	Data   AppointmentRes `json:"data"`
}

type SwConfirmAppointmentRes struct {
	Code   int                `json:"code"`
	Status string             `json:"status"`
	Data   FullAppointmentRes `json:"data"`
}

type SwFinByEmailAppointmentRes struct {
	Code   int                  `json:"code"`
	Status string               `json:"status"`
	Data   []FullAppointmentRes `json:"data"`
}

type SwCreateClinicRes struct {
	Code   int       `json:"code"`
	Status string    `json:"status"`
	Data   ClinicRes `json:"data"`
}

type SwListClinicRes struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   []ClinicRes `json:"data"`
}

type SwDeleteClinicRes struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   string `json:"data"`
}
