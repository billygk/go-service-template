package middleware

import (
	"github.com/gin-gonic/gin"
)

func ExampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add any necessary middleware logic here
		c.Next()
	}
}
