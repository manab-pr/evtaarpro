package crm

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/config"
	"github.com/manab-pr/evtaarpro/internal/datastore"
	"github.com/manab-pr/evtaarpro/internal/middleware"
	"github.com/manab-pr/evtaarpro/modules/crm/infra/postgresql"
	"github.com/manab-pr/evtaarpro/modules/crm/presentation/http/handlers"
)

// RegisterRoutes registers CRM module routes
func RegisterRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	// Initialize repositories
	customerRepo := postgresql.NewCustomerRepository(pgStore.DB)
	interactionRepo := postgresql.NewCustomerInteractionRepository(pgStore.DB)

	// Initialize handlers
	crmHandlers := handlers.NewCRMHandlers(customerRepo, interactionRepo)

	// Register routes with auth middleware
	crm := rg.Group("/crm")
	crm.Use(middleware.AuthMiddleware(cfg.JWT.Secret))
	{
		// Customer routes
		crm.POST("/customers", crmHandlers.CreateCustomer)
		crm.GET("/customers", crmHandlers.ListCustomers)
		crm.GET("/customers/:customer_id", crmHandlers.GetCustomer)
		crm.PUT("/customers/:customer_id", crmHandlers.UpdateCustomer)
		crm.DELETE("/customers/:customer_id", crmHandlers.DeleteCustomer)
		crm.GET("/customers/:customer_id/interactions", crmHandlers.ListInteractionsByCustomer)

		// Interaction routes
		crm.POST("/interactions", crmHandlers.CreateInteraction)
		crm.GET("/interactions/:id", crmHandlers.GetInteraction)
	}
}
