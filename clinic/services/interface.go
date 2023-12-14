package services

import (
	"clinic/models/dto"
	"clinic/models/entity"
	"context"
)

type User interface {
	Register(ctx context.Context, d *dto.UserRegisterReq) (*dto.UserResponse, int, error)
	Login(ctx context.Context, d *dto.UserLoginReq) (*entity.Patients, int, error)
}

type Employee interface {
	Register(ctx context.Context, d *dto.EmployeeRegisterReq) (*entity.Employees, int, error)
	Login(ctx context.Context, d *dto.EmployeeLoginReq) (*entity.Employees, int, error)
}

type Clinic interface {
	Create(ctx context.Context, d *dto.ClinicReq) (*dto.ClinicRes, int, error)
	List(ctx context.Context) ([]dto.ClinicRes, int, error)
	FindByID(ctx context.Context, id int) (*dto.ClinicRes, int, error)
	Update(ctx context.Context, d *dto.ClinicUpdateReq, id int) (*dto.ClinicRes, int, error)
	Delete(ctx context.Context, id int) (int, error)
}

type Appointment interface {
	Create(ctx context.Context, d *entity.Appointments) (*dto.CreateAppointmentRes, int, error)
	Cancel(ctx context.Context, id int, email string) (*dto.AppointmentRes, int, error)
	Confirm(ctx context.Context, id int, email string, price int) (*dto.FullAppointmentRes, int, error)
	FindByPatientEmail(ctx context.Context, filter map[string]interface{}) ([]dto.FullAppointmentRes, int, error)
}
