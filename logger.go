package main

import (
	"os"

	charmLog "github.com/charmbracelet/log"

	"starrequest/config"
)

// NewLogger creates a new logger instance
func NewLogger(cfg *config.Config) *charmLog.Logger {
	var logger *charmLog.Logger

	if cfg.Env == "dev" {
		// development mode logger config
		logger = charmLog.NewWithOptions(os.Stdout, charmLog.Options{
			Level:      charmLog.DebugLevel,
			TimeFormat: "03:04:05 PM",
		})
	} else {
		// production mode logger config
		logger = charmLog.NewWithOptions(os.Stdout, charmLog.Options{
			Level:      charmLog.InfoLevel,
			TimeFormat: "01-02-2006 03:04:05 PM",
		})
	}

	// common options
	logger.SetReportCaller(true)
	logger.SetReportTimestamp(true)
	logger.SetCallerFormatter(charmLog.ShortCallerFormatter)

	return logger
}
