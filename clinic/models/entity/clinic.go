package entity

type Clinics struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Slot      int    `json:"slot"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
