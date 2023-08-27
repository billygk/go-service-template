package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/billygk/go-service-template/go-service-a/internal/config"
	"github.com/billygk/go-service-template/go-service-a/internal/handlers"
	"github.com/billygk/go-service-template/go-service-a/internal/middleware"
	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func New(cfg *config.Config) (*gin.Engine, error) {
	r := gin.Default()

	// Add any necessary middleware here
	r.Use(middleware.ExampleMiddleware())

	// Set up OAuth2 authentication
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, cfg.OIDC.IssuerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get provider: %v", err)
	}
	oauth2Config := oauth2.Config{
		ClientID:     cfg.OIDC.ClientID,
		ClientSecret: cfg.OIDC.ClientSecret,
		RedirectURL:  cfg.OIDC.RedirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	state := "some-random-state"
	r.GET("/login", func(c *gin.Context) {
		url := oauth2Config.AuthCodeURL(state)
		c.Redirect(http.StatusFound, url)
	})
	r.GET("/callback", func(c *gin.Context) {
		if c.Query("state") != state {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid state"))
			return
		}
		token, err := oauth2Config.Exchange(ctx, c.Query("code"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to exchange token: %v", err))
			return
		}
		rawIDToken, ok := token.Extra("id_token").(string)
		if !ok {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("no id_token in token response"))
			return
		}
		oidcConfig := &oidc.Config{
			ClientID: cfg.OIDC.ClientID,
		}
		idToken, err := provider.Verifier(oidcConfig).Verify(ctx, rawIDToken)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to verify ID token: %v", err))
			return
		}
		var claims struct {
			Email         string `json:"email"`
			EmailVerified bool   `json:"email_verified"`
			PreferredName string `json:"preferred_username"`
		}
		if err := idToken.Claims(&claims); err != nil {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("failed to parse claims: %v", err))
			return
		}
		c.Set("email", claims.Email)
		c.Set("preferred_name", claims.PreferredName)
		c.Next()
	})

	// Define the endpoints
	r.GET("/api/v1/service-a/", handlers.ServiceAHandler)
	r.GET("/api/v1/service-a/auth-user-1", handlers.AuthUser1Handler)
	r.GET("/api/v1/service-a/auth-user-2", handlers.AuthUser2Handler)

	return r, nil
}
