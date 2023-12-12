package entity

type Appointments struct {
	ID               int    `json:"id"`
	PatientEmail     string `json:"patient_email"`
	ClinicID         int    `json:"clinic_id"`
	EmployeeUsername string `json:"employee_username"`
	AppointmentDate  int64  `json:"appointment_date"`
	QueueNumber      int    `json:"queue_number"`
	Status           string `json:"status"`
	Price            int    `json:"price"`
	Description      string `json:"description"`
	CreatedAt        int64  `json:"created_at"`
	UpdatedAt        int64  `json:"updated_at"`
}
