package api

import (
	"net/http"
)

// Router defines routing logic for the fairshare server.
func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", PlainText(http.StatusNotFound))
	return mux
}
