package router

import (
	"github.com/gin-gonic/gin"

	"github.com/billygk/go-service-template/go-gateway/internal/config"
	"github.com/billygk/go-service-template/go-gateway/internal/handlers"
	"github.com/billygk/go-service-template/go-gateway/internal/middleware"
)

// New returns a new Gin router
func New(cfg *config.Config) *gin.Engine {
	// Create the router
	r := gin.Default()

	// Add middleware for rate limiting, circuit breaking, and retrying
	r.Use(middleware.RateLimitMiddleware(cfg.RateLimit))
	r.Use(middleware.CircuitBreakerMiddleware(cfg.CircuitBreaker))
	r.Use(middleware.RetryMiddleware(cfg.Retry))

	// Add middleware for logging, metrics, and tracing
	r.Use(middleware.LoggingMiddleware())
	r.Use(middleware.MetricsMiddleware())
	r.Use(middleware.TracingMiddleware())

	// Add a health check endpoint
	r.GET("/health", handlers.HealthCheckHandler())

	// Add an OAuth2 endpoint using Keycloak
	r.GET("/oauth2", handlers.OAuth2Handler(
		handlers.NewOAuth2Client(cfg.OAuth2.Timeout, cfg.OAuth2.RetryInterval, cfg.OAuth2.MaxRequests, cfg.OAuth2.Interval),
		cfg.OAuth2.Endpoint,
	))

	return r
}
