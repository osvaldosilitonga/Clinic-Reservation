package repositories

import (
	"clinic/models/entity"
	"context"

	"gorm.io/gorm"
)

type EmployeeImpl struct {
	DB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) Employee {
	return &EmployeeImpl{
		DB: db,
	}
}

func (e *EmployeeImpl) Save(c context.Context, d *entity.Employees) (*entity.Employees, error) {
	if err := e.DB.WithContext(c).Create(&d).Scan(&d).Error; err != nil {
		return nil, err
	}

	return d, nil
}

func (e *EmployeeImpl) FindByUsername(c context.Context, username string) (*entity.Employees, error) {
	employee := entity.Employees{}
	res := e.DB.WithContext(c).Where("username = ?", username).First(&employee)
	if res.Error != nil {
		return nil, res.Error
	}

	return &employee, res.Error
}
