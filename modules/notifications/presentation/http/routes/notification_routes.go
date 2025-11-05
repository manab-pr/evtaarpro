package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/middleware"
	"github.com/manab-pr/evtaarpro/modules/notifications/presentation/http/handlers"
)

// RegisterRoutes registers notification routes
func RegisterRoutes(rg *gin.RouterGroup, notificationHandlers *handlers.NotificationHandlers, jwtSecret string) {
	notifications := rg.Group("/notifications")
	notifications.Use(middleware.AuthMiddleware(jwtSecret))
	{
		// Notification routes
		notifications.POST("", notificationHandlers.CreateNotification)
		notifications.GET("", notificationHandlers.ListNotifications)
		notifications.GET("/unread-count", notificationHandlers.GetUnreadCount)
		notifications.PUT("/:id/read", notificationHandlers.MarkAsRead)
		notifications.POST("/read-all", notificationHandlers.MarkAllAsRead)
		notifications.DELETE("/:id", notificationHandlers.DeleteNotification)
	}
}
