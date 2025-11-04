package ports

import (
	"context"

	"github.com/manab-pr/evtaarpro/modules/auth/domain/entities"
)

// UserRepository defines methods for user data access
type UserRepository interface {
	// Create creates a new user
	Create(ctx context.Context, user *entities.User) error

	// GetByID retrieves a user by ID
	GetByID(ctx context.Context, id string) (*entities.User, error)

	// GetByEmail retrieves a user by email
	GetByEmail(ctx context.Context, email string) (*entities.User, error)

	// Update updates a user
	Update(ctx context.Context, user *entities.User) error

	// Delete deletes a user
	Delete(ctx context.Context, id string) error

	// Exists checks if a user exists by email
	Exists(ctx context.Context, email string) (bool, error)
}
