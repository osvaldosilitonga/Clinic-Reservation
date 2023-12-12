package repositories

import "clinic/models/entity"

type Patient interface {
	Save(d *entity.Patients) (*entity.Patients, error)
	FindByID(id int) (*entity.Patients, error)
	List() ([]*entity.Patients, error)
	Update(d *entity.Patients) (*entity.Patients, error)
	Delete(id int) error
}
