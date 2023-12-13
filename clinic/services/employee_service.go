package services

import (
	"clinic/models/dto"
	"clinic/models/entity"
	"clinic/repositories"
	"context"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type EmployeeImpl struct {
	EmployeeRepo repositories.Employee
}

func NewEmployeeService(employeeRepo repositories.Employee) Employee {
	return &EmployeeImpl{
		EmployeeRepo: employeeRepo,
	}
}

func (e *EmployeeImpl) Register(ctx context.Context, d *dto.EmployeeRegisterReq) (*entity.Employees, int, error) {
	employee := entity.Employees{
		Username: d.Username,
		Role:     "admin",
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(d.Password), 14)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	employee.Password = string(hash)

	res, err := e.EmployeeRepo.Save(ctx, &employee)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, http.StatusBadRequest, err
		}
		return nil, http.StatusInternalServerError, err
	}

	return res, http.StatusCreated, nil
}
