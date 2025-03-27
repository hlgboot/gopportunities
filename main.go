package main

import (
	"github.com/hlgboot/gopportunities/config"
	"github.com/hlgboot/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
