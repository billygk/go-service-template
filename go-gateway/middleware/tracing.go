package middleware

import (
	"github.com/gin-gonic/gin"
)

// TracingMiddleware adds tracing to requests
func TracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement tracing
		c.Next()
	}
}
