package main

import (
	"errors"
	"flag"
	"log/slog"
	"os"

	"github.com/manojnakp/fairshare/cli"
	"github.com/manojnakp/fairshare/cli/config"
)

// Main is a wrapper to convert errors cli.Main into exit codes.
func Main() (code int) {
	err := cli.Main()
	if err == nil || errors.Is(err, flag.ErrHelp) {
		return
	}
	slog.Error("main failed", "error", err)
	switch {
	case errors.Is(err, config.ErrConfigParse):
		return 2
	}
	return 1
}

func main() {
	os.Exit(Main())
}
