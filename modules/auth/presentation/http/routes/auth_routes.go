package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/middleware"
	"github.com/manab-pr/evtaarpro/modules/auth/presentation/http/handlers"
)

// RegisterRoutes registers auth routes
func RegisterRoutes(
	rg *gin.RouterGroup,
	registerHandler *handlers.RegisterHandler,
	loginHandler *handlers.LoginHandler,
	logoutHandler *handlers.LogoutHandler,
	jwtSecret string,
) {
	auth := rg.Group("/auth")
	{
		// Public routes
		auth.POST("/register", registerHandler.Handle)
		auth.POST("/login", loginHandler.Handle)

		// Protected routes
		auth.POST("/logout", middleware.AuthMiddleware(jwtSecret), logoutHandler.Handle)
	}
}
