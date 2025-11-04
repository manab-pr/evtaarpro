package httpx

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/manab-pr/evtaarpro/internal/config"
	"github.com/manab-pr/evtaarpro/internal/datastore"
	"github.com/manab-pr/evtaarpro/internal/middleware"
	"github.com/manab-pr/evtaarpro/internal/response"

	// Import modules
	authModule "github.com/manab-pr/evtaarpro/modules/auth"
	usersModule "github.com/manab-pr/evtaarpro/modules/users"
	meetingsModule "github.com/manab-pr/evtaarpro/modules/meetings"
	crmModule "github.com/manab-pr/evtaarpro/modules/crm"
	payrollModule "github.com/manab-pr/evtaarpro/modules/payroll"
	notificationsModule "github.com/manab-pr/evtaarpro/modules/notifications"
)

// NewRouter creates and configures the application router
func NewRouter(cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) *gin.Engine {
	router := gin.New()

	// Global middleware
	router.Use(middleware.Recovery())
	router.Use(middleware.RequestLogger())
	router.Use(middleware.CORS(cfg.CORS))

	// Health check endpoint
	router.GET("/health", healthCheckHandler(pgStore, redisStore))
	router.GET("/ready", readinessHandler(pgStore, redisStore))

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Initialize and register module routes
		registerAuthRoutes(v1, cfg, pgStore, redisStore)
		registerUserRoutes(v1, cfg, pgStore, redisStore)
		registerMeetingRoutes(v1, cfg, pgStore, redisStore)
		registerCRMRoutes(v1, cfg, pgStore, redisStore)
		registerPayrollRoutes(v1, cfg, pgStore, redisStore)
		registerNotificationRoutes(v1, cfg, pgStore, redisStore)
	}

	// 404 handler
	router.NoRoute(func(c *gin.Context) {
		response.NotFound(c, "Route not found")
	})

	return router
}

// Health check handler
func healthCheckHandler(pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "evtaarpro",
		})
	}
}

// Readiness check handler
func readinessHandler(pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check PostgreSQL
		if err := pgStore.Health(c.Request.Context()); err != nil {
			response.InternalServerError(c, "Database not ready")
			return
		}

		// Check Redis
		if err := redisStore.Health(c.Request.Context()); err != nil {
			response.InternalServerError(c, "Redis not ready")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":   "ready",
			"postgres": "connected",
			"redis":    "connected",
		})
	}
}

// Module route registration functions
func registerAuthRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	authModule.RegisterRoutes(rg, cfg, pgStore, redisStore)
}

func registerUserRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	usersModule.RegisterRoutes(rg, cfg, pgStore, redisStore)
}

func registerMeetingRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	meetingsModule.RegisterRoutes(rg, cfg, pgStore, redisStore)
}

func registerCRMRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	crmModule.RegisterRoutes(rg, cfg, pgStore, redisStore)
}

func registerPayrollRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	payrollModule.RegisterRoutes(rg, cfg, pgStore, redisStore)
}

func registerNotificationRoutes(rg *gin.RouterGroup, cfg *config.Config, pgStore *datastore.PostgresStore, redisStore *datastore.RedisStore) {
	notificationsModule.RegisterRoutes(rg, cfg, pgStore, redisStore)
}
