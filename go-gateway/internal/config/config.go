package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config represents the configuration for the gateway service
type Config struct {
	Server         ServerConfig
	RateLimit      RateLimitConfig
	CircuitBreaker CircuitBreakerConfig
	Retry          RetryConfig
	OAuth2         OAuth2Config
}

// ServerConfig represents the configuration for the server
type ServerConfig struct {
	Address string
}

// RateLimitConfig represents the configuration for rate limiting
type RateLimitConfig struct {
	// TODO: Add rate limiting configuration
}

// CircuitBreakerConfig represents the configuration for circuit breaking
type CircuitBreakerConfig struct {
	// TODO: Add circuit breaker configuration
}

// RetryConfig represents the configuration for retrying
type RetryConfig struct {
	// TODO: Add retry configuration
}

// OAuth2Config represents the configuration for OAuth2 using Keycloak
type OAuth2Config struct {
	Endpoint      string
	Timeout       time.Duration
	RetryInterval time.Duration
	MaxRequests   uint32
	Interval      time.Duration
}

// Load loads the configuration from environment variables and/or configuration files
func Load() *Config {
	// Set default values
	viper.SetDefault("server.address", ":8080")
	viper.SetDefault("oauth2.timeout", "2s")
	viper.SetDefault("oauth2.retry_interval", "100ms")
	viper.SetDefault("oauth2.max_requests", uint32(1))
	viper.SetDefault("oauth2.interval", "10s")

	// Bind environment variables
	viper.BindEnv("server.address", "SERVER_ADDRESS")
	viper.BindEnv("oauth2.endpoint", "OAUTH2_ENDPOINT")
	viper.BindEnv("oauth2.timeout", "OAUTH2_TIMEOUT")
	viper.BindEnv("oauth2.retry_interval", "OAUTH2_RETRY_INTERVAL")
	viper.BindEnv("oauth2.max_requests", "OAUTH2_MAX_REQUESTS")
	viper.BindEnv("oauth2.interval", "OAUTH2_INTERVAL")

	// Read configuration files
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/project/")
	viper.AddConfigPath("$HOME/.project")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Unmarshal configuration
	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
