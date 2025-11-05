package notifications

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/config"
	"github.com/manab-pr/evtaarpro/internal/datastore"
	"github.com/manab-pr/evtaarpro/internal/middleware"
	"github.com/manab-pr/evtaarpro/modules/notifications/infra/postgresql"
	"github.com/manab-pr/evtaarpro/modules/notifications/presentation/http/handlers"
)

// RegisterRoutes registers notifications module routes
func RegisterRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	// Initialize repository
	notificationRepo := postgresql.NewNotificationRepository(pgStore.DB)

	// Initialize handlers
	notificationHandlers := handlers.NewNotificationHandlers(notificationRepo)

	// Register routes with auth middleware
	notifications := rg.Group("/notifications")
	notifications.Use(middleware.AuthMiddleware(cfg.JWT.Secret))
	{
		notifications.POST("", notificationHandlers.CreateNotification)
		notifications.GET("/me", notificationHandlers.ListMyNotifications)
		notifications.GET("/me/unread-count", notificationHandlers.GetUnreadCount)
		notifications.GET("/:id", notificationHandlers.GetNotification)
		notifications.POST("/:id/read", notificationHandlers.MarkAsRead)
		notifications.POST("/:id/unread", notificationHandlers.MarkAsUnread)
		notifications.POST("/mark-all-read", notificationHandlers.MarkAllAsRead)
		notifications.DELETE("/:id", notificationHandlers.DeleteNotification)
	}
}
