package repository

import (
	"context"

	"github.com/manab-pr/evtaarpro/modules/meetings/domain/entities"
)

// MeetingRepository defines methods for meeting data access
type MeetingRepository interface {
	// Create creates a new meeting
	Create(ctx context.Context, meeting *entities.Meeting) error

	// GetByID retrieves a meeting by ID
	GetByID(ctx context.Context, id string) (*entities.Meeting, error)

	// List retrieves meetings with pagination
	List(ctx context.Context, page, pageSize int, userID string) ([]*entities.Meeting, int64, error)

	// ListByOrganizer retrieves meetings by organizer
	ListByOrganizer(ctx context.Context, organizerID string, page, pageSize int) ([]*entities.Meeting, int64, error)

	// Update updates a meeting
	Update(ctx context.Context, meeting *entities.Meeting) error

	// Delete deletes a meeting
	Delete(ctx context.Context, id string) error

	// GetUpcoming retrieves upcoming meetings
	GetUpcoming(ctx context.Context, userID string, limit int) ([]*entities.Meeting, error)
}
