package entity

type Employees struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Role     string `json:"role"`
}
