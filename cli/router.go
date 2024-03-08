package cli

import "net/http"

// Router defines routing logic for the fairshare server.
func Router() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
