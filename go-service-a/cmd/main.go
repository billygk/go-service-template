package main

import (
	"fmt"

	"github.com/billygk/go-service-template/go-service-a/internal/config"
	"github.com/billygk/go-service-template/go-service-a/internal/router"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		fmt.Printf("failed to load config")
		panic(err)
	}
	r, err := router.New(cfg)
	if err != nil {
		fmt.Printf("failed to create router")
		panic(err)
	}
	r.Run()
}
