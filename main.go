package main

import (
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/fmo/oauth/internal/handlers"
)

func main() {
	// Setup logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				source := a.Value.Any().(*slog.Source)

				wd, _ := os.Getwd()

				source.File = strings.TrimPrefix(source.File, wd+"/")
			}

			return a
		},
	}))

	// Setup app
	app := handlers.NewApp(logger)

	// Setup router
	mux := http.NewServeMux()
	mux.HandleFunc("/signin", app.Signin)
	mux.HandleFunc("/consent", app.Consent)
	mux.HandleFunc("/oauth/authorize", app.Authorize)
	mux.HandleFunc("/oauth/token", app.Token)
	mux.HandleFunc("/oauth/sessions", app.ListSessions)

	// Run the server
	logger.Info("Server starting", "port", 8080)
	http.ListenAndServe(":8080", mux)
}
