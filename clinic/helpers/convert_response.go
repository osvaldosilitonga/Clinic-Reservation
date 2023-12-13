package helpers

import (
	"clinic/models/dto"
	"clinic/models/entity"
)

func ToClinicResponse(d *entity.Clinics) *dto.ClinicRes {
	return &dto.ClinicRes{
		ID:        d.ID,
		Name:      d.Name,
		Phone:     d.Phone,
		Address:   d.Address,
		Slot:      d.Slot,
		CreatedAt: ConvertTimeLocal(d.CreatedAt),
		UpdatedAt: ConvertTimeLocal(d.UpdatedAt),
	}
}
