package postgresql

import (
	"context"
	"database/sql"

	"github.com/manab-pr/evtaarpro/modules/crm/domain/entities"
)

// CustomerRepository implements customer persistence
type CustomerRepository struct {
	db *sql.DB
}

// NewCustomerRepository creates a new repository
func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

// Create creates a new customer
func (r *CustomerRepository) Create(ctx context.Context, customer *entities.Customer) error {
	query := `
		INSERT INTO customers (id, company_id, name, email, phone, company, status, source, assigned_to, notes, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`
	_, err := r.db.ExecContext(ctx, query,
		customer.ID, customer.CompanyID, customer.Name, customer.Email, customer.Phone,
		customer.Company, customer.Status, customer.Source, customer.AssignedTo, customer.Notes,
		customer.CreatedBy, customer.CreatedAt, customer.UpdatedAt,
	)
	return err
}

// GetByID retrieves a customer by ID
func (r *CustomerRepository) GetByID(ctx context.Context, id string) (*entities.Customer, error) {
	query := `
		SELECT id, company_id, name, email, phone, company, status, source, assigned_to, notes, created_by, created_at, updated_at
		FROM customers WHERE id = $1
	`
	customer := &entities.Customer{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&customer.ID, &customer.CompanyID, &customer.Name, &customer.Email, &customer.Phone,
		&customer.Company, &customer.Status, &customer.Source, &customer.AssignedTo, &customer.Notes,
		&customer.CreatedBy, &customer.CreatedAt, &customer.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

// List retrieves customers with pagination
func (r *CustomerRepository) List(ctx context.Context, companyID string, limit, offset int) ([]*entities.Customer, int, error) {
	query := `
		SELECT id, company_id, name, email, phone, company, status, source, assigned_to, notes, created_by, created_at, updated_at
		FROM customers WHERE company_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := r.db.QueryContext(ctx, query, companyID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var customers []*entities.Customer
	for rows.Next() {
		customer := &entities.Customer{}
		err := rows.Scan(
			&customer.ID, &customer.CompanyID, &customer.Name, &customer.Email, &customer.Phone,
			&customer.Company, &customer.Status, &customer.Source, &customer.AssignedTo, &customer.Notes,
			&customer.CreatedBy, &customer.CreatedAt, &customer.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		customers = append(customers, customer)
	}

	// Get total count
	var total int
	countQuery := `SELECT COUNT(*) FROM customers WHERE company_id = $1`
	err = r.db.QueryRowContext(ctx, countQuery, companyID).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}

// Update updates a customer
func (r *CustomerRepository) Update(ctx context.Context, customer *entities.Customer) error {
	query := `
		UPDATE customers
		SET name = $2, email = $3, phone = $4, company = $5, status = $6, source = $7,
		    assigned_to = $8, notes = $9, updated_at = $10
		WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query,
		customer.ID, customer.Name, customer.Email, customer.Phone, customer.Company,
		customer.Status, customer.Source, customer.AssignedTo, customer.Notes, customer.UpdatedAt,
	)
	return err
}

// Delete deletes a customer
func (r *CustomerRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM customers WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// CreateInteraction creates a customer interaction
func (r *CustomerRepository) CreateInteraction(ctx context.Context, interaction *entities.CustomerInteraction) error {
	query := `
		INSERT INTO customer_interactions (id, customer_id, user_id, type, subject, description, scheduled_at, completed_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.ExecContext(ctx, query,
		interaction.ID, interaction.CustomerID, interaction.UserID, interaction.Type,
		interaction.Subject, interaction.Description, interaction.ScheduledAt,
		interaction.CompletedAt, interaction.CreatedAt,
	)
	return err
}

// GetInteractions retrieves interactions for a customer
func (r *CustomerRepository) GetInteractions(ctx context.Context, customerID string) ([]*entities.CustomerInteraction, error) {
	query := `
		SELECT id, customer_id, user_id, type, subject, description, scheduled_at, completed_at, created_at
		FROM customer_interactions WHERE customer_id = $1
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var interactions []*entities.CustomerInteraction
	for rows.Next() {
		interaction := &entities.CustomerInteraction{}
		err := rows.Scan(
			&interaction.ID, &interaction.CustomerID, &interaction.UserID, &interaction.Type,
			&interaction.Subject, &interaction.Description, &interaction.ScheduledAt,
			&interaction.CompletedAt, &interaction.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		interactions = append(interactions, interaction)
	}

	return interactions, nil
}
