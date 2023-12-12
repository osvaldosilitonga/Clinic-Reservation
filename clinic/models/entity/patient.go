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
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}
