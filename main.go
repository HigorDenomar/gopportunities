package main

import (
	"github.com/higordenomar/gopportunities/config"
	"github.com/higordenomar/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	err := config.Init()
	logger = config.GetLogger("main")

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Init("localhost:8080")
}
