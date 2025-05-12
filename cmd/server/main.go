package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ycyr/splunk2alertmanager/pkg/api"
	"github.com/ycyr/splunk2alertmanager/pkg/config"
	"github.com/ycyr/splunk2alertmanager/pkg/logger"
)

var version string // set via -ldflags "-X main.version=..."

func main() {
	// Version flag
	showVersion := flag.Bool("version", false, "Print the version and exit")
	flag.Parse()
	if *showVersion {
		if version == "" {
			fmt.Println("Version: unknown")
		} else {
			fmt.Printf("Version: %s\n", version)
		}
		os.Exit(0)
	}

	cfg := config.LoadConfig()

	log := logger.NewLogger(cfg.LogLevel, cfg.LogFormat)

	log.Info("Starting HTTP server", "bind_address", cfg.BindAddress)
	if err := api.StartServer(cfg, log); err != nil {
		log.Error("Failed to start HTTP server", "error", err)
		os.Exit(1)
	}
}

