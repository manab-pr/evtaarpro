package repository

import (
	"context"

	"github.com/manab-pr/evtaarpro/modules/users/domain/entities"
)

// UserRepository defines methods for user data access
type UserRepository interface {
	// Create creates a new user
	Create(ctx context.Context, user *entities.User) error

	// GetByID retrieves a user by ID
	GetByID(ctx context.Context, id string) (*entities.User, error)

	// GetByEmail retrieves a user by email
	GetByEmail(ctx context.Context, email string) (*entities.User, error)

	// List retrieves a list of users with pagination
	List(ctx context.Context, page, pageSize int) ([]*entities.User, int64, error)

	// Update updates a user
	Update(ctx context.Context, user *entities.User) error

	// Delete deletes a user
	Delete(ctx context.Context, id string) error

	// Search searches users by name or email
	Search(ctx context.Context, query string, page, pageSize int) ([]*entities.User, int64, error)
}
