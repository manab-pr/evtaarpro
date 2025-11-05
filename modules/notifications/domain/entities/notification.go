package entities

import "time"

// Notification represents a notification entity
type Notification struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Type      string    `json:"type"` // meeting_reminder, payroll_generated, customer_assigned, etc.
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Data      string    `json:"data"` // JSON data
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"created_at"`
	ReadAt    *time.Time `json:"read_at,omitempty"`
}
