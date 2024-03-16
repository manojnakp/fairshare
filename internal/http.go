package internal

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

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
	w := new(strings.Builder)
	_, _ = writeMessage(w, int(status))
	return w.String()
}

// writeMessage writes code and HTTP status message for given status code.
func writeMessage(w io.Writer, code int) (int, error) {
	return fmt.Fprintf(w, "%d %s", code, http.StatusText(code))
}

// AllowOnly is [http.Handler] in case of HTTP method not in this list of
// allowed methods. It responds with a `405 Method Not Allowed` PlainText
// response along with 'Allow' header set to the list of allowed methods.
type AllowOnly []string

// ServeHTTP implements [http.Handler].
//
// Note: Empty allow list has semantic meaning of no method allowed on the
// target. See [HTTP Semantics] for more info.
//
// [HTTP Semantics]: https://www.rfc-editor.org/rfc/rfc9110.html#name-allow
func (allowed AllowOnly) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", strings.Join(allowed, ", "))
	PlainText(http.StatusMethodNotAllowed).ServeHTTP(w, r)
}
