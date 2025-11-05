package dto

import "time"

type CreateNotificationRequest struct {
	UserID   string                 `json:"user_id" binding:"required"`
	Type     string                 `json:"type" binding:"required"`
	Title    string                 `json:"title" binding:"required"`
	Message  string                 `json:"message" binding:"required"`
	Data     map[string]interface{} `json:"data"`
	Priority string                 `json:"priority"`
}

type NotificationResponse struct {
	ID        string                 `json:"id"`
	UserID    string                 `json:"user_id"`
	Type      string                 `json:"type"`
	Title     string                 `json:"title"`
	Message   string                 `json:"message"`
	Data      map[string]interface{} `json:"data,omitempty"`
	IsRead    bool                   `json:"is_read"`
	ReadAt    *time.Time             `json:"read_at,omitempty"`
	Priority  string                 `json:"priority"`
	CreatedAt time.Time              `json:"created_at"`
}

type ListResponse struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}

type UnreadCountResponse struct {
	Count int `json:"count"`
}
