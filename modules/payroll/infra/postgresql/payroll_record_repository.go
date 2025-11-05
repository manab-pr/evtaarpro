package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/manab-pr/evtaarpro/modules/payroll/domain/entities"
)

type PayrollRecordRepository struct {
	db *sql.DB
}

func NewPayrollRecordRepository(db *sql.DB) *PayrollRecordRepository {
	return &PayrollRecordRepository{db: db}
}

func (r *PayrollRecordRepository) Create(ctx context.Context, record *entities.PayrollRecord) error {
	query := `
		INSERT INTO payroll_records (id, employee_id, period_start, period_end,
			gross_salary, deductions, net_salary, status, paid_on, payment_reference,
			created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`
	_, err := r.db.ExecContext(ctx, query,
		record.ID, record.EmployeeID, record.PeriodStart, record.PeriodEnd,
		record.GrossSalary, record.Deductions, record.NetSalary, record.Status,
		record.PaidOn, record.PaymentReference, record.CreatedBy,
		record.CreatedAt, record.UpdatedAt,
	)
	return err
}

func (r *PayrollRecordRepository) GetByID(ctx context.Context, id string) (*entities.PayrollRecord, error) {
	query := `
		SELECT id, employee_id, period_start, period_end, gross_salary, deductions,
			net_salary, status, paid_on, payment_reference, created_by, created_at, updated_at
		FROM payroll_records WHERE id = $1
	`
	record := &entities.PayrollRecord{}
	var paidOn sql.NullTime
	var paymentReference sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&record.ID, &record.EmployeeID, &record.PeriodStart, &record.PeriodEnd,
		&record.GrossSalary, &record.Deductions, &record.NetSalary, &record.Status,
		&paidOn, &paymentReference, &record.CreatedBy,
		&record.CreatedAt, &record.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("payroll record not found")
	}
	if err != nil {
		return nil, err
	}

	if paidOn.Valid {
		record.PaidOn = &paidOn.Time
	}
	if paymentReference.Valid {
		record.PaymentReference = &paymentReference.String
	}

	return record, nil
}

func (r *PayrollRecordRepository) List(ctx context.Context, employeeID string, status entities.PayrollStatus, offset, limit int) ([]*entities.PayrollRecord, int, error) {
	var countQuery string
	var query string
	var args []interface{}
	argCount := 1

	countQueryBase := `SELECT COUNT(*) FROM payroll_records WHERE 1=1`
	queryBase := `
		SELECT id, employee_id, period_start, period_end, gross_salary, deductions,
			net_salary, status, paid_on, payment_reference, created_by, created_at, updated_at
		FROM payroll_records WHERE 1=1
	`

	if employeeID != "" {
		countQueryBase += ` AND employee_id = $` + string(rune(argCount+'0'))
		queryBase += ` AND employee_id = $` + string(rune(argCount+'0'))
		args = append(args, employeeID)
		argCount++
	}

	if status != "" {
		countQueryBase += ` AND status = $` + string(rune(argCount+'0'))
		queryBase += ` AND status = $` + string(rune(argCount+'0'))
		args = append(args, status)
		argCount++
	}

	countQuery = countQueryBase
	query = queryBase + ` ORDER BY created_at DESC LIMIT $` + string(rune(argCount+'0')) + ` OFFSET $` + string(rune(argCount+1+'0'))

	countArgs := args
	args = append(args, limit, offset)

	var total int
	err := r.db.QueryRowContext(ctx, countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	records := []*entities.PayrollRecord{}
	for rows.Next() {
		record := &entities.PayrollRecord{}
		var paidOn sql.NullTime
		var paymentReference sql.NullString

		err := rows.Scan(
			&record.ID, &record.EmployeeID, &record.PeriodStart, &record.PeriodEnd,
			&record.GrossSalary, &record.Deductions, &record.NetSalary, &record.Status,
			&paidOn, &paymentReference, &record.CreatedBy,
			&record.CreatedAt, &record.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		if paidOn.Valid {
			record.PaidOn = &paidOn.Time
		}
		if paymentReference.Valid {
			record.PaymentReference = &paymentReference.String
		}

		records = append(records, record)
	}

	return records, total, nil
}

func (r *PayrollRecordRepository) Update(ctx context.Context, record *entities.PayrollRecord) error {
	query := `
		UPDATE payroll_records SET
			employee_id = $2, period_start = $3, period_end = $4,
			gross_salary = $5, deductions = $6, net_salary = $7, status = $8,
			paid_on = $9, payment_reference = $10, updated_at = $11
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		record.ID, record.EmployeeID, record.PeriodStart, record.PeriodEnd,
		record.GrossSalary, record.Deductions, record.NetSalary, record.Status,
		record.PaidOn, record.PaymentReference, record.UpdatedAt,
	)
	return err
}

func (r *PayrollRecordRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM payroll_records WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
