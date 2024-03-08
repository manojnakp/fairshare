package api

import (
	"net/http"
)

// Router defines routing logic for the fairshare server.
func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", PlainText(http.StatusNotFound))
	mux.HandleFunc("GET /health", HealthCheck)
	mux.Handle("/health", AllowOnly{"GET"})
	return mux
}

// HealthCheck handles the server health endpoint. Basically responds
// `200 OK` as long as server is alive and kicking.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	PlainText(http.StatusOK).ServeHTTP(w, r)
}
