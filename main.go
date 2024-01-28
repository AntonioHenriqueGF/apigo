package main

import (
	"github.com/AntonioHenriqueGF/apigo/config"
	"github.com/AntonioHenriqueGF/apigo/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configs: %s", err)
		return
	}

	// Initializes the router
	router.Init()
}
