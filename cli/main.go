package cli

import (
	"io"
	"log/slog"
	"os"

	"github.com/manojnakp/fairshare/api"
	"github.com/manojnakp/fairshare/cli/config"
)

// LogSource ensures that source code position gets logged along with
// log record. TODO: Set to false or remove in production release.
var LogSource = true

// Main is the entrypoint of the fairshare server CLI.
func Main() error {
	err := config.Parse()
	if err != nil {
		return err
	}
	InitSlog(os.Stdout, config.Log())
	slog.Info("config parse", "config", config.Config)
	srv := api.ServerBuilder{
		Host: config.Host(),
		Port: config.Port(),
	}.Build()
	slog.Info("server listen", "address", srv.Addr)
	return srv.ListenAndServe()
}

// InitSlog sets up a structured JSON slog.Logger as default.
func InitSlog(w io.Writer, level slog.Level) {
	handler := slog.NewJSONHandler(w, &slog.HandlerOptions{
		AddSource: LogSource,
		Level:     level,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
