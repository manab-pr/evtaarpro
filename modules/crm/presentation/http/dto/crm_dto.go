package dto

import "time"

// Customer DTOs
type CreateCustomerRequest struct {
	CompanyName string  `json:"company_name" binding:"required"`
	ContactName string  `json:"contact_name"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	Address     *string `json:"address"`
	Industry    string  `json:"industry"`
	Status      string  `json:"status" binding:"required"`
	AssignedTo  *string `json:"assigned_to"`
}

type UpdateCustomerRequest struct {
	CompanyName string  `json:"company_name"`
	ContactName string  `json:"contact_name"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	Address     *string `json:"address"`
	Industry    string  `json:"industry"`
	Status      string  `json:"status"`
	AssignedTo  *string `json:"assigned_to"`
}

type CustomerResponse struct {
	ID          string     `json:"id"`
	CompanyName string     `json:"company_name"`
	ContactName string     `json:"contact_name"`
	Email       *string    `json:"email,omitempty"`
	Phone       *string    `json:"phone,omitempty"`
	Address     *string    `json:"address,omitempty"`
	Industry    string     `json:"industry"`
	Status      string     `json:"status"`
	AssignedTo  *string    `json:"assigned_to,omitempty"`
	CreatedBy   string     `json:"created_by"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// Customer Interaction DTOs
type CreateInteractionRequest struct {
	CustomerID      string     `json:"customer_id" binding:"required"`
	InteractionType string     `json:"interaction_type" binding:"required"`
	Subject         *string    `json:"subject"`
	Notes           *string    `json:"notes"`
	InteractionDate *time.Time `json:"interaction_date"`
}

type InteractionResponse struct {
	ID              string     `json:"id"`
	CustomerID      string     `json:"customer_id"`
	UserID          string     `json:"user_id"`
	InteractionType string     `json:"interaction_type"`
	Subject         *string    `json:"subject,omitempty"`
	Notes           *string    `json:"notes,omitempty"`
	InteractionDate time.Time  `json:"interaction_date"`
	CreatedAt       time.Time  `json:"created_at"`
}

type ListResponse struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}
