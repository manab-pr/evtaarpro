package usecases

import (
	"context"

	"github.com/manab-pr/evtaarpro/modules/crm/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/crm/domain/ports"
)

// ListCustomersUseCase handles listing customers
type ListCustomersUseCase struct {
	customerRepo ports.CustomerRepository
}

// NewListCustomersUseCase creates a new use case
func NewListCustomersUseCase(customerRepo ports.CustomerRepository) *ListCustomersUseCase {
	return &ListCustomersUseCase{
		customerRepo: customerRepo,
	}
}

// Execute lists customers with pagination
func (uc *ListCustomersUseCase) Execute(ctx context.Context, companyID string, page, pageSize int) ([]*entities.Customer, int, error) {
	offset := (page - 1) * pageSize
	return uc.customerRepo.List(ctx, companyID, pageSize, offset)
}
