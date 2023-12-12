package helpers

import (
	"clinic/models/dto"
	"clinic/models/entity"
)

func ToUserResponse(u *entity.Patients) dto.UserResponse {
	return dto.UserResponse{
		Email:   u.Email,
		Name:    u.Name,
		Phone:   u.Phone,
		Gender:  u.Gender,
		Address: u.Address,
		Birth:   u.Birth.Format("2006-01-02"),
	}
}
