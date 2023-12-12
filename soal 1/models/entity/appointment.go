package entity

type Appointments struct {
	ID               int    `json:"id"`
	PatientEmail     string `json:"patient_email"`
	ClinicId         int    `json:"clinic_id"`
	EmployeeUsername string `json:"employee_username"`
	AppointmentDate  string `json:"appointment_date"`
	QueueingNumber   int    `json:"queueing_number"`
	Status           string `json:"status"`
	Price            int    `json:"price"`
	Description      string `json:"description"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
