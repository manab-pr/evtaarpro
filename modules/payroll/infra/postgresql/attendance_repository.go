package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/manab-pr/evtaarpro/modules/payroll/domain/entities"
)

type AttendanceRepository struct {
	db *sql.DB
}

func NewAttendanceRepository(db *sql.DB) *AttendanceRepository {
	return &AttendanceRepository{db: db}
}

func (r *AttendanceRepository) Create(ctx context.Context, attendance *entities.Attendance) error {
	query := `
		INSERT INTO attendance (id, employee_id, date, check_in, check_out, status, notes, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := r.db.ExecContext(ctx, query,
		attendance.ID, attendance.EmployeeID, attendance.Date,
		attendance.CheckIn, attendance.CheckOut, attendance.Status,
		attendance.Notes, attendance.CreatedAt,
	)
	return err
}

func (r *AttendanceRepository) GetByID(ctx context.Context, id string) (*entities.Attendance, error) {
	query := `
		SELECT id, employee_id, date, check_in, check_out, status, notes, created_at
		FROM attendance WHERE id = $1
	`
	attendance := &entities.Attendance{}
	var checkIn, checkOut sql.NullTime
	var notes sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&attendance.ID, &attendance.EmployeeID, &attendance.Date,
		&checkIn, &checkOut, &attendance.Status, &notes, &attendance.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("attendance record not found")
	}
	if err != nil {
		return nil, err
	}

	if checkIn.Valid {
		attendance.CheckIn = &checkIn.Time
	}
	if checkOut.Valid {
		attendance.CheckOut = &checkOut.Time
	}
	if notes.Valid {
		attendance.Notes = &notes.String
	}

	return attendance, nil
}

func (r *AttendanceRepository) GetByEmployeeAndDate(ctx context.Context, employeeID string, date string) (*entities.Attendance, error) {
	query := `
		SELECT id, employee_id, date, check_in, check_out, status, notes, created_at
		FROM attendance WHERE employee_id = $1 AND date = $2
	`
	attendance := &entities.Attendance{}
	var checkIn, checkOut sql.NullTime
	var notes sql.NullString

	err := r.db.QueryRowContext(ctx, query, employeeID, date).Scan(
		&attendance.ID, &attendance.EmployeeID, &attendance.Date,
		&checkIn, &checkOut, &attendance.Status, &notes, &attendance.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("attendance record not found")
	}
	if err != nil {
		return nil, err
	}

	if checkIn.Valid {
		attendance.CheckIn = &checkIn.Time
	}
	if checkOut.Valid {
		attendance.CheckOut = &checkOut.Time
	}
	if notes.Valid {
		attendance.Notes = &notes.String
	}

	return attendance, nil
}

func (r *AttendanceRepository) List(ctx context.Context, employeeID string, startDate, endDate string, offset, limit int) ([]*entities.Attendance, int, error) {
	var countQuery string
	var query string
	var args []interface{}

	baseCondition := `WHERE 1=1`
	if employeeID != "" {
		baseCondition += ` AND employee_id = $1`
		args = append(args, employeeID)
	}
	if startDate != "" && endDate != "" {
		if employeeID != "" {
			baseCondition += ` AND date BETWEEN $2 AND $3`
			args = append(args, startDate, endDate)
		} else {
			baseCondition += ` AND date BETWEEN $1 AND $2`
			args = append(args, startDate, endDate)
		}
	}

	countQuery = `SELECT COUNT(*) FROM attendance ` + baseCondition

	argPos := len(args) + 1
	query = `
		SELECT id, employee_id, date, check_in, check_out, status, notes, created_at
		FROM attendance ` + baseCondition + `
		ORDER BY date DESC, created_at DESC
		LIMIT $` + string(rune(argPos+'0')) + ` OFFSET $` + string(rune(argPos+1+'0'))

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

	attendances := []*entities.Attendance{}
	for rows.Next() {
		attendance := &entities.Attendance{}
		var checkIn, checkOut sql.NullTime
		var notes sql.NullString

		err := rows.Scan(
			&attendance.ID, &attendance.EmployeeID, &attendance.Date,
			&checkIn, &checkOut, &attendance.Status, &notes, &attendance.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		if checkIn.Valid {
			attendance.CheckIn = &checkIn.Time
		}
		if checkOut.Valid {
			attendance.CheckOut = &checkOut.Time
		}
		if notes.Valid {
			attendance.Notes = &notes.String
		}

		attendances = append(attendances, attendance)
	}

	return attendances, total, nil
}

func (r *AttendanceRepository) Update(ctx context.Context, attendance *entities.Attendance) error {
	query := `
		UPDATE attendance SET
			check_in = $2, check_out = $3, status = $4, notes = $5
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		attendance.ID, attendance.CheckIn, attendance.CheckOut,
		attendance.Status, attendance.Notes,
	)
	return err
}

func (r *AttendanceRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM attendance WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
