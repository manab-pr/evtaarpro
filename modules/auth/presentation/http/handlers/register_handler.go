package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/response"
	"github.com/manab-pr/evtaarpro/modules/auth/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/auth/domain/usecases"
	"github.com/manab-pr/evtaarpro/modules/auth/presentation/http/dto"
)

// RegisterHandler handles user registration
type RegisterHandler struct {
	registerUseCase *usecases.RegisterUseCase
}

// NewRegisterHandler creates a new RegisterHandler
func NewRegisterHandler(registerUseCase *usecases.RegisterUseCase) *RegisterHandler {
	return &RegisterHandler{
		registerUseCase: registerUseCase,
	}
}

// Handle handles the register request
// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Registration details"
// @Success 201 {object} response.Response{data=dto.UserResponse}
// @Failure 400 {object} response.Response
// @Failure 409 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /auth/register [post]
func (h *RegisterHandler) Handle(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// Execute use case
	output, err := h.registerUseCase.Execute(c.Request.Context(), usecases.RegisterInput{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      entities.Role(req.Role),
	})

	if err != nil {
		if err == usecases.ErrUserAlreadyExists {
			response.Conflict(c, "User already exists")
			return
		}
		response.InternalServerError(c, "Failed to register user")
		return
	}

	response.Created(c, "User registered successfully", dto.UserResponse{
		UserID: output.UserID,
		Email:  output.Email,
		Role:   req.Role,
	})
}
