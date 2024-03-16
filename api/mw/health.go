package mw

import (
	"net/http"
	"strings"

	"github.com/manojnakp/fairshare/internal"
)

// HealthChecker is the URI endpoint at which heart beat messages are served.
var HealthChecker = "/health"

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
			internal.AllowOnly{MethodGet}.ServeHTTP(w, r)
			return
		}
		/* prevent caching */
		w.Header().Set("Cache-Control", "no-cache, no-store")
		internal.PlainText(http.StatusOK).ServeHTTP(w, r)
	})
}
