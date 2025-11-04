package notifications

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/config"
	"github.com/manab-pr/evtaarpro/internal/datastore"
)

// RegisterRoutes registers notifications module routes
func RegisterRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	notifications := rg.Group("/notifications")
	{
		// TODO: Implement notification routes
		notifications.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Notifications module - Coming soon"})
		})
	}
}
