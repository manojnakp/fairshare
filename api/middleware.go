package api

import (
	"encoding/json"
	"github.com/coreos/go-oidc/v3/oidc"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"runtime/debug"
	"strings"
	"time"

	"github.com/manojnakp/fairshare/internal"
)

// bearerPrefix is the prefix in Authorization header.
const bearerPrefix = "Bearer "

// HealthChecker is the URI endpoint at which heart beat messages are served.
var HealthChecker = "/health"

// Middleware is an alias for HTTP middleware.
type Middleware = func(http.Handler) http.Handler

// With constructs a chain of middlewares using the list of passed ones.
// The first middleware in the argument gets wrapped around by the later
// ones. So, outermost middleware (say logger) should be last argument.
func With(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, mw := range middlewares {
			next = mw(next)
		}
		return next
	}
}

// HeartBeat is an HTTP Middleware that listens to heart beat (alive)
// requests at the HealthChecker endpoint. Basically, it returns with a
// `200 OK` message as long as the server is alive and kicking.
func HeartBeat(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.EqualFold(r.URL.Path, HealthChecker) {
			/* path does not match */
			next.ServeHTTP(w, r)
			return
		}
		const MethodGet = "GET"
		if r.Method != MethodGet {
			/* method does not match */
			AllowOnly{MethodGet}.ServeHTTP(w, r)
			return
		}
		/* prevent caching */
		w.Header().Set("Cache-Control", "no-cache, no-store")
		PlainText(http.StatusOK).ServeHTTP(w, r)
	})
}

// Authenticator is an HTTP Middleware that checks Authorization header for
// OpenID Connect Bearer Token.
func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		prefix := authHeader[:len(bearerPrefix)]
		if !strings.EqualFold(prefix, bearerPrefix) {
			/* invalid authorization header */
			PlainText(http.StatusUnauthorized).ServeHTTP(w, r)
			slog.Warn("invalid Authorization header", "prefix", prefix)
			return
		}
		ctx := r.Context()
		rawToken := authHeader[len(bearerPrefix):]
		verifier := ctx.Value(internal.AuthKey).(*oidc.IDTokenVerifier)
		token, err := verifier.Verify(ctx, rawToken)
		if err != nil {
			/* not authenticated */
			PlainText(http.StatusUnauthorized).ServeHTTP(w, r)
			slog.Warn("failed to verify id token", "error", err)
			return
		}
		/* TODO: use token */
		_ = token
		next.ServeHTTP(w, r)
	})
}

// Logger is an HTTP Middleware to log handling information about every
// incoming request.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* panic handler */
		defer func() {
			if err := recover(); err != nil {
				code := http.StatusInternalServerError
				PlainText(code).ServeHTTP(w, r)
				log.Println(err, string(debug.Stack()))
			}
		}()
		/* build log entry */
		wrapper := internal.NewResponseWriter(w)
		entry := &LogEntry{
			Addr:     r.RemoteAddr,
			Proto:    r.Proto,
			Method:   r.Method,
			Host:     r.Host,
			URI:      r.URL,
			Incoming: r.Header,
		}
		start := time.Now()
		next.ServeHTTP(wrapper, r)
		entry.Duration = time.Since(start)
		entry.Status = wrapper.StatusCode
		entry.Size = wrapper.BodyLength
		entry.Outgoing = wrapper.Header()
		/* log the record */
		buf, _ := json.Marshal(entry)
		log.Println(string(buf))
	})
}

// LogEntry denotes the single record being logged on every request.
type LogEntry struct {
	/* request properties */

	Addr   string   `json:"address,omitempty"`
	Proto  string   `json:"proto,omitempty"`
	Method string   `json:"method,omitempty"`
	Host   string   `json:"host,omitempty"`
	URI    *url.URL `json:"uri,omitempty"`

	/* headers */

	Incoming http.Header `json:"request_headers,omitempty"`
	Outgoing http.Header `json:"response_headers,omitempty"`

	/* overall properties */

	Duration time.Duration `json:"duration,omitempty"`

	/* response properties */

	Status int `json:"status,omitempty"`
	Size   int `json:"size,omitempty"`
}
