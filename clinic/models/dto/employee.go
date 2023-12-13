package dto

type EmployeeRegisterReq struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}

type EmployeeLoginReq struct {
	Username string `json:"username" validate:"required,min=5,max=20"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}

type EmployeeResponse struct {
	Username string `json:"username"`
}
