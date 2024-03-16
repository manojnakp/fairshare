package api

import (
	"io"
	"net/http"

	"github.com/manojnakp/fairshare/api/mw"
	"github.com/manojnakp/fairshare/internal"

	"github.com/coreos/go-oidc/v3/oidc"
)

// Router defines routing logic for the fairshare server.
func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", internal.PlainText(http.StatusNotFound))
	mux.Handle("GET /sub", mw.HandleFunc(GetSubject, mw.Authn))
	return mux
}

// GetSubject returns the `subject` claim of the id token presented.
func GetSubject(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token := ctx.Value(internal.TokenKey).(*oidc.IDToken)
	subject := token.Subject
	w.WriteHeader(http.StatusOK)
	_, _ = io.WriteString(w, subject)
}
