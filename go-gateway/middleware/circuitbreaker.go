package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

// CircuitBreakerMiddleware adds circuit breaking to requests
func CircuitBreakerMiddleware(cfg CircuitBreakerConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement circuit breaking
		c.Next()
	}
}

// NewCircuitBreaker returns a new circuit breaker
func NewCircuitBreaker(cfg CircuitBreakerConfig) *gobreaker.CircuitBreaker {
	return gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "default",
		MaxRequests: cfg.MaxRequests,
		Interval:    cfg.Interval,
		Timeout:     cfg.Timeout,
	})
}
