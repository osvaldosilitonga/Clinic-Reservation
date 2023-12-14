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

func (a *AppointmentImpl) Cancel(ctx context.Context, id int, email string) (*dto.AppointmentRes, int, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	appointment, err := a.AppointmentRepo.FindByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, err
	}

	if appointment.PatientEmail != email {
		return nil, http.StatusForbidden, errors.New("you are not authorized to cancel this appointment")
	}

	now := time.Now().Unix()
	if appointment.AppointmentDate <= now {
		return nil, http.StatusConflict, errors.New("cannot cancel appointment. please cancel 1 day before the appointment")
	}

	update := map[string]interface{}{
		"status": "canceled",
	}
	appointmentUpdate, err := a.AppointmentRepo.Update(c, id, &update)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, err
	}

	res := dto.AppointmentRes{
		ID:              appointment.ID,
		PatientEmail:    appointment.PatientEmail,
		ClinicID:        appointment.ClinicID,
		AppointmentDate: helpers.ConvertTimeLocal(appointment.AppointmentDate),
		QueueNumber:     appointment.QueueNumber,
		Status:          appointmentUpdate.Status,
		Description:     appointment.Description,
		CreatedAt:       helpers.ConvertTimeLocal(appointment.CreatedAt),
		UpdatedAt:       helpers.ConvertTimeLocal(appointment.UpdatedAt),
	}

	return &res, http.StatusOK, nil
}

func (a *AppointmentImpl) Confirm(ctx context.Context, id int, email string, price int) (*dto.FullAppointmentRes, int, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	appointment, err := a.AppointmentRepo.FindByID(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, err
	}

	switch appointment.Status {
	case "canceled":
		return nil, http.StatusConflict, errors.New("appointment is already canceled")
	case "success":
		return nil, http.StatusConflict, errors.New("appointment is already confirmed")
	}

	update := map[string]interface{}{
		"employee_username": email,
		"status":            "success",
		"price":             price,
		"updated_at":        time.Now().Unix(),
	}
	appointmentUpdate, err := a.AppointmentRepo.Update(c, id, &update)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, err
	}

	res := dto.FullAppointmentRes{
		ID:               appointment.ID,
		PatientEmail:     appointment.PatientEmail,
		ClinicID:         appointment.ClinicID,
		EmployeeUsername: appointmentUpdate.EmployeeUsername,
		AppointmentDate:  helpers.ConvertTimeLocal(appointment.AppointmentDate),
		QueueNumber:      appointment.QueueNumber,
		Status:           appointmentUpdate.Status,
		Price:            appointmentUpdate.Price,
		Description:      appointment.Description,
		CreatedAt:        helpers.ConvertTimeLocal(appointment.CreatedAt),
		UpdatedAt:        helpers.ConvertTimeLocal(appointment.UpdatedAt),
	}

	return &res, http.StatusOK, nil
}
