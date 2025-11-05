package postgresql

import (
	"context"
	"database/sql"

	"github.com/manab-pr/evtaarpro/modules/payroll/domain/entities"
)

// PayrollRepository implements payroll persistence
type PayrollRepository struct {
	db *sql.DB
}

// NewPayrollRepository creates a new repository
func NewPayrollRepository(db *sql.DB) *PayrollRepository {
	return &PayrollRepository{db: db}
}

// CreateEmployee creates a new employee
func (r *PayrollRepository) CreateEmployee(ctx context.Context, employee *entities.Employee) error {
	query := `
		INSERT INTO employees (id, user_id, employee_code, department, designation, joining_date, salary_amount, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := r.db.ExecContext(ctx, query,
		employee.ID, employee.UserID, employee.EmployeeCode, employee.Department,
		employee.Designation, employee.JoiningDate, employee.SalaryAmount,
		employee.IsActive, employee.CreatedAt, employee.UpdatedAt,
	)
	return err
}

// GetEmployee retrieves an employee by ID
func (r *PayrollRepository) GetEmployee(ctx context.Context, id string) (*entities.Employee, error) {
	query := `
		SELECT id, user_id, employee_code, department, designation, joining_date, salary_amount, is_active, created_at, updated_at
		FROM employees WHERE id = $1
	`
	employee := &entities.Employee{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&employee.ID, &employee.UserID, &employee.EmployeeCode, &employee.Department,
		&employee.Designation, &employee.JoiningDate, &employee.SalaryAmount,
		&employee.IsActive, &employee.CreatedAt, &employee.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

// ListEmployees retrieves employees with pagination
func (r *PayrollRepository) ListEmployees(ctx context.Context, limit, offset int) ([]*entities.Employee, int, error) {
	query := `
		SELECT id, user_id, employee_code, department, designation, joining_date, salary_amount, is_active, created_at, updated_at
		FROM employees
		WHERE is_active = true
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var employees []*entities.Employee
	for rows.Next() {
		employee := &entities.Employee{}
		err := rows.Scan(
			&employee.ID, &employee.UserID, &employee.EmployeeCode, &employee.Department,
			&employee.Designation, &employee.JoiningDate, &employee.SalaryAmount,
			&employee.IsActive, &employee.CreatedAt, &employee.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		employees = append(employees, employee)
	}

	// Get total count
	var total int
	countQuery := `SELECT COUNT(*) FROM employees WHERE is_active = true`
	err = r.db.QueryRowContext(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return employees, total, nil
}

// UpdateEmployee updates an employee
func (r *PayrollRepository) UpdateEmployee(ctx context.Context, employee *entities.Employee) error {
	query := `
		UPDATE employees
		SET department = $2, designation = $3, salary_amount = $4, is_active = $5, updated_at = $6
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		employee.ID, employee.Department, employee.Designation,
		employee.SalaryAmount, employee.IsActive, employee.UpdatedAt,
	)
	return err
}

// CreateAttendance creates an attendance record
func (r *PayrollRepository) CreateAttendance(ctx context.Context, attendance *entities.Attendance) error {
	query := `
		INSERT INTO attendance (id, employee_id, date, check_in, check_out, status, hours_worked, notes, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.ExecContext(ctx, query,
		attendance.ID, attendance.EmployeeID, attendance.Date, attendance.CheckIn,
		attendance.CheckOut, attendance.Status, attendance.HoursWorked,
		attendance.Notes, attendance.CreatedAt,
	)
	return err
}

// GetAttendance retrieves attendance records
func (r *PayrollRepository) GetAttendance(ctx context.Context, employeeID string, month, year int) ([]*entities.Attendance, error) {
	query := `
		SELECT id, employee_id, date, check_in, check_out, status, hours_worked, notes, created_at
		FROM attendance
		WHERE employee_id = $1 AND EXTRACT(MONTH FROM date) = $2 AND EXTRACT(YEAR FROM date) = $3
		ORDER BY date DESC
	`
	rows, err := r.db.QueryContext(ctx, query, employeeID, month, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []*entities.Attendance
	for rows.Next() {
		record := &entities.Attendance{}
		err := rows.Scan(
			&record.ID, &record.EmployeeID, &record.Date, &record.CheckIn,
			&record.CheckOut, &record.Status, &record.HoursWorked,
			&record.Notes, &record.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

// CreatePayrollRecord creates a payroll record
func (r *PayrollRepository) CreatePayrollRecord(ctx context.Context, record *entities.PayrollRecord) error {
	query := `
		INSERT INTO payroll_records (id, employee_id, month, year, basic_salary, allowances, deductions, net_salary, payment_status, payment_date, notes, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`
	_, err := r.db.ExecContext(ctx, query,
		record.ID, record.EmployeeID, record.Month, record.Year, record.BasicSalary,
		record.Allowances, record.Deductions, record.NetSalary, record.PaymentStatus,
		record.PaymentDate, record.Notes, record.CreatedAt, record.UpdatedAt,
	)
	return err
}

// GetPayrollRecord retrieves a payroll record by ID
func (r *PayrollRepository) GetPayrollRecord(ctx context.Context, id string) (*entities.PayrollRecord, error) {
	query := `
		SELECT id, employee_id, month, year, basic_salary, allowances, deductions, net_salary, payment_status, payment_date, notes, created_at, updated_at
		FROM payroll_records WHERE id = $1
	`
	record := &entities.PayrollRecord{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&record.ID, &record.EmployeeID, &record.Month, &record.Year, &record.BasicSalary,
		&record.Allowances, &record.Deductions, &record.NetSalary, &record.PaymentStatus,
		&record.PaymentDate, &record.Notes, &record.CreatedAt, &record.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// ListPayrollRecords retrieves payroll records with pagination
func (r *PayrollRepository) ListPayrollRecords(ctx context.Context, limit, offset int) ([]*entities.PayrollRecord, int, error) {
	query := `
		SELECT id, employee_id, month, year, basic_salary, allowances, deductions, net_salary, payment_status, payment_date, notes, created_at, updated_at
		FROM payroll_records
		ORDER BY year DESC, month DESC, created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var records []*entities.PayrollRecord
	for rows.Next() {
		record := &entities.PayrollRecord{}
		err := rows.Scan(
			&record.ID, &record.EmployeeID, &record.Month, &record.Year, &record.BasicSalary,
			&record.Allowances, &record.Deductions, &record.NetSalary, &record.PaymentStatus,
			&record.PaymentDate, &record.Notes, &record.CreatedAt, &record.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		records = append(records, record)
	}

	// Get total count
	var total int
	countQuery := `SELECT COUNT(*) FROM payroll_records`
	err = r.db.QueryRowContext(ctx, countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

// UpdatePayrollRecord updates a payroll record
func (r *PayrollRepository) UpdatePayrollRecord(ctx context.Context, record *entities.PayrollRecord) error {
	query := `
		UPDATE payroll_records
		SET payment_status = $2, payment_date = $3, notes = $4, updated_at = $5
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		record.ID, record.PaymentStatus, record.PaymentDate, record.Notes, record.UpdatedAt,
	)
	return err
}
