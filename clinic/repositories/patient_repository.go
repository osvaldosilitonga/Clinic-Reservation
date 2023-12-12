package repositories

import (
	"clinic/models/entity"

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

func (p *PatientImpl) Save(d *entity.Patients) (*entity.Patients, error) {
	res := p.DB.Create(p).Scan(&d)
	if res.Error != nil {
		return nil, res.Error
	}

	return d, nil
}

func (p *PatientImpl) FindByID(id int) (*entity.Patients, error) {
	var d entity.Patients
	res := p.DB.Where("id = ?", id).First(&d)
	if res.Error != nil {
		return nil, res.Error
	}

	return &d, nil
}

func (p *PatientImpl) List() ([]*entity.Patients, error) {
	var d []*entity.Patients
	res := p.DB.Find(&d)
	if res.Error != nil {
		return nil, res.Error
	}

	return d, nil
}

func (p *PatientImpl) Update(d *entity.Patients) (*entity.Patients, error) {
	res := p.DB.Save(&d)
	if res.Error != nil {
		return nil, res.Error
	}

	return d, nil
}

func (p *PatientImpl) Delete(id int) error {
	res := p.DB.Delete(&entity.Patients{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
