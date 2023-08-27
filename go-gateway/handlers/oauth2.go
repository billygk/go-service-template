package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gojektech/heimdall/v6/httpclient"
	"github.com/sony/gobreaker"
)

// OAuth2Handler returns a handler function for the OAuth2 endpoint using Keycloak
func OAuth2Handler(client *httpclient.Client, endpoint string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Make a request to the Keycloak server
		resp, err := client.Get(endpoint)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		// Return the response from the Keycloak server
		c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	}
}

// NewOAuth2Client returns a new HTTP client for the OAuth2 endpoint using Keycloak
func NewOAuth2Client(timeout, retryInterval time.Duration, maxRequests uint32, interval time.Duration) *httpclient.Client {
	return httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
		httpclient.WithRetrier(heimdall.NewRetrier(heimdall.NewConstantBackoff(retryInterval))),
		httpclient.WithCircuitBreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "oauth2",
			MaxRequests: maxRequests,
			Interval:    interval,
			Timeout:     timeout,
		})),
	)
}
