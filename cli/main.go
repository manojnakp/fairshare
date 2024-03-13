package cli

import (
	"context"
	"database/sql"
	"io"
	"log/slog"
	"os"

	"github.com/manojnakp/fairshare/api"
	"github.com/manojnakp/fairshare/cli/config"
	"github.com/manojnakp/fairshare/internal"

	_ "github.com/lib/pq"
)

// LogSource ensures that source code position gets logged along with
// log record. TODO: Set to false or remove in production release.
var LogSource = true

// Main is the entrypoint of the fairshare server CLI.
func Main() error {
	var ctx = context.Background()
	/* parse config */
	err := config.Parse()
	if err != nil {
		return err
	}
	/* setup slog */
	InitSlog(os.Stdout, config.Log())
	slog.Info("config parse", "config", config.Config)
	/* setup database */
	db, err := InitDB()
	if err != nil {
		return err
	}
	slog.Info("db connected")
	/* setup server */
	ctx = context.WithValue(ctx, internal.DBKey, db)
	srv := api.ServerBuilder{
		Host:    config.Host(),
		Port:    config.Port(),
		Context: ctx,
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

// InitDB sets up database connection, or returns error in case of failure.
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.DB())
	if err != nil {
		return nil, err
	}
	/* ping to verify db is really connected */
	err = db.PingContext(context.Background())
	if err != nil {
		return nil, err
	}
	return db, nil
}
