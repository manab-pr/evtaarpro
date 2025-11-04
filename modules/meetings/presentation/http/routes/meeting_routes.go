package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/middleware"
	"github.com/manab-pr/evtaarpro/modules/meetings/presentation/http/handlers"
)

// RegisterRoutes registers meeting routes
func RegisterRoutes(rg *gin.RouterGroup, handlers *handlers.MeetingHandlers, jwtSecret string) {
	meetings := rg.Group("/meetings")
	meetings.Use(middleware.AuthMiddleware(jwtSecret))
	{
		meetings.POST("", handlers.CreateMeeting)
		meetings.GET("", handlers.ListMeetings)
		meetings.GET("/:id", handlers.GetMeeting)
		meetings.POST("/:id/join", handlers.JoinMeeting)
	}
}
