package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/response"
	"github.com/manab-pr/evtaarpro/modules/auth/domain/usecases"
)

// LogoutHandler handles user logout
type LogoutHandler struct {
	logoutUseCase *usecases.LogoutUseCase
}

// NewLogoutHandler creates a new LogoutHandler
func NewLogoutHandler(logoutUseCase *usecases.LogoutUseCase) *LogoutHandler {
	return &LogoutHandler{
		logoutUseCase: logoutUseCase,
	}
}

// Handle handles the logout request
// @Summary Logout
// @Description Logout user and invalidate session
// @Tags auth
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /auth/logout [post]
func (h *LogoutHandler) Handle(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "User not authenticated")
		return
	}

	// Execute use case
	if err := h.logoutUseCase.Execute(c.Request.Context(), userID.(string)); err != nil {
		response.InternalServerError(c, "Failed to logout")
		return
	}

	response.OK(c, "Logout successful", nil)
}
