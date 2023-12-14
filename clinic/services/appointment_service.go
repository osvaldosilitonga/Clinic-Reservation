package services

import (
	"clinic/helpers"
	"clinic/models/dto"
	"clinic/models/entity"
	"clinic/repositories"
	"context"
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type AppointmentImpl struct {
	AppointmentRepo repositories.Appointment
	ClinicRepo      repositories.Clinic
}

func NewAppointmentService(ar repositories.Appointment, cr repositories.Clinic) Appointment {
	return &AppointmentImpl{
		AppointmentRepo: ar,
		ClinicRepo:      cr,
	}
}

func (a *AppointmentImpl) Create(ctx context.Context, d *entity.Appointments) (*dto.CreateAppointmentRes, int, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Queue Number - Count Record Clinic
	d.Status = "booked"
	record, err := a.AppointmentRepo.CountRecordClinic(c, d.AppointmentDate, d.ClinicID, d.Status)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	d.QueueNumber = int(record) + 1

	// Get Clinic Slot
	clinic, err := a.ClinicRepo.FindByID(c, int(d.ClinicID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, err
	}

	// Check Clinic Slot
	if record >= int64(clinic.Slot) {
		return nil, http.StatusConflict, errors.New("appointment for this clinic is full")
	}

	// Save Appointment
	appointment, err := a.AppointmentRepo.Save(c, d)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res := dto.CreateAppointmentRes{
		ID:              appointment.ID,
		PatientEmail:    appointment.PatientEmail,
		ClinicID:        appointment.ClinicID,
		AppointmentDate: helpers.ConvertTimeLocal(appointment.AppointmentDate),
		QueueNumber:     appointment.QueueNumber,
		Status:          appointment.Status,
		Description:     appointment.Description,
		CreatedAt:       helpers.ConvertTimeLocal(appointment.CreatedAt),
		UpdatedAt:       helpers.ConvertTimeLocal(appointment.UpdatedAt),
	}

	return &res, http.StatusCreated, nil
}
