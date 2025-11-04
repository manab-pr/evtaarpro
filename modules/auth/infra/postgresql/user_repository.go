package postgresql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/manab-pr/evtaarpro/modules/auth/domain/entities"
)

// UserRepository implements ports.UserRepository using PostgreSQL
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create creates a new user
func (r *UserRepository) Create(ctx context.Context, user *entities.User) error {
	query := `
		INSERT INTO users (id, email, password_hash, first_name, last_name, role, is_active, email_verified, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.Role,
		user.IsActive,
		user.EmailVerified,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return err
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id string) (*entities.User, error) {
	query := `
		SELECT id, email, password_hash, first_name, last_name, role, is_active, email_verified, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &entities.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.IsActive,
		&user.EmailVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

// GetByEmail retrieves a user by email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := `
		SELECT id, email, password_hash, first_name, last_name, role, is_active, email_verified, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	user := &entities.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Role,
		&user.IsActive,
		&user.EmailVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

// Update updates a user
func (r *UserRepository) Update(ctx context.Context, user *entities.User) error {
	query := `
		UPDATE users
		SET email = $2, first_name = $3, last_name = $4, role = $5, is_active = $6, email_verified = $7, updated_at = $8
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Role,
		user.IsActive,
		user.EmailVerified,
		user.UpdatedAt,
	)

	return err
}

// Delete deletes a user
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// Exists checks if a user exists by email
func (r *UserRepository) Exists(ctx context.Context, email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	var exists bool
	err := r.db.QueryRowContext(ctx, query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
