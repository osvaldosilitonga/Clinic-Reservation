package dto

import "clinic/models/entity"

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

type SwLoginRes struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   string `json:"data"`
}
