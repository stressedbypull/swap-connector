package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORS is a middleware that sets CORS headers to allow cross-origin requests.
// For security, only GET method is allowed.
// The allowedOrigins parameter can be:
// - "*" to allow all origins (default, not recommended for production)
// - A comma-separated list of specific origins (e.g., "https://example.com,https://app.example.com")
func CORS(allowedOrigins string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		
		// If allowedOrigins is "*", allow any origin
		if allowedOrigins == "*" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		} else if origin != "" {
			// Check if the origin is in the allowed list
			origins := strings.Split(allowedOrigins, ",")
			for _, allowedOrigin := range origins {
				trimmedOrigin := strings.TrimSpace(allowedOrigin)
				if trimmedOrigin == origin {
					c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
					c.Writer.Header().Set("Vary", "Origin")
					break
				}
			}
		}
		
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// Handle preflight OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		
		c.Next()
	}
}
