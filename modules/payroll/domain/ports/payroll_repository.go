package ports

import (
	"context"
	"github.com/manab-pr/evtaarpro/modules/payroll/domain/entities"
)

// PayrollRepository defines payroll persistence operations
type PayrollRepository interface {
	// Employee operations
	CreateEmployee(ctx context.Context, employee *entities.Employee) error
	GetEmployee(ctx context.Context, id string) (*entities.Employee, error)
	ListEmployees(ctx context.Context, limit, offset int) ([]*entities.Employee, int, error)
	UpdateEmployee(ctx context.Context, employee *entities.Employee) error

	// Attendance operations
	CreateAttendance(ctx context.Context, attendance *entities.Attendance) error
	GetAttendance(ctx context.Context, employeeID string, month, year int) ([]*entities.Attendance, error)

	// Payroll operations
	CreatePayrollRecord(ctx context.Context, record *entities.PayrollRecord) error
	GetPayrollRecord(ctx context.Context, id string) (*entities.PayrollRecord, error)
	ListPayrollRecords(ctx context.Context, limit, offset int) ([]*entities.PayrollRecord, int, error)
	UpdatePayrollRecord(ctx context.Context, record *entities.PayrollRecord) error
}
