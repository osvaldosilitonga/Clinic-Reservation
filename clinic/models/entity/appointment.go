package entity

type Appointments struct {
	ID               uint   `json:"id,omitempty"`
	PatientEmail     string `json:"patient_email,omitempty"`
	ClinicID         uint   `json:"clinic_id,omitempty"`
	EmployeeUsername string `json:"employee_username,omitempty" gorm:"<-:update"`
	AppointmentDate  int64  `json:"appointment_date,omitempty"`
	QueueNumber      int    `json:"queue_number,omitempty"`
	Status           string `json:"status,omitempty"`
	Price            int    `json:"price,omitempty"`
	Description      string `json:"description,omitempty"`
	CreatedAt        int64  `json:"created_at,omitempty"`
	UpdatedAt        int64  `json:"updated_at,omitempty"`
}
