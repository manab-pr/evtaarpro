package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorInfo  `json:"error,omitempty"`
}

// ErrorInfo represents error details
type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// PaginatedResponse represents a paginated API response
type PaginatedResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

// Pagination represents pagination metadata
type Pagination struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalPages int   `json:"total_pages"`
	TotalItems int64 `json:"total_items"`
}

// Success sends a successful response
func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error sends an error response
func Error(c *gin.Context, statusCode int, code, message, details string) {
	c.JSON(statusCode, Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
			Details: details,
		},
	})
}

// Paginated sends a paginated response
func Paginated(c *gin.Context, data interface{}, page, pageSize int, totalItems int64) {
	totalPages := int((totalItems + int64(pageSize) - 1) / int64(pageSize))

	c.JSON(http.StatusOK, PaginatedResponse{
		Success: true,
		Data:    data,
		Pagination: Pagination{
			Page:       page,
			PageSize:   pageSize,
			TotalPages: totalPages,
			TotalItems: totalItems,
		},
	})
}

// BadRequest sends a 400 Bad Request response
func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, "BAD_REQUEST", message, "")
}

// Unauthorized sends a 401 Unauthorized response
func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, "UNAUTHORIZED", message, "")
}

// Forbidden sends a 403 Forbidden response
func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, "FORBIDDEN", message, "")
}

// NotFound sends a 404 Not Found response
func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, "NOT_FOUND", message, "")
}

// Conflict sends a 409 Conflict response
func Conflict(c *gin.Context, message string) {
	Error(c, http.StatusConflict, "CONFLICT", message, "")
}

// InternalServerError sends a 500 Internal Server Error response
func InternalServerError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", message, "")
}

// Created sends a 201 Created response
func Created(c *gin.Context, message string, data interface{}) {
	Success(c, http.StatusCreated, message, data)
}

// OK sends a 200 OK response
func OK(c *gin.Context, message string, data interface{}) {
	Success(c, http.StatusOK, message, data)
}

// NoContent sends a 204 No Content response
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
