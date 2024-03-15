package api

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"

	"github.com/manojnakp/fairshare/internal"

	"github.com/coreos/go-oidc/v3/oidc"
)

// Router defines routing logic for the fairshare server.
func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", PlainText(http.StatusNotFound))
	mux.Handle("GET /sub", Authn(http.HandlerFunc(GetSubject)))
	return mux
}

// GetSubject returns the `subject` claim of the id token presented.
func GetSubject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token := ctx.Value(internal.TokenKey).(*oidc.IDToken)
	subject := token.Subject
	w.WriteHeader(http.StatusOK)
	_, _ = io.WriteString(w, subject)
}

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
		Handler:     With(HeartBeat, Logger)(Router()),
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
