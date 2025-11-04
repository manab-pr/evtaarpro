package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/response"
	"github.com/manab-pr/evtaarpro/modules/auth/domain/usecases"
	"github.com/manab-pr/evtaarpro/modules/auth/presentation/http/dto"
)

// LoginHandler handles user login
type LoginHandler struct {
	loginUseCase *usecases.LoginUseCase
}

// NewLoginHandler creates a new LoginHandler
func NewLoginHandler(loginUseCase *usecases.LoginUseCase) *LoginHandler {
	return &LoginHandler{
		loginUseCase: loginUseCase,
	}
}

// Handle handles the login request
// @Summary Login
// @Description Authenticate user and return JWT tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login credentials"
// @Success 200 {object} response.Response{data=dto.AuthResponse}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /auth/login [post]
func (h *LoginHandler) Handle(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// Execute use case
	output, err := h.loginUseCase.Execute(c.Request.Context(), usecases.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		if err == usecases.ErrInvalidCredentials {
			response.Unauthorized(c, "Invalid email or password")
			return
		}
		if err == usecases.ErrUserNotActive {
			response.Forbidden(c, "User account is not active")
			return
		}
		response.InternalServerError(c, "Failed to login")
		return
	}

	response.OK(c, "Login successful", dto.AuthResponse{
		UserID:       output.UserID,
		Email:        output.Email,
		Role:         output.Role,
		AccessToken:  output.AccessToken,
		RefreshToken: output.RefreshToken,
	})
}
