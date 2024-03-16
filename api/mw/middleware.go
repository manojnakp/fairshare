package mw

import (
	"net/http"
)

// Middleware is an alias for HTTP middleware.
type Middleware = func(http.Handler) http.Handler

// Handle constructs a chain of middlewares using the list of passed ones.
// The first middleware in the argument gets wrapped around by the later
// ones. So, outermost middleware (say logger) should be last argument. Returns
// a http.Handler that passes through the list of middlewares.
func Handle(next http.Handler, middlewares ...Middleware) http.Handler {
	for _, mw := range middlewares {
		next = mw(next)
	}
	return next
}

// HandleFunc is like Handle, but for func.
func HandleFunc(
	next func(w http.ResponseWriter, r *http.Request),
	middlewares ...Middleware,
) http.Handler {
	return Handle(http.HandlerFunc(next), middlewares...)
}
