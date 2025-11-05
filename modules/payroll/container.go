package payroll

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/config"
	"github.com/manab-pr/evtaarpro/internal/datastore"
	"github.com/manab-pr/evtaarpro/modules/payroll/infra/postgresql"
	"github.com/manab-pr/evtaarpro/modules/payroll/presentation/http/handlers"
	"github.com/manab-pr/evtaarpro/modules/payroll/presentation/http/routes"
)

// RegisterRoutes registers payroll module routes
func RegisterRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	// Infrastructure
	payrollRepo := postgresql.NewPayrollRepository(pgStore.DB)

	// Handlers
	payrollHandlers := handlers.NewPayrollHandlers(payrollRepo)

	// Register routes
	routes.RegisterRoutes(rg, payrollHandlers, cfg.JWT.Secret)
}
