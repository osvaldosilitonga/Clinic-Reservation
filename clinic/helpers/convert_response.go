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

func ToAppointmentResponse(d *entity.Appointments) *dto.FullAppointmentRes {
	return &dto.FullAppointmentRes{
		ID:               d.ID,
		PatientEmail:     d.PatientEmail,
		ClinicID:         d.ClinicID,
		EmployeeUsername: d.EmployeeUsername,
		AppointmentDate:  ConvertTimeLocal(d.AppointmentDate),
		QueueNumber:      d.QueueNumber,
		Status:           d.Status,
		Price:            d.Price,
		Description:      d.Description,
		CreatedAt:        ConvertTimeLocal(d.CreatedAt),
		UpdatedAt:        ConvertTimeLocal(d.UpdatedAt),
	}
}
