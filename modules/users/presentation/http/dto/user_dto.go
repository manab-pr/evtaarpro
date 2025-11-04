package dto

import "time"

// UserResponse represents a user response
type UserResponse struct {
	ID            string    `json:"id"`
	Email         string    `json:"email"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	FullName      string    `json:"full_name"`
	Phone         string    `json:"phone,omitempty"`
	Avatar        string    `json:"avatar,omitempty"`
	Role          string    `json:"role"`
	Department    string    `json:"department,omitempty"`
	IsActive      bool      `json:"is_active"`
	EmailVerified bool      `json:"email_verified"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// UpdateUserRequest represents a user update request
type UpdateUserRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Department string `json:"department"`
}

// ListUsersQuery represents list users query parameters
type ListUsersQuery struct {
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
	Search   string `form:"search"`
}
