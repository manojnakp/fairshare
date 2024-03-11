package config

import (
	"errors"
	"log/slog"
	"strconv"
)

// ErrInvalidLogLevel reports that an invalid log level value has been
// given as configuration.
var ErrInvalidLogLevel = errors.New("config: invalid log level")

// String is a custom flag type that represents any string value.
type String string

// String implements flag.Value. String is nil-safe.
func (s *String) String() string {
	if s == nil {
		return ""
	}
	return string(*s)
}

// Set implements flag.Value.
func (s *String) Set(value string) error {
	*s = String(value)
	return nil
}

// Get implements flag.Getter. Returns a value of type `string`,
// or `nil` in case of a nil receiver.
func (s *String) Get() any {
	if s == nil {
		return nil
	}
	return string(*s)
}

// LogLevel is a custom flag type that represents slog.LogLevel.
type LogLevel slog.Level

// String implements flag.Value. String is nil-safe. Wraps around
// [slog.Level.String].
func (s *LogLevel) String() string {
	var level slog.Level
	if s != nil {
		level = slog.Level(*s)
	}
	return level.String()
}

// Set implements flag.Value.
func (s *LogLevel) Set(value string) error {
	var level slog.Level
	n, err := strconv.Atoi(value)
	if err == nil {
		/* numeric log level */
		level = slog.Level(n)
	} else {
		/* string log level */
		level, err = ParseLogLevel(value)
	}
	if err != nil {
		return err
	}
	*s = LogLevel(level)
	return nil
}

// Get implements flag.Getter. Returns a value of type slog.Level,
// or `nil` in case of a nil receiver.
func (s *LogLevel) Get() any {
	if s == nil {
		return nil
	}
	return slog.Level(*s)
}

// ParseLogLevel parses the string log level into slog.Level. In case of
// empty string, returns default log level 0 (slog.LevelInfo). Given an
// incorrect log level string, returns ErrInvalidLogLevel.
func ParseLogLevel(s string) (slog.Level, error) {
	switch s {
	case "":
		return 0, nil
	case "DEBUG":
		return slog.LevelDebug, nil
	case "INFO":
		return slog.LevelInfo, nil
	case "WARN":
		return slog.LevelWarn, nil
	case "ERROR":
		return slog.LevelError, nil
	}
	return 0, ErrInvalidLogLevel
}
