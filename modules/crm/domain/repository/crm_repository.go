package repository

import (
	"context"

	"github.com/manab-pr/evtaarpro/modules/crm/domain/entities"
)

type CustomerRepository interface {
	Create(ctx context.Context, customer *entities.Customer) error
	GetByID(ctx context.Context, id string) (*entities.Customer, error)
	List(ctx context.Context, status entities.CustomerStatus, assignedTo string, offset, limit int) ([]*entities.Customer, int, error)
	Update(ctx context.Context, customer *entities.Customer) error
	Delete(ctx context.Context, id string) error
}

type CustomerInteractionRepository interface {
	Create(ctx context.Context, interaction *entities.CustomerInteraction) error
	GetByID(ctx context.Context, id string) (*entities.CustomerInteraction, error)
	ListByCustomer(ctx context.Context, customerID string, offset, limit int) ([]*entities.CustomerInteraction, int, error)
	Delete(ctx context.Context, id string) error
}
