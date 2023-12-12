package dto

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Name     string `json:"name" validate:"required,min=3,max=32"`
	Phone    string `json:"phone" validate:"required,min=10,max=13"`
	Gender   string `json:"gender" validate:"required,min=4,max=6"`
	Address  string `json:"address" validate:"required,min=10,max=128"`
	Birth    string `json:"birth" validate:"required,min=10,max=10"`
}
