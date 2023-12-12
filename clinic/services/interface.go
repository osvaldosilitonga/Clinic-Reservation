package services

import (
	"clinic/models/dto"
	"context"
)

type User interface {
	Register(ctx context.Context, d *dto.UserRegisterReq) (*dto.UserResponse, int, error)
}
