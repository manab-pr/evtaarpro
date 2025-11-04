package middleware

import (
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/response"
)

// Recovery middleware recovers from panics
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic
				stack := debug.Stack()
				fmt.Printf("PANIC: %v\n%s\n", err, stack)

				// Return error response
				response.InternalServerError(c, "An unexpected error occurred")
				c.Abort()
			}
		}()

		c.Next()
	}
}
