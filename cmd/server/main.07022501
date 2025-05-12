package main

import (
	"os"

	"github.com/ycyr/splunk2alertmanager/pkg/api"
	"github.com/ycyr/splunk2alertmanager/pkg/config"
	"github.com/ycyr/splunk2alertmanager/pkg/logger"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger with log level and format
	log := logger.NewLogger(cfg.LogLevel, cfg.LogFormat)

	// Start HTTP server
	log.Info("Starting HTTP server", "bind_address", cfg.BindAddress)
	if err := api.StartServer(cfg, log); err != nil {
		log.Error("Failed to start HTTP server", "error", err)
		os.Exit(1)
	}
}
