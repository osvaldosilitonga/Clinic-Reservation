package dto

import "time"

type ClinicReq struct {
	Name    string `json:"name" validate:"required,min=6,max=50"`
	Address string `json:"address" validate:"required,min=5"`
	Phone   string `json:"phone" validate:"required,min=6,max=13"`
	Slot    int    `json:"slot" validate:"required,min=1"`
}

type ClinicUpdateReq struct {
	Name    string `json:"name,omitempty" validate:"omitempty,min=6,max=50"`
	Address string `json:"address,omitempty" validate:"omitempty,min=5"`
	Phone   string `json:"phone,omitempty" validate:"omitempty,min=6,max=13"`
	Slot    int    `json:"slot,omitempty" validate:"omitempty,min=1"`
}

type ClinicRes struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Slot      int       `json:"slot"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
