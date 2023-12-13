package services

import (
	"clinic/models/dto"
	"clinic/models/entity"
	"clinic/repositories"
	"context"
	"net/http"
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

	clinic, err := cl.ClinicRepo.Save(ctx, clinic)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return clinic, http.StatusCreated, nil
}
