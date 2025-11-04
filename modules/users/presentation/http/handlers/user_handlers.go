package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/response"
	"github.com/manab-pr/evtaarpro/modules/users/domain/entities"
	"github.com/manab-pr/evtaarpro/modules/users/domain/usecases"
	"github.com/manab-pr/evtaarpro/modules/users/presentation/http/dto"
)

// UserHandlers contains user-related HTTP handlers
type UserHandlers struct {
	getUserUC   *usecases.GetUserUseCase
	listUsersUC *usecases.ListUsersUseCase
	updateUserUC *usecases.UpdateUserUseCase
}

// NewUserHandlers creates new UserHandlers
func NewUserHandlers(
	getUserUC *usecases.GetUserUseCase,
	listUsersUC *usecases.ListUsersUseCase,
	updateUserUC *usecases.UpdateUserUseCase,
) *UserHandlers {
	return &UserHandlers{
		getUserUC:   getUserUC,
		listUsersUC: listUsersUC,
		updateUserUC: updateUserUC,
	}
}

// GetMe handles getting the current user's profile
// @Summary Get current user profile
// @Description Get the authenticated user's profile
// @Tags users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} response.Response{data=dto.UserResponse}
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/me [get]
func (h *UserHandlers) GetMe(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "User not authenticated")
		return
	}

	user, err := h.getUserUC.Execute(c.Request.Context(), userID.(string))
	if err != nil {
		response.InternalServerError(c, "Failed to get user")
		return
	}

	response.OK(c, "User retrieved successfully", mapUserToResponse(user))
}

// GetUser handles getting a user by ID
// @Summary Get user by ID
// @Description Get a specific user by their ID
// @Tags users
// @Security BearerAuth
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response{data=dto.UserResponse}
// @Failure 401 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/{id} [get]
func (h *UserHandlers) GetUser(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.getUserUC.Execute(c.Request.Context(), userID)
	if err != nil {
		response.NotFound(c, "User not found")
		return
	}

	response.OK(c, "User retrieved successfully", mapUserToResponse(user))
}

// ListUsers handles listing users with pagination
// @Summary List users
// @Description Get a paginated list of users
// @Tags users
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Page size" default(20)
// @Param search query string false "Search query"
// @Success 200 {object} response.PaginatedResponse
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users [get]
func (h *UserHandlers) ListUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	users, total, err := h.listUsersUC.Execute(c.Request.Context(), page, pageSize)
	if err != nil {
		response.InternalServerError(c, "Failed to list users")
		return
	}

	userResponses := make([]dto.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = mapUserToResponse(user)
	}

	response.Paginated(c, userResponses, page, pageSize, total)
}

// UpdateUser handles updating user profile
// @Summary Update user profile
// @Description Update the current user's profile
// @Tags users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.UpdateUserRequest true "Update details"
// @Success 200 {object} response.Response{data=dto.UserResponse}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/me [put]
func (h *UserHandlers) UpdateUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Unauthorized(c, "User not authenticated")
		return
	}

	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.updateUserUC.Execute(c.Request.Context(), usecases.UpdateInput{
		UserID:     userID.(string),
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Phone:      req.Phone,
		Department: req.Department,
	}); err != nil {
		response.InternalServerError(c, "Failed to update user")
		return
	}

	// Get updated user
	user, _ := h.getUserUC.Execute(c.Request.Context(), userID.(string))

	response.OK(c, "User updated successfully", mapUserToResponse(user))
}

func mapUserToResponse(user *entities.User) dto.UserResponse {
	return dto.UserResponse{
		ID:            user.ID,
		Email:         user.Email,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		FullName:      user.FullName(),
		Phone:         user.Phone,
		Avatar:        user.Avatar,
		Role:          user.Role,
		Department:    user.Department,
		IsActive:      user.IsActive,
		EmailVerified: user.EmailVerified,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}
