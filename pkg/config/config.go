package config

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Config holds the application configuration
var version string // set via -ldflags "-X github.com/ycyr/splunk2alertmanager/pkg/config.version=..."

type Config struct {
	ShowVersion      bool
	AlertmanagerURL  string
	BindAddress      string
	LogLevel         string
	LogFormat        string
	EndsAtDuration   string
	AdditionalLabels []string
	AnnotationPrefix string // New field for annotation prefix
}

// LoadConfig parses command-line flags and returns a Config struct
func LoadConfig() Config {
	versionFlag := flag.Bool("version", false, "Print the version and exit")
	alertmanagerURL := flag.String("u", "http://localhost:9093", "URL of the Alertmanager instance (`-u`, `--alertmanager-url`)")
	bindAddress := flag.String("b", "localhost:8080", "Bind address for the HTTP server (`-b`, `--bind`)")
	logLevel := flag.String("l", "info", "Log level (debug, info, warn, error) (`-l`, `--log-level`)")
	logFormat := flag.String("f", "text", "Log format (json or text) (`-f`, `--log-format`)")
	endsAtDuration := flag.String("e", "", "Duration for EndsAt (e.g., 1h, 30m, 15s) (`-e`, `--ends-at`); leave empty for no EndsAt")
	additionalLabels := flag.String("add-labels", "", "Comma-separated list of additional labels to include from the Splunk result (`--add-labels`)")
	annotationPrefix := flag.String("p", "ann.", "Prefix for detecting annotations (`-p`, `--annotation-prefix`)") // Default = "ann."

	flag.Parse()
	if *versionFlag {
		fmt.Printf("Version: %s\n", version)
		os.Exit(0)
	}

	return Config{
		ShowVersion:      *versionFlag,
		AlertmanagerURL:  *alertmanagerURL,
		BindAddress:      *bindAddress,
		LogLevel:         *logLevel,
		LogFormat:        *logFormat,
		EndsAtDuration:   *endsAtDuration,
		AdditionalLabels: strings.Split(*additionalLabels, ","),
		AnnotationPrefix: *annotationPrefix,
	}
}

