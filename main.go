package main

import (
	"fmt"

	"starrequest/config"

	charmLog "github.com/charmbracelet/log"
)

var log *charmLog.Logger

func init() {
	// load the app config
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %v", err)
	}

	// create a logger
	log = NewLogger(cfg)

	log.Info("Config loaded.")
}

func main() {
	log.Info("Hello, world!")
}
