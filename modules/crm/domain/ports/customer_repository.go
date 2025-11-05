package ports

import (
	"context"
	"github.com/manab-pr/evtaarpro/modules/crm/domain/entities"
)

// CustomerRepository defines customer persistence operations
type CustomerRepository interface {
	Create(ctx context.Context, customer *entities.Customer) error
	GetByID(ctx context.Context, id string) (*entities.Customer, error)
	List(ctx context.Context, companyID string, limit, offset int) ([]*entities.Customer, int, error)
	Update(ctx context.Context, customer *entities.Customer) error
	Delete(ctx context.Context, id string) error

	// Interactions
	CreateInteraction(ctx context.Context, interaction *entities.CustomerInteraction) error
	GetInteractions(ctx context.Context, customerID string) ([]*entities.CustomerInteraction, error)
}
