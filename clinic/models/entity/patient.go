package entity

import "time"

type Patients struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Gender    string    `json:"gender"`
	Address   string    `json:"address"`
	Birth     time.Time `json:"birth"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
