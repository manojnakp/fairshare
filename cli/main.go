package cli

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/manojnakp/fairshare/api"
)

// Main is the entrypoint of the fairshare server CLI.
func Main() error {
	flag.Parse()
	addr := net.JoinHostPort(Config.Host, Config.Port)
	srv := &http.Server{
		Addr:        addr,
		Handler:     api.With(api.Logger, api.HeartBeat)(api.Router()),
		BaseContext: nil,
	}
	log.Println("server listening on", addr)
	return srv.ListenAndServe()
}
