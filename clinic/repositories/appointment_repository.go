package repositories

import (
	"clinic/models/entity"
	"context"

	"gorm.io/gorm"
)

type AppointmentImpl struct {
	DB *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) Appointment {
	return &AppointmentImpl{
		DB: db,
	}
}

func (a *AppointmentImpl) Save(c context.Context, d *entity.Appointments) (*entity.Appointments, error) {
	err := a.DB.WithContext(c).Create(&d).Error
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (a *AppointmentImpl) FindByDate(c context.Context, date int64) ([]*entity.Appointments, error) {
	var appointments []*entity.Appointments

	err := a.DB.WithContext(c).Where("appointment_date = ?", date).Find(&appointments).Error
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (a *AppointmentImpl) FindWithFilter(c context.Context, filter *map[string]interface{}) ([]*entity.Appointments, error) {
	var appointments []*entity.Appointments

	err := a.DB.WithContext(c).Where(filter).Find(&appointments).Error
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (a *AppointmentImpl) CountRecordClinic(c context.Context, date int64, clinicID uint, status string) (int64, error) {
	var count int64

	err := a.DB.WithContext(c).Model(&entity.Appointments{}).Where("appointment_date = ? AND clinic_id = ? AND status = ?", date, clinicID, status).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
