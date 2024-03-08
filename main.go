package main

import (
	"flag"
	"log"
	"net"
	"net/http"
)

var (
	Port = "8080"
	Host = ""
)

func init() {
	flag.StringVar(&Host, "host", Host, "host to listen on")
	flag.StringVar(&Port, "port", Port, "port to listen on")
}

func main() {
	flag.Parse()
	addr := net.JoinHostPort(Host, Port)
	log.Println("server listening on", addr)
	err := http.ListenAndServe(addr, nil)
	log.Fatal(err)
}
