package services

import (
	"clinic/models/dto"
	"clinic/repositories"
	"context"
)

type UserImpl struct {
	PatientRepo repositories.Patient
}

type UserConfig struct {
	PatientRepo repositories.Patient
}

func NewUserService(uc *UserConfig) User {
	return &UserImpl{
		PatientRepo: uc.PatientRepo,
	}
}

func (u *UserImpl) Register(ctx context.Context, d *dto.User) error {
	return nil
}
