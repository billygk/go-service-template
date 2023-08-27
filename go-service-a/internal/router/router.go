package router

import (
	"github.com/billygk/go-service-template/go-service-a/internal/config"
	"github.com/billygk/go-service-template/go-service-a/internal/handlers"
	"github.com/billygk/go-service-template/go-service-a/internal/middleware"
	"github.com/gin-gonic/gin"
)

func New(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// Add any necessary middleware here
	r.Use(middleware.ExampleMiddleware())

	// Define the endpoints
	r.GET("/api/v1/service-a/", handlers.ServiceAHandler)
	r.GET("/api/v1/service-a/auth-user-1", handlers.AuthUser1Handler)
	r.GET("/api/v1/service-a/auth-user-2", handlers.AuthUser2Handler)

	return r
}
