package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/middleware"
	"github.com/manab-pr/evtaarpro/modules/crm/presentation/http/handlers"
)

// RegisterRoutes registers CRM routes
func RegisterRoutes(rg *gin.RouterGroup, customerHandlers *handlers.CustomerHandlers, jwtSecret string) {
	crm := rg.Group("/crm")
	crm.Use(middleware.AuthMiddleware(jwtSecret))
	{
		// Customer routes
		crm.POST("/customers", customerHandlers.CreateCustomer)
		crm.GET("/customers", customerHandlers.ListCustomers)
		crm.GET("/customers/:id", customerHandlers.GetCustomer)
		crm.PUT("/customers/:id", customerHandlers.UpdateCustomer)
		crm.DELETE("/customers/:id", customerHandlers.DeleteCustomer)

		// Interaction routes
		crm.POST("/customers/:id/interactions", customerHandlers.AddInteraction)
		crm.GET("/customers/:id/interactions", customerHandlers.GetInteractions)
	}
}
