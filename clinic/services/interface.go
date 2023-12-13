package services

import (
	"clinic/models/dto"
	"clinic/models/entity"
	"context"
)

type User interface {
	Register(ctx context.Context, d *dto.UserRegisterReq) (*dto.UserResponse, int, error)
	Login(ctx context.Context, d *dto.UserLoginReq) (*entity.Patients, int, error)
}

type Employee interface {
	Register(ctx context.Context, d *dto.EmployeeRegisterReq) (*entity.Employees, int, error)
	Login(ctx context.Context, d *dto.EmployeeLoginReq) (*entity.Employees, int, error)
}
