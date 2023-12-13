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

	"gorm.io/gorm"
)

type ClinicImpl struct {
	ClinicRepo repositories.Clinic
}

func NewClinicService(cr repositories.Clinic) Clinic {
	return &ClinicImpl{
		ClinicRepo: cr,
	}
}

func (cl *ClinicImpl) Create(ctx context.Context, d *dto.ClinicReq) (*entity.Clinics, int, error) {
	clinic := &entity.Clinics{
		Name:    d.Name,
		Address: d.Address,
		Phone:   d.Phone,
		Slot:    d.Slot,
	}

	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	clinic, err := cl.ClinicRepo.Save(c, clinic)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return clinic, http.StatusCreated, nil
}

func (cl *ClinicImpl) List(ctx context.Context) ([]dto.ClinicRes, int, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	clinics, err := cl.ClinicRepo.List(c)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	res := []dto.ClinicRes{}
	for _, clinic := range clinics {
		r := dto.ClinicRes{
			ID:        clinic.ID,
			Name:      clinic.Name,
			Phone:     clinic.Phone,
			Address:   clinic.Address,
			Slot:      clinic.Slot,
			CreatedAt: helpers.ConvertTimeLocal(clinic.CreatedAt),
			UpdatedAt: helpers.ConvertTimeLocal(clinic.UpdatedAt),
		}

		res = append(res, r)
	}

	return res, http.StatusOK, nil
}

func (cl *ClinicImpl) FindByID(ctx context.Context, id int) (*dto.ClinicRes, int, error) {
	clinic, err := cl.ClinicRepo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, err
		}

		return nil, http.StatusInternalServerError, err
	}

	res := dto.ClinicRes{
		ID:        clinic.ID,
		Name:      clinic.Name,
		Phone:     clinic.Phone,
		Address:   clinic.Address,
		Slot:      clinic.Slot,
		CreatedAt: helpers.ConvertTimeLocal(clinic.CreatedAt),
		UpdatedAt: helpers.ConvertTimeLocal(clinic.UpdatedAt),
	}

	return &res, http.StatusOK, nil
}

func (cl *ClinicImpl) Update(ctx context.Context, d *dto.ClinicUpdateReq, id int) (*dto.ClinicRes, int, error) {
	clinic := &entity.Clinics{
		Name:      d.Name,
		Address:   d.Address,
		Phone:     d.Phone,
		Slot:      d.Slot,
		UpdatedAt: time.Now().UnixMilli(),
	}

	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	clinic, err := cl.ClinicRepo.Update(c, clinic, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, err
		}

		return nil, http.StatusInternalServerError, err
	}

	res := dto.ClinicRes{
		ID:        clinic.ID,
		Name:      clinic.Name,
		Phone:     clinic.Phone,
		Address:   clinic.Address,
		Slot:      clinic.Slot,
		CreatedAt: helpers.ConvertTimeLocal(clinic.CreatedAt),
		UpdatedAt: helpers.ConvertTimeLocal(clinic.UpdatedAt),
	}

	return &res, http.StatusOK, nil
}

func (cl *ClinicImpl) Delete(ctx context.Context, id int) (int, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := cl.ClinicRepo.Delete(c, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, err
		}

		return http.StatusInternalServerError, err
	}

	if rows == 0 {
		return http.StatusNotFound, errors.New("clinic not found")
	}

	return http.StatusOK, nil
}
