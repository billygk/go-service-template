package middleware

import (
	"github.com/gin-gonic/gin"
)

// RetryMiddleware adds retrying to requests
func RetryMiddleware(cfg RetryConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement retrying
		c.Next()
	}
}
