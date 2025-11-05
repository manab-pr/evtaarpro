package notifications

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/config"
	"github.com/manab-pr/evtaarpro/internal/datastore"
	"github.com/manab-pr/evtaarpro/modules/notifications/infra/postgresql"
	"github.com/manab-pr/evtaarpro/modules/notifications/presentation/http/handlers"
	"github.com/manab-pr/evtaarpro/modules/notifications/presentation/http/routes"
)

// RegisterRoutes registers notifications module routes
func RegisterRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	// Infrastructure
	notificationRepo := postgresql.NewNotificationRepository(pgStore.DB)

	// Handlers
	notificationHandlers := handlers.NewNotificationHandlers(notificationRepo)

	// Register routes
	routes.RegisterRoutes(rg, notificationHandlers, cfg.JWT.Secret)
}
