package middleware

import (
	"github.com/gin-gonic/gin"
)

// MetricsMiddleware adds metrics to requests
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement metrics
		c.Next()
	}
}
