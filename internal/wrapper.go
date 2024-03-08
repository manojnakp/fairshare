package internal

import "net/http"

// ResponseWriter is a wrapper around [http.ResponseWriter] to capture
// status code and size of response body for logging and other purposes.
type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	BodyLength int
}

// NewResponseWriter is the constructor for ResponseWriter.
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{ResponseWriter: w, StatusCode: http.StatusOK}
}

// Write is a wrapper around calls to underlying [http.ResponseWriter].
func (rw *ResponseWriter) Write(buf []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(buf)
	rw.BodyLength += n
	return n, err
}

// WriteHeader is a wrapper around calls to underlying [http.ResponseWriter].
func (rw *ResponseWriter) WriteHeader(code int) {
	rw.ResponseWriter.WriteHeader(code)
	rw.StatusCode = code
}

// Unwrap returns the underlying [http.ResponseWriter]. As for why this is
// needed, see [http.ResponseController].
func (rw *ResponseWriter) Unwrap() http.ResponseWriter {
	return rw.ResponseWriter
}
