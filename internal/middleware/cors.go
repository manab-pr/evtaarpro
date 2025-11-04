package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/manab-pr/evtaarpro/internal/config"
)

// CORS middleware
func CORS(cfg config.CORSConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Check if origin is allowed
		allowed := false
		for _, allowedOrigin := range cfg.AllowedOrigins {
			if allowedOrigin == "*" || allowedOrigin == origin {
				allowed = true
				break
			}
		}

		if allowed {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		if cfg.AllowCredentials {
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// Set allowed methods
		if len(cfg.AllowedMethods) > 0 {
			methods := ""
			for i, method := range cfg.AllowedMethods {
				if i > 0 {
					methods += ", "
				}
				methods += method
			}
			c.Writer.Header().Set("Access-Control-Allow-Methods", methods)
		}

		// Set allowed headers
		if len(cfg.AllowedHeaders) > 0 {
			headers := ""
			for i, header := range cfg.AllowedHeaders {
				if i > 0 {
					headers += ", "
				}
				headers += header
			}
			c.Writer.Header().Set("Access-Control-Allow-Headers", headers)
		}

		// Set exposed headers
		if len(cfg.ExposedHeaders) > 0 {
			headers := ""
			for i, header := range cfg.ExposedHeaders {
				if i > 0 {
					headers += ", "
				}
				headers += header
			}
			c.Writer.Header().Set("Access-Control-Expose-Headers", headers)
		}

		// Set max age
		if cfg.MaxAge > 0 {
			c.Writer.Header().Set("Access-Control-Max-Age", fmt.Sprintf("%d", cfg.MaxAge))
		}

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
