package services

import (
	"clinic/helpers"
	"clinic/models/dto"
	"clinic/models/entity"
	"clinic/repositories"
	"context"
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func (u *UserImpl) Register(c context.Context, d *dto.UserRegisterReq) (*dto.UserResponse, int, error) {
	user := &entity.Patients{
		Email:   d.Email,
		Name:    d.Name,
		Phone:   d.Phone,
		Gender:  d.Gender,
		Address: d.Address,
		Role:    "patient",
	}

	birth, err := time.Parse("2006-01-02", d.Birth)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	user.Birth = birth

	hash, err := bcrypt.GenerateFromPassword([]byte(d.Password), 14)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	user.Password = string(hash)

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	res, err := u.PatientRepo.Save(ctx, user)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, http.StatusBadRequest, err
		}
		return nil, http.StatusInternalServerError, err
	}

	res.Scan(&user)

	response := helpers.ToUserResponse(user)

	return &response, http.StatusCreated, nil
}
