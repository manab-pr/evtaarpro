package usecases

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/manab-pr/evtaarpro/modules/crm/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/crm/domain/ports"
)

// AddInteractionInput represents input for adding an interaction
type AddInteractionInput struct {
	CustomerID  string
	UserID      string
	Type        string
	Subject     string
	Description string
	ScheduledAt *time.Time
}

// AddInteractionUseCase handles adding customer interactions
type AddInteractionUseCase struct {
	customerRepo ports.CustomerRepository
}

// NewAddInteractionUseCase creates a new use case
func NewAddInteractionUseCase(customerRepo ports.CustomerRepository) *AddInteractionUseCase {
	return &AddInteractionUseCase{
		customerRepo: customerRepo,
	}
}

// Execute adds a new interaction
func (uc *AddInteractionUseCase) Execute(ctx context.Context, input AddInteractionInput) (*entities.CustomerInteraction, error) {
	interaction := &entities.CustomerInteraction{
		ID:          uuid.New().String(),
		CustomerID:  input.CustomerID,
		UserID:      input.UserID,
		Type:        input.Type,
		Subject:     input.Subject,
		Description: input.Description,
		ScheduledAt: input.ScheduledAt,
		CreatedAt:   time.Now(),
	}

	if err := uc.customerRepo.CreateInteraction(ctx, interaction); err != nil {
		return nil, err
	}

	return interaction, nil
}
