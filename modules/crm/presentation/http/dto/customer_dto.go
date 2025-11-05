package dto

import "time"

// CreateCustomerRequest represents a customer creation request
type CreateCustomerRequest struct {
	Name       string  `json:"name" binding:"required"`
	Email      string  `json:"email" binding:"required,email"`
	Phone      string  `json:"phone"`
	Company    string  `json:"company"`
	Status     string  `json:"status" binding:"required"`
	Source     string  `json:"source"`
	AssignedTo *string `json:"assigned_to"`
	Notes      string  `json:"notes"`
}

// UpdateCustomerRequest represents a customer update request
type UpdateCustomerRequest struct {
	Name       string  `json:"name"`
	Email      string  `json:"email" binding:"email"`
	Phone      string  `json:"phone"`
	Company    string  `json:"company"`
	Status     string  `json:"status"`
	Source     string  `json:"source"`
	AssignedTo *string `json:"assigned_to"`
	Notes      string  `json:"notes"`
}

// CustomerResponse represents a customer response
type CustomerResponse struct {
	ID         string     `json:"id"`
	CompanyID  string     `json:"company_id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	Company    string     `json:"company"`
	Status     string     `json:"status"`
	Source     string     `json:"source"`
	AssignedTo *string    `json:"assigned_to,omitempty"`
	Notes      string     `json:"notes"`
	CreatedBy  string     `json:"created_by"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// CreateInteractionRequest represents an interaction creation request
type CreateInteractionRequest struct {
	Type        string     `json:"type" binding:"required"`
	Subject     string     `json:"subject" binding:"required"`
	Description string     `json:"description"`
	ScheduledAt *time.Time `json:"scheduled_at"`
}

// InteractionResponse represents an interaction response
type InteractionResponse struct {
	ID          string     `json:"id"`
	CustomerID  string     `json:"customer_id"`
	UserID      string     `json:"user_id"`
	Type        string     `json:"type"`
	Subject     string     `json:"subject"`
	Description string     `json:"description"`
	ScheduledAt *time.Time `json:"scheduled_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}
