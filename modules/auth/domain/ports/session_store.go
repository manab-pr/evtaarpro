package ports

import (
	"context"
	"time"
)

// SessionStore defines methods for session management
type SessionStore interface {
	// Create creates a new session
	Create(ctx context.Context, userID string, token string, ttl time.Duration) error

	// Get retrieves a session
	Get(ctx context.Context, userID string) (string, error)

	// Delete deletes a session
	Delete(ctx context.Context, userID string) error

	// Exists checks if a session exists
	Exists(ctx context.Context, userID string) (bool, error)
}
