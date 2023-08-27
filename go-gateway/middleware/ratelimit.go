package middleware

import (
	"github.com/gin-gonic/gin"
)

// RateLimitMiddleware adds rate limiting to requests
func RateLimitMiddleware(cfg RateLimitConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement rate limiting
		c.Next()
	}
}
