package main

// This command should be run in the terminal, not in the code:
// go get github.com/billygk/go-service-template/go-gateway/internal/config

import (
	"github.com/billygk/go-service-template/go-gateway/internal/config"
	"github.com/billygk/go-service-template/go-gateway/internal/router"
)

func main() {
	// Load the configuration
	cfg := config.Load()

	// Create the router
	r := router.New(cfg)

	// Start the server
	r.Run(cfg.Server.Address)
}
go get github.com/billygk/go-service-template/go-gateway/internal/router
func main() {
	// Load the configuration
	cfg := config.Load()

	// Create the router
	r := router.New(cfg)

	// Start the server
	r.Run(cfg.Server.Address)
}
