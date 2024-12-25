package main

import (
	"fmt"
	"starrequest/config"

	"starrequest/pkg/parser"

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
	// create a parser
	parser := parser.NewParser(log)

	// parse the file
	req, err := parser.Parse("./examples/ex1.sr.yaml")
	if err != nil {
		log.Error("Failed to parse yaml", "error", err)
	}

	fmt.Println(req)
}
