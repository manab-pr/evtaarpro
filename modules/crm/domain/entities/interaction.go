package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type InteractionType string

const (
	InteractionCall    InteractionType = "call"
	InteractionEmail   InteractionType = "email"
	InteractionMeeting InteractionType = "meeting"
	InteractionNote    InteractionType = "note"
)

type CustomerInteraction struct {
	ID              string
	CustomerID      string
	UserID          string
	InteractionType InteractionType
	Subject         *string
	Notes           *string
	InteractionDate time.Time
	CreatedAt       time.Time
}

func NewCustomerInteraction(customerID, userID string, interactionType InteractionType) (*CustomerInteraction, error) {
	if customerID == "" {
		return nil, errors.New("customer ID is required")
	}
	if userID == "" {
		return nil, errors.New("user ID is required")
	}

	return &CustomerInteraction{
		ID:              uuid.New().String(),
		CustomerID:      customerID,
		UserID:          userID,
		InteractionType: interactionType,
		InteractionDate: time.Now(),
		CreatedAt:       time.Now(),
	}, nil
}
