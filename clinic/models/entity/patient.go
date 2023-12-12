package entity

import "time"

type Patients struct {
	ID        int       `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	Name      string    `json:"name,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Gender    string    `json:"gender,omitempty"`
	Address   string    `json:"address,omitempty"`
	Birth     time.Time `json:"birth,omitempty"`
	CreatedAt int64     `json:"created_at,omitempty"`
	UpdatedAt int64     `json:"updated_at,omitempty"`
}
