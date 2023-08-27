package config

import (
	"github.com/kelseyhightower/envconfig"
)

type OIDCConfig struct {
	IssuerURL    string `envconfig:"OIDC_ISSUER_URL"`
	ClientID     string `envconfig:"OIDC_CLIENT_ID"`
	ClientSecret string `envconfig:"OIDC_CLIENT_SECRET"`
	RedirectURL  string `envconfig:"OIDC_REDIRECT_URL"`
}

type Config struct {
	OIDC *OIDCConfig
	// Add any other necessary configuration options here
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
