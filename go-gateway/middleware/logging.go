package middleware

import (
	"github.com/gin-gonic/gin"
)

// LoggingMiddleware adds logging to requests
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement logging
		c.Next()
	}
}
