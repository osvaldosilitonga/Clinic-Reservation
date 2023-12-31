package dto

import "time"

type AppointmentFilter struct {
	PatientEmail     string `json:"patient_email,omitempty"`
	ClinicID         int    `json:"clinic_id,omitempty"`
	EmployeeUsername string `json:"employee_username,omitempty"`
	AppointmentDate  int64  `json:"appointment_date,omitempty"`
	Status           string `json:"status,omitempty"`
}

type CreateAppointmentReq struct {
	AppointmentDate string `json:"appointment_date" validate:"required"`
	ClinicID        uint   `json:"clinic_id" validate:"required,numeric"`
	Description     string `json:"description" validate:"required"`
}

type ConfirmAppointmentReq struct {
	Price int `json:"price" validate:"required,numeric,min=10000"`
}

type CreateAppointmentRes struct {
	ID              uint      `json:"id"`
	PatientEmail    string    `json:"patient_email"`
	ClinicID        uint      `json:"clinic_id"`
	AppointmentDate time.Time `json:"appointment_date"`
	QueueNumber     int       `json:"queue_number"`
	Status          string    `json:"status"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type AppointmentRes struct {
	ID              uint      `json:"id"`
	PatientEmail    string    `json:"patient_email"`
	ClinicID        uint      `json:"clinic_id"`
	AppointmentDate time.Time `json:"appointment_date"`
	QueueNumber     int       `json:"queue_number"`
	Status          string    `json:"status"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type FullAppointmentRes struct {
	ID               uint      `json:"id"`
	PatientEmail     string    `json:"patient_email"`
	ClinicID         uint      `json:"clinic_id"`
	EmployeeUsername string    `json:"employee_username"`
	AppointmentDate  time.Time `json:"appointment_date"`
	QueueNumber      int       `json:"queue_number"`
	Status           string    `json:"status"`
	Price            int       `json:"price"`
	Description      string    `json:"description"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
