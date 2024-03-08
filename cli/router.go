package cli

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Router defines routing logic for the fairshare server.
func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", PlainText(http.StatusNotFound))
	return mux
}

// PlainText is [http.Handler] to write a plain text response consisting
// of status code and text message.
type PlainText int

// ServeHTTP implements [http.Handler].
func (status PlainText) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	code := int(status)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_, _ = writeMessage(w, code)
}

// String implements [fmt.Stringer].
func (status PlainText) String() string {
	var w = new(strings.Builder)
	_, _ = writeMessage(w, int(status))
	return w.String()
}

// writeMessage writes code and HTTP status message for given status code.
func writeMessage(w io.Writer, code int) (int, error) {
	return fmt.Fprintf(w, "%d %s", code, http.StatusText(code))
}
