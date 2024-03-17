package mw

import (
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"runtime/debug"
	"time"

	"github.com/manojnakp/fairshare/internal"
)

// Logger is an HTTP Middleware to log handling information about every
// incoming request.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* panic handler */
		defer func() {
			if err := recover(); err != nil {
				code := http.StatusInternalServerError
				internal.PlainText(code).ServeHTTP(w, r)
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
		slog.Info("log record", "log-entry", entry)
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
