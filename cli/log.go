package cli

import (
	"io"
	"log/slog"
)

// LogSource ensures that source code position gets logged along with
// log record. TODO: Set to false or remove in production release.
var LogSource = true

// InitSlog sets up a structured JSON slog.Logger as default.
func InitSlog(w io.Writer, source bool, level slog.Level) {
	handler := slog.NewJSONHandler(w, &slog.HandlerOptions{
		AddSource: source,
		Level:     level,
	})
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

// ParseLogLevel parses the string log level into slog.Level. In case of
// no matches, returns default log level 0 (slog.LevelInfo).
func ParseLogLevel(log string) (level slog.Level) {
	switch log {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	}
	return
}
