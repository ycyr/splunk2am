package api

import (
	"net/http"

	"log/slog"

	"github.com/ycyr/splunk2alertmanager/pkg/alertmanager"
	"github.com/ycyr/splunk2alertmanager/pkg/config"
)

// StartServer starts the HTTP server
func StartServer(cfg config.Config, logger *slog.Logger) error {
	http.HandleFunc("/splunk-webhook", func(w http.ResponseWriter, r *http.Request) {
		alertmanager.HandleSplunkWebhook(w, r, cfg, logger)
	})
	return http.ListenAndServe(cfg.BindAddress, nil)
}
