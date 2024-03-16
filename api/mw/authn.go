package mw

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/manojnakp/fairshare/internal"

	"github.com/coreos/go-oidc/v3/oidc"
)

// bearerPrefix is the prefix in Authorization header.
const bearerPrefix = "Bearer "

// Authn is an HTTP Middleware that checks Authorization header for
// OpenID Connect Bearer Token.
func Authn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) < len(bearerPrefix) {
			/* too small authorization header */
			internal.PlainText(http.StatusUnauthorized).ServeHTTP(w, r)
			slog.Warn("missing or empty Authorization header", "header", authHeader)
			return
		}
		prefix := authHeader[:len(bearerPrefix)]
		if !strings.EqualFold(prefix, bearerPrefix) {
			/* invalid authorization header */
			internal.PlainText(http.StatusUnauthorized).ServeHTTP(w, r)
			slog.Warn("invalid Authorization header", "prefix", prefix)
			return
		}
		ctx := r.Context()
		rawToken := authHeader[len(bearerPrefix):]
		verifier := ctx.Value(internal.VerifierKey).(*oidc.IDTokenVerifier)
		token, err := verifier.Verify(ctx, rawToken)
		if err != nil {
			/* not authenticated */
			internal.PlainText(http.StatusUnauthorized).ServeHTTP(w, r)
			slog.Warn("failed to verify id token", "error", err)
			return
		}
		ctx = context.WithValue(ctx, internal.TokenKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
