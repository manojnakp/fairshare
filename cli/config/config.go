package config

import (
	"errors"
	"flag"
	"log/slog"
	"os"
)

// Config default values.
const (
	DefaultPort     = "8080"
	DefaultDB       = "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable"
	DefaultAuth     = "http://localhost:9090"
	DefaultAudience = "fairshare-web"
)

// ErrConfigParse is a generic error that represents config parsing failed.
var ErrConfigParse = errors.New("config: failed to parse configuration")

// cmd is a custom flag set used by config.
var cmd = flag.NewFlagSet("fairshare", flag.ContinueOnError)

// Config is the configuration options passed to the fairshare server.
// Means of conveying configuration may be:
//
//   - Command line flags
//   - Environment Variables
var Config = struct {
	Host     String   `json:"host,omitempty"`
	Port     String   `json:"port,omitempty"`
	Log      LogLevel `json:"log,omitempty"`
	DB       String   `json:"db,omitempty"`
	Auth     String   `json:"auth,omitempty"`
	Audience String   `json:"audience,omitempty"`
}{
	Port:     DefaultPort,
	DB:       DefaultDB,
	Auth:     DefaultAuth,
	Audience: DefaultAudience,
}

// Parse parses the configuration from command line flags and environment
// variables. Entrypoint of this package.
func Parse() error {
	var err error
	err = ParseEnv()
	if err != nil {
		return wrap(err)
	}
	err = cmd.Parse(os.Args[1:])
	if err != nil {
		return wrap(err)
	}
	return nil
}

// wrap is a utility to wrap any given error with ErrConfigParse.
func wrap(err error) error {
	return errors.Join(ErrConfigParse, err)
}

// Host return Config.Host as a string. Not thread-safe.
func Host() string {
	return Config.Host.Get().(string)
}

// Port returns Config.Port as a string. Not thread-safe.
func Port() string {
	return Config.Port.Get().(string)
}

// Log returns Config.Log as slog.Level. Not thread-safe.
func Log() slog.Level {
	return Config.Log.Get().(slog.Level)
}

// DB returns Config.DB as a string. Not thread-safe.
func DB() string {
	return Config.DB.Get().(string)
}

// Auth returns Config.Auth as a string. Not thread-safe.
func Auth() string {
	return Config.Auth.Get().(string)
}

// Audience returns Config.Audience as a string. Not thread-safe.
func Audience() string {
	return Config.Audience.Get().(string)
}
