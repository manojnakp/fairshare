package cli

import (
	"flag"
	"log"
	"net"
	"net/http"
)

// Main is the entrypoint of the fairshare server CLI.
func Main() error {
	flag.Parse()
	addr := net.JoinHostPort(Config.Host, Config.Port)
	log.Println("server listening on", addr)
	return http.ListenAndServe(addr, Router())
}
