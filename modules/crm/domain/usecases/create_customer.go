package usecases

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/manab-pr/evtaarpro/modules/crm/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/crm/domain/ports"
)

// CreateCustomerInput represents input for creating a customer
type CreateCustomerInput struct {
	CompanyID  string
	Name       string
	Email      string
	Phone      string
	Company    string
	Status     entities.CustomerStatus
	Source     string
	AssignedTo *string
	Notes      string
	CreatedBy  string
}

// CreateCustomerUseCase handles customer creation
type CreateCustomerUseCase struct {
	customerRepo ports.CustomerRepository
}

// NewCreateCustomerUseCase creates a new use case
func NewCreateCustomerUseCase(customerRepo ports.CustomerRepository) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		customerRepo: customerRepo,
	}
}

// Execute creates a new customer
func (uc *CreateCustomerUseCase) Execute(ctx context.Context, input CreateCustomerInput) (*entities.Customer, error) {
	customer := &entities.Customer{
		ID:         uuid.New().String(),
		CompanyID:  input.CompanyID,
		Name:       input.Name,
		Email:      input.Email,
		Phone:      input.Phone,
		Company:    input.Company,
		Status:     input.Status,
		Source:     input.Source,
		AssignedTo: input.AssignedTo,
		Notes:      input.Notes,
		CreatedBy:  input.CreatedBy,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := uc.customerRepo.Create(ctx, customer); err != nil {
		return nil, err
	}

	return customer, nil
}
