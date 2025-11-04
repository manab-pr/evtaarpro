package payroll

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/config"
	"github.com/manab-pr/evtaarpro/internal/datastore"
)

// RegisterRoutes registers payroll module routes
func RegisterRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	payroll := rg.Group("/payroll")
	{
		// TODO: Implement payroll routes
		payroll.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Payroll module - Coming soon"})
		})
	}
}
