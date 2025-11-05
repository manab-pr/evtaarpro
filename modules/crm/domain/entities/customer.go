package entities

import "time"

// CustomerStatus represents the status of a customer
type CustomerStatus string

const (
	CustomerStatusLead       CustomerStatus = "lead"
	CustomerStatusProspect   CustomerStatus = "prospect"
	CustomerStatusActive     CustomerStatus = "active"
	CustomerStatusInactive   CustomerStatus = "inactive"
	CustomerStatusChurned    CustomerStatus = "churned"
)

// Customer represents a customer entity
type Customer struct {
	ID          string         `json:"id"`
	CompanyID   string         `json:"company_id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Phone       string         `json:"phone"`
	Company     string         `json:"company"`
	Status      CustomerStatus `json:"status"`
	Source      string         `json:"source"`
	AssignedTo  *string        `json:"assigned_to,omitempty"`
	Notes       string         `json:"notes"`
	CreatedBy   string         `json:"created_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// CustomerInteraction represents an interaction with a customer
type CustomerInteraction struct {
	ID           string    `json:"id"`
	CustomerID   string    `json:"customer_id"`
	UserID       string    `json:"user_id"`
	Type         string    `json:"type"` // call, email, meeting, note
	Subject      string    `json:"subject"`
	Description  string    `json:"description"`
	ScheduledAt  *time.Time `json:"scheduled_at,omitempty"`
	CompletedAt  *time.Time `json:"completed_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
}
