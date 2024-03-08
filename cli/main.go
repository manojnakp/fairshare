package cli

import (
	"flag"
	"log"
	"net"
	"os"

	"github.com/manojnakp/fairshare/api"
)

// Main is the entrypoint of the fairshare server CLI.
func Main() error {
	flag.Parse()
	level := ParseLogLevel(Config.Log)
	InitSlog(os.Stdout, LogSource, level)
	addr := net.JoinHostPort(Config.Host, Config.Port)
	srv := api.ServerBuilder{
		Host: Config.Host,
		Port: Config.Port,
	}.Build()
	log.Println("server listening on", addr)
	return srv.ListenAndServe()
}
