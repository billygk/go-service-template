package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthUser1Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from Auth User 1!",
	})
}
