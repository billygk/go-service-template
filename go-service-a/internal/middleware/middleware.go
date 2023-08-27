package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/billygk/go-service-template/go-service-a/internal/config"
	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
)

func ExampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add any necessary middleware logic here
		c.Next()
	}
}

func authMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}
		authHeaderparts := strings.Split(authHeader, " ")
		if len(authHeaderparts) != 2 || authHeaderparts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			return
		}
		tokenString := authHeaderparts[1]
		ctx := context.Background()
		provider, err := oidc.NewProvider(ctx, cfg.OIDC.IssuerURL)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get provider: %v", err))
			return
		}
		oidcConfig := &oidc.Config{
			ClientID: cfg.OIDC.ClientID,
		}
		verifier := provider.Verifier(oidcConfig)
		_, err = verifier.Verify(ctx, tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Next()
	}
}
