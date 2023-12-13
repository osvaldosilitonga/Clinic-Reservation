package repositories

import (
	"clinic/models/entity"
	"context"

	"gorm.io/gorm"
)

type ClinicImpl struct {
	DB *gorm.DB
}

func NewClinicRepository(db *gorm.DB) Clinic {
	return &ClinicImpl{
		DB: db,
	}
}

func (cl *ClinicImpl) Save(c context.Context, d *entity.Clinics) (*entity.Clinics, error) {
	if err := cl.DB.WithContext(c).Create(d).Scan(&d).Error; err != nil {
		return nil, err
	}

	return d, nil
}

func (cl *ClinicImpl) FindByID(c context.Context, id int) (*entity.Clinics, error) {
	clinc := &entity.Clinics{}

	if err := cl.DB.WithContext(c).Where("id = ?", id).First(&clinc).Error; err != nil {
		return nil, err
	}

	return clinc, nil
}

func (cl *ClinicImpl) Update(c context.Context, d *entity.Clinics) (*entity.Clinics, error) {
	if err := cl.DB.WithContext(c).Save(d).Error; err != nil {
		return nil, err
	}

	return d, nil
}

func (cl *ClinicImpl) Delete(c context.Context, id int) error {
	if err := cl.DB.WithContext(c).Where("id = ?", id).Delete(&entity.Clinics{}).Error; err != nil {
		return err
	}

	return nil
}