package api

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"

	"github.com/manojnakp/fairshare/api/mw"
)

// ServerBuilder facilitates constructing http.Server for this API.
type ServerBuilder struct {
	Host      string
	Port      string
	Context   context.Context
	TLSConfig *tls.Config
}

// Build constructs a new http.Server with the provided configurations
// to serve this API.
func (sb ServerBuilder) Build() *http.Server {
	return &http.Server{
		Addr:        net.JoinHostPort(sb.Host, sb.Port),
		Handler:     mw.Handle(Router(), mw.HeartBeat, mw.Logger),
		TLSConfig:   sb.TLSConfig,
		BaseContext: sb.getContext,
	}
}

// getContext returns the configured base context (or context.Background if nil).
// This is used to set `BaseContext` field of http.Server.
func (sb ServerBuilder) getContext(net.Listener) (ctx context.Context) {
	ctx = sb.Context
	if ctx == nil {
		ctx = context.Background()
	}
	return ctx
}
