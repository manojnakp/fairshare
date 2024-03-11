package config

import (
	"flag"
	"os"
)

// env is a mapping from env var name to flag.Value.
var env = map[string]flag.Value{}

// EnvRegister registers a flag.Value at the specified variable name.
func EnvRegister(value flag.Value, name string) {
	env[name] = value
}

// ParseEnv parses registered environment variables from process env.
func ParseEnv() error {
	for name, flagValue := range env {
		value, ok := os.LookupEnv(name)
		if !ok {
			continue
		}
		err := flagValue.Set(value)
		if err != nil {
			return err
		}
	}
	return nil
}
