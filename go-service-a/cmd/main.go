package main

import (
	"github.com/billygk/go-service-template/go-service-a/internal/config"
	"github.com/billygk/go-service-template/go-service-a/internal/router"
)

func main() {
	cfg := config.New()
	r := router.New(cfg)
	r.Run()
}
