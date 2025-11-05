package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/manab-pr/evtaarpro/modules/payroll/domain/entities"
)

type EmployeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(ctx context.Context, employee *entities.Employee) error {
	query := `
		INSERT INTO employees (id, employee_code, department, designation, date_of_joining,
			salary_amount, salary_currency, bank_account, tax_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := r.db.ExecContext(ctx, query,
		employee.ID, employee.EmployeeCode, employee.Department, employee.Designation,
		employee.DateOfJoining, employee.SalaryAmount, employee.SalaryCurrency,
		employee.BankAccount, employee.TaxID, employee.CreatedAt, employee.UpdatedAt,
	)
	return err
}

func (r *EmployeeRepository) GetByID(ctx context.Context, id string) (*entities.Employee, error) {
	query := `
		SELECT id, employee_code, department, designation, date_of_joining,
			salary_amount, salary_currency, bank_account, tax_id, created_at, updated_at
		FROM employees WHERE id = $1
	`
	employee := &entities.Employee{}
	var bankAccount, taxID sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&employee.ID, &employee.EmployeeCode, &employee.Department, &employee.Designation,
		&employee.DateOfJoining, &employee.SalaryAmount, &employee.SalaryCurrency,
		&bankAccount, &taxID, &employee.CreatedAt, &employee.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("employee not found")
	}
	if err != nil {
		return nil, err
	}

	if bankAccount.Valid {
		employee.BankAccount = &bankAccount.String
	}
	if taxID.Valid {
		employee.TaxID = &taxID.String
	}

	return employee, nil
}

func (r *EmployeeRepository) GetByEmployeeCode(ctx context.Context, employeeCode string) (*entities.Employee, error) {
	query := `
		SELECT id, employee_code, department, designation, date_of_joining,
			salary_amount, salary_currency, bank_account, tax_id, created_at, updated_at
		FROM employees WHERE employee_code = $1
	`
	employee := &entities.Employee{}
	var bankAccount, taxID sql.NullString

	err := r.db.QueryRowContext(ctx, query, employeeCode).Scan(
		&employee.ID, &employee.EmployeeCode, &employee.Department, &employee.Designation,
		&employee.DateOfJoining, &employee.SalaryAmount, &employee.SalaryCurrency,
		&bankAccount, &taxID, &employee.CreatedAt, &employee.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("employee not found")
	}
	if err != nil {
		return nil, err
	}

	if bankAccount.Valid {
		employee.BankAccount = &bankAccount.String
	}
	if taxID.Valid {
		employee.TaxID = &taxID.String
	}

	return employee, nil
}

func (r *EmployeeRepository) List(ctx context.Context, department string, offset, limit int) ([]*entities.Employee, int, error) {
	var countQuery string
	var query string
	var args []interface{}

	if department != "" {
		countQuery = `SELECT COUNT(*) FROM employees WHERE department = $1`
		query = `
			SELECT id, employee_code, department, designation, date_of_joining,
				salary_amount, salary_currency, bank_account, tax_id, created_at, updated_at
			FROM employees WHERE department = $1
			ORDER BY created_at DESC
			LIMIT $2 OFFSET $3
		`
		args = []interface{}{department, limit, offset}
	} else {
		countQuery = `SELECT COUNT(*) FROM employees`
		query = `
			SELECT id, employee_code, department, designation, date_of_joining,
				salary_amount, salary_currency, bank_account, tax_id, created_at, updated_at
			FROM employees
			ORDER BY created_at DESC
			LIMIT $1 OFFSET $2
		`
		args = []interface{}{limit, offset}
	}

	var total int
	if department != "" {
		err := r.db.QueryRowContext(ctx, countQuery, department).Scan(&total)
		if err != nil {
			return nil, 0, err
		}
	} else {
		err := r.db.QueryRowContext(ctx, countQuery).Scan(&total)
		if err != nil {
			return nil, 0, err
		}
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	employees := []*entities.Employee{}
	for rows.Next() {
		employee := &entities.Employee{}
		var bankAccount, taxID sql.NullString

		err := rows.Scan(
			&employee.ID, &employee.EmployeeCode, &employee.Department, &employee.Designation,
			&employee.DateOfJoining, &employee.SalaryAmount, &employee.SalaryCurrency,
			&bankAccount, &taxID, &employee.CreatedAt, &employee.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		if bankAccount.Valid {
			employee.BankAccount = &bankAccount.String
		}
		if taxID.Valid {
			employee.TaxID = &taxID.String
		}

		employees = append(employees, employee)
	}

	return employees, total, nil
}

func (r *EmployeeRepository) Update(ctx context.Context, employee *entities.Employee) error {
	query := `
		UPDATE employees SET
			employee_code = $2, department = $3, designation = $4,
			date_of_joining = $5, salary_amount = $6, salary_currency = $7,
			bank_account = $8, tax_id = $9, updated_at = $10
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		employee.ID, employee.EmployeeCode, employee.Department, employee.Designation,
		employee.DateOfJoining, employee.SalaryAmount, employee.SalaryCurrency,
		employee.BankAccount, employee.TaxID, employee.UpdatedAt,
	)
	return err
}

func (r *EmployeeRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM employees WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
