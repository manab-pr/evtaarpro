package crm

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/config"
	"github.com/manab-pr/evtaarpro/internal/datastore"
	"github.com/manab-pr/evtaarpro/modules/crm/domain/usecases"
	"github.com/manab-pr/evtaarpro/modules/crm/infra/postgresql"
	"github.com/manab-pr/evtaarpro/modules/crm/presentation/http/handlers"
	"github.com/manab-pr/evtaarpro/modules/crm/presentation/http/routes"
)

// RegisterRoutes registers CRM module routes
func RegisterRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	// Infrastructure
	customerRepo := postgresql.NewCustomerRepository(pgStore.DB)

	// Use cases
	createCustomerUC := usecases.NewCreateCustomerUseCase(customerRepo)
	listCustomersUC := usecases.NewListCustomersUseCase(customerRepo)
	addInteractionUC := usecases.NewAddInteractionUseCase(customerRepo)

	// Handlers
	customerHandlers := handlers.NewCustomerHandlers(
		createCustomerUC,
		listCustomersUC,
		addInteractionUC,
		customerRepo,
	)

	// Register routes
	routes.RegisterRoutes(rg, customerHandlers, cfg.JWT.Secret)
}
