package main

import (
	"os"

	"github.com/manojnakp/fairshare/cli"
)

func main() {
	if err := cli.Main(); err != nil {
		os.Exit(1)
	}
}
