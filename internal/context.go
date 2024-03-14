package internal

// dbKey is used to create a value for context key.
type dbKey struct{}

// authKey is used to create a value for context key.
type authKey struct{}

// Context Keys used across fairshare.
var (
	DBKey   dbKey   // *sql.DB
	AuthKey authKey // *oidc.IDTokenVerifier
)
