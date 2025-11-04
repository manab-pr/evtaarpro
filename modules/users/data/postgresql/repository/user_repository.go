package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/manab-pr/evtaarpro/modules/users/domain/entities"
)

// UserRepository implements repository.UserRepository using PostgreSQL
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
		INSERT INTO users (id, email, first_name, last_name, phone, avatar, role, department, is_active, email_verified, created_at, updated_at, password_hash)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, '')
	`

	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Phone,
		user.Avatar,
		user.Role,
		user.Department,
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
		SELECT 
			id,
			email,
			first_name,
			last_name,
			COALESCE(phone, '') AS phone,
			COALESCE(avatar, '') AS avatar,
			role,
			COALESCE(department, '') AS department,
			is_active,
			email_verified,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
	`

	user := &entities.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.Avatar,
		&user.Role,
		&user.Department,
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
		SELECT 
			id,
			email,
			first_name,
			last_name,
			COALESCE(phone, '') AS phone,
			COALESCE(avatar, '') AS avatar,
			role,
			COALESCE(department, '') AS department,
			is_active,
			email_verified,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
	`

	user := &entities.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.Avatar,
		&user.Role,
		&user.Department,
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

// List retrieves users with pagination
func (r *UserRepository) List(ctx context.Context, page, pageSize int) ([]*entities.User, int64, error) {
	offset := (page - 1) * pageSize

	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM users`
	if err := r.db.QueryRowContext(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	// Get users
	query := `
		SELECT 
			id,
			email,
			first_name,
			last_name,
			COALESCE(phone, '') AS phone,
			COALESCE(avatar, '') AS avatar,
			role,
			COALESCE(department, '') AS department,
			is_active,
			email_verified,
			created_at,
			updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	users := make([]*entities.User, 0)
	for rows.Next() {
		user := &entities.User{}
		if err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Phone,
			&user.Avatar,
			&user.Role,
			&user.Department,
			&user.IsActive,
			&user.EmailVerified,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}

	return users, total, nil
}

// Update updates a user
func (r *UserRepository) Update(ctx context.Context, user *entities.User) error {
	query := `
		UPDATE users
		SET first_name = $2, last_name = $3, phone = $4, avatar = $5, department = $6, updated_at = $7
		WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Phone,
		user.Avatar,
		user.Department,
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

// Search searches users by name or email
func (r *UserRepository) Search(ctx context.Context, query string, page, pageSize int) ([]*entities.User, int64, error) {
	offset := (page - 1) * pageSize
	searchPattern := fmt.Sprintf("%%%s%%", query)

	// Get total count
	var total int64
	countQuery := `
		SELECT COUNT(*) FROM users
		WHERE first_name ILIKE $1 OR last_name ILIKE $1 OR email ILIKE $1
	`
	if err := r.db.QueryRowContext(ctx, countQuery, searchPattern).Scan(&total); err != nil {
		return nil, 0, err
	}

	// Get users
	searchQuery := `
		SELECT 
			id,
			email,
			first_name,
			last_name,
			COALESCE(phone, '') AS phone,
			COALESCE(avatar, '') AS avatar,
			role,
			COALESCE(department, '') AS department,
			is_active,
			email_verified,
			created_at,
			updated_at
		FROM users
		WHERE first_name ILIKE $1 OR last_name ILIKE $1 OR email ILIKE $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := r.db.QueryContext(ctx, searchQuery, searchPattern, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	users := make([]*entities.User, 0)
	for rows.Next() {
		user := &entities.User{}
		if err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Phone,
			&user.Avatar,
			&user.Role,
			&user.Department,
			&user.IsActive,
			&user.EmailVerified,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}

	return users, total, nil
}
