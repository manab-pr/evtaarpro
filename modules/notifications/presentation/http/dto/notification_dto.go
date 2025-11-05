package dto

import "time"

// CreateNotificationRequest represents a notification creation request
type CreateNotificationRequest struct {
	UserID  string `json:"user_id" binding:"required"`
	Type    string `json:"type" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Message string `json:"message" binding:"required"`
	Data    string `json:"data"`
}

// NotificationResponse represents a notification response
type NotificationResponse struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	Type      string     `json:"type"`
	Title     string     `json:"title"`
	Message   string     `json:"message"`
	Data      string     `json:"data"`
	Read      bool       `json:"read"`
	CreatedAt time.Time  `json:"created_at"`
	ReadAt    *time.Time `json:"read_at,omitempty"`
}
