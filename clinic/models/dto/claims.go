package dto

type Claims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
