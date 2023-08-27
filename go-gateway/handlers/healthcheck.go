package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler returns a handler function for the health check endpoint
func HealthCheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
