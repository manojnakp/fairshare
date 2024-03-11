package config

import (
	"flag"
	"log/slog"
)

// Config is the configuration options passed to the fairshare server.
// Means of conveying configuration may be:
//
//   - Command line flags
var Config = struct {
	Host String   `json:"host,omitempty"`
	Port String   `json:"port,omitempty"`
	Log  LogLevel `json:"log,omitempty"`
}{
	Port: "8080",
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

// initialize command line flags for parsing configuration
func init() {
	flag.Var(&Config.Host, "host", "host to listen on")
	flag.Var(&Config.Port, "port", "port to listen on")
	flag.Var(&Config.Log, "log", "log level above which records get logged")
}

// Parse is a wrapper around flag.Parse.
func Parse() {
	flag.Parse()
}
