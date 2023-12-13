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

func (cl *ClinicImpl) List(c context.Context) ([]*entity.Clinics, error) {
	clincs := []*entity.Clinics{}

	if err := cl.DB.WithContext(c).Find(&clincs).Error; err != nil {
		return nil, err
	}

	return clincs, nil
}

func (cl *ClinicImpl) Update(c context.Context, d *entity.Clinics, id int) (*entity.Clinics, error) {
	if err := cl.DB.WithContext(c).Where("id = ?", id).Updates(&d).Scan(&d).Error; err != nil {
		return nil, err
	}

	// if err := cl.DB.WithContext(c).Save(d).Error; err != nil {
	// 	return nil, err
	// }

	return d, nil
}

func (cl *ClinicImpl) Delete(c context.Context, id int) (int, error) {
	res := cl.DB.WithContext(c).Where("id = ?", id).Delete(&entity.Clinics{})
	if res.Error != nil {
		return 0, res.Error
	}

	return int(res.RowsAffected), nil
}
