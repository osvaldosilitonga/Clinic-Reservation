package services

import (
	"clinic/helpers"
	"clinic/models/dto"
	"clinic/models/entity"
	"clinic/repositories"
	"context"
	"net/http"
	"time"
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
