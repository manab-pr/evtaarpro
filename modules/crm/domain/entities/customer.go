package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type CustomerStatus string

const (
	StatusLead     CustomerStatus = "lead"
	StatusProspect CustomerStatus = "prospect"
	StatusActive   CustomerStatus = "active"
	StatusInactive CustomerStatus = "inactive"
)

type Customer struct {
	ID          string
	CompanyName string
	ContactName string
	Email       *string
	Phone       *string
	Address     *string
	Industry    string
	Status      CustomerStatus
	AssignedTo  *string
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewCustomer(companyName, contactName, industry string, status CustomerStatus, createdBy string) (*Customer, error) {
	if companyName == "" {
		return nil, errors.New("company name is required")
	}

	return &Customer{
		ID:          uuid.New().String(),
		CompanyName: companyName,
		ContactName: contactName,
		Industry:    industry,
		Status:      status,
		CreatedBy:   createdBy,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (c *Customer) Assign(userID string) {
	c.AssignedTo = &userID
	c.UpdatedAt = time.Now()
}

func (c *Customer) UpdateStatus(status CustomerStatus) {
	c.Status = status
	c.UpdatedAt = time.Now()
}
