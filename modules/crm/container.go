package crm

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/config"
	"github.com/manab-pr/evtaarpro/internal/datastore"
)

// RegisterRoutes registers CRM module routes
func RegisterRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	crm := rg.Group("/crm")
	{
		// TODO: Implement CRM routes
		crm.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "CRM module - Coming soon"})
		})
	}
}
