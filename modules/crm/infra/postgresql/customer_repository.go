package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/manab-pr/evtaarpro/modules/crm/domain/entities"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Create(ctx context.Context, customer *entities.Customer) error {
	query := `
		INSERT INTO customers (id, company_name, contact_name, email, phone, address,
			industry, status, assigned_to, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	_, err := r.db.ExecContext(ctx, query,
		customer.ID, customer.CompanyName, customer.ContactName, customer.Email,
		customer.Phone, customer.Address, customer.Industry, customer.Status,
		customer.AssignedTo, customer.CreatedBy, customer.CreatedAt, customer.UpdatedAt,
	)
	return err
}

func (r *CustomerRepository) GetByID(ctx context.Context, id string) (*entities.Customer, error) {
	query := `
		SELECT id, company_name, contact_name, email, phone, address,
			industry, status, assigned_to, created_by, created_at, updated_at
		FROM customers WHERE id = $1
	`
	customer := &entities.Customer{}
	var email, phone, address, assignedTo sql.NullString

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&customer.ID, &customer.CompanyName, &customer.ContactName, &email,
		&phone, &address, &customer.Industry, &customer.Status,
		&assignedTo, &customer.CreatedBy, &customer.CreatedAt, &customer.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("customer not found")
	}
	if err != nil {
		return nil, err
	}

	if email.Valid {
		customer.Email = &email.String
	}
	if phone.Valid {
		customer.Phone = &phone.String
	}
	if address.Valid {
		customer.Address = &address.String
	}
	if assignedTo.Valid {
		customer.AssignedTo = &assignedTo.String
	}

	return customer, nil
}

func (r *CustomerRepository) List(ctx context.Context, status entities.CustomerStatus, assignedTo string, offset, limit int) ([]*entities.Customer, int, error) {
	var countQuery string
	var query string
	var args []interface{}

	baseCondition := `WHERE 1=1`
	if status != "" {
		baseCondition += ` AND status = $` + "1"
		args = append(args, status)
	}
	if assignedTo != "" {
		if status != "" {
			baseCondition += ` AND assigned_to = $2`
		} else {
			baseCondition += ` AND assigned_to = $1`
		}
		args = append(args, assignedTo)
	}

	countQuery = `SELECT COUNT(*) FROM customers ` + baseCondition

	argPos := len(args) + 1
	query = `
		SELECT id, company_name, contact_name, email, phone, address,
			industry, status, assigned_to, created_by, created_at, updated_at
		FROM customers ` + baseCondition + `
		ORDER BY created_at DESC
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

	customers := []*entities.Customer{}
	for rows.Next() {
		customer := &entities.Customer{}
		var email, phone, address, assignedTo sql.NullString

		err := rows.Scan(
			&customer.ID, &customer.CompanyName, &customer.ContactName, &email,
			&phone, &address, &customer.Industry, &customer.Status,
			&assignedTo, &customer.CreatedBy, &customer.CreatedAt, &customer.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		if email.Valid {
			customer.Email = &email.String
		}
		if phone.Valid {
			customer.Phone = &phone.String
		}
		if address.Valid {
			customer.Address = &address.String
		}
		if assignedTo.Valid {
			customer.AssignedTo = &assignedTo.String
		}

		customers = append(customers, customer)
	}

	return customers, total, nil
}

func (r *CustomerRepository) Update(ctx context.Context, customer *entities.Customer) error {
	query := `
		UPDATE customers SET
			company_name = $2, contact_name = $3, email = $4, phone = $5,
			address = $6, industry = $7, status = $8, assigned_to = $9, updated_at = $10
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		customer.ID, customer.CompanyName, customer.ContactName, customer.Email,
		customer.Phone, customer.Address, customer.Industry, customer.Status,
		customer.AssignedTo, customer.UpdatedAt,
	)
	return err
}

func (r *CustomerRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM customers WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
