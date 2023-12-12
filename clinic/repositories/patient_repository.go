package repositories

import (
	"clinic/models/entity"
	"context"

	"gorm.io/gorm"
)

type PatientImpl struct {
	DB *gorm.DB
}

func NewPatientRepository(db *gorm.DB) Patient {
	return &PatientImpl{
		DB: db,
	}
}

func (p *PatientImpl) Save(c context.Context, d *entity.Patients) (*gorm.DB, error) {
	res := p.DB.WithContext(c).Create(&d)
	if res.Error != nil {
		return res, res.Error
	}

	return res, nil
}

func (p *PatientImpl) FindByID(c context.Context, id int) (*entity.Patients, error) {
	var d entity.Patients
	res := p.DB.WithContext(c).Where("id = ?", id).First(&d)
	if res.Error != nil {
		return nil, res.Error
	}

	return &d, nil
}

func (p *PatientImpl) List(c context.Context) ([]*entity.Patients, error) {
	var d []*entity.Patients
	res := p.DB.WithContext(c).Find(&d)
	if res.Error != nil {
		return nil, res.Error
	}

	return d, nil
}

func (p *PatientImpl) Update(c context.Context, d *entity.Patients) (*entity.Patients, error) {
	res := p.DB.WithContext(c).Save(&d)
	if res.Error != nil {
		return nil, res.Error
	}

	return d, nil
}

func (p *PatientImpl) Delete(c context.Context, id int) error {
	res := p.DB.WithContext(c).Delete(&entity.Patients{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
