package dto

type ClinicReq struct {
	Name    string `json:"name" validate:"required,min=6,max=50"`
	Address string `json:"address" validate:"required,min=5"`
	Phone   string `json:"phone" validate:"required,min=6,max=13"`
	Slot    int    `json:"slot" validate:"required,min=1"`
}
