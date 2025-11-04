package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestLogger logs HTTP requests
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Generate request ID
		requestID := uuid.New().String()
		c.Set("request_id", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Log request details
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		// TODO: Replace with proper structured logging
		if errorMessage != "" {
			// Log error
			_ = errorMessage
		}

		// Log info
		_ = statusCode
		_ = method
		_ = path
		_ = clientIP
		_ = latency
		_ = requestID
	}
}
