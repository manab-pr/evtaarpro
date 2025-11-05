package repository

import (
	"context"

	"github.com/manab-pr/evtaarpro/modules/payroll/domain/entities"
)

type EmployeeRepository interface {
	Create(ctx context.Context, employee *entities.Employee) error
	GetByID(ctx context.Context, id string) (*entities.Employee, error)
	GetByEmployeeCode(ctx context.Context, employeeCode string) (*entities.Employee, error)
	List(ctx context.Context, department string, offset, limit int) ([]*entities.Employee, int, error)
	Update(ctx context.Context, employee *entities.Employee) error
	Delete(ctx context.Context, id string) error
}

type PayrollRecordRepository interface {
	Create(ctx context.Context, record *entities.PayrollRecord) error
	GetByID(ctx context.Context, id string) (*entities.PayrollRecord, error)
	List(ctx context.Context, employeeID string, status entities.PayrollStatus, offset, limit int) ([]*entities.PayrollRecord, int, error)
	Update(ctx context.Context, record *entities.PayrollRecord) error
	Delete(ctx context.Context, id string) error
}

type AttendanceRepository interface {
	Create(ctx context.Context, attendance *entities.Attendance) error
	GetByID(ctx context.Context, id string) (*entities.Attendance, error)
	GetByEmployeeAndDate(ctx context.Context, employeeID string, date string) (*entities.Attendance, error)
	List(ctx context.Context, employeeID string, startDate, endDate string, offset, limit int) ([]*entities.Attendance, int, error)
	Update(ctx context.Context, attendance *entities.Attendance) error
	Delete(ctx context.Context, id string) error
}
