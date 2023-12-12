package entity

type Employees struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
}

type Roles struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Access []string `json:"access"`
}
