package repositories

import (
	"clinic/models/entity"
	"context"

	"gorm.io/gorm"
)

type Patient interface {
	Save(c context.Context, d *entity.Patients) (*gorm.DB, error)
	FindByID(c context.Context, id int) (*entity.Patients, error)
	List(c context.Context) ([]*entity.Patients, error)
	Update(c context.Context, d *entity.Patients) (*entity.Patients, error)
	Delete(c context.Context, id int) error
}
