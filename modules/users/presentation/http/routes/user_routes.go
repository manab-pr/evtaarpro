package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/middleware"
	"github.com/manab-pr/evtaarpro/modules/users/presentation/http/handlers"
)

// RegisterRoutes registers user routes
func RegisterRoutes(rg *gin.RouterGroup, handlers *handlers.UserHandlers, jwtSecret string) {
	users := rg.Group("/users")
	users.Use(middleware.AuthMiddleware(jwtSecret))
	{
		users.GET("/me", handlers.GetMe)
		users.PUT("/me", handlers.UpdateUser)
		users.GET("", handlers.ListUsers)
		users.GET("/:id", handlers.GetUser)
	}
}
