package repositories

import (
	"clinic/models/entity"
	"context"

	"gorm.io/gorm"
)

type Patient interface {
	Save(c context.Context, d *entity.Patients) (*gorm.DB, error)
	FindByID(c context.Context, id int) (*entity.Patients, error)
	FindByEmail(c context.Context, email string) (*entity.Patients, error)
	List(c context.Context) ([]*entity.Patients, error)
	Update(c context.Context, d *entity.Patients) (*entity.Patients, error)
	Delete(c context.Context, id int) error
}

type Employee interface {
	Save(c context.Context, d *entity.Employees) (*entity.Employees, error)
	FindByUsername(c context.Context, username string) (*entity.Employees, error)
}

type Clinic interface {
	Save(c context.Context, d *entity.Clinics) (*entity.Clinics, error)
	FindByID(c context.Context, id int) (*entity.Clinics, error)
	List(c context.Context) ([]*entity.Clinics, error)
	Update(c context.Context, d *entity.Clinics) (*entity.Clinics, error)
	Delete(c context.Context, id int) error
}
