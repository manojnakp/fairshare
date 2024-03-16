package internal

// dbKey is used to create a value for context key.
type dbKey struct{}

// verifierKey is used to create a value for context key.
type verifierKey struct{}

// tokenKey is used to create a value for context key.
type tokenKey struct{}

// Context Keys used across fairshare.
var (
	DBKey       dbKey       // *sql.DB
	VerifierKey verifierKey // *oidc.IDTokenVerifier
	TokenKey    tokenKey    // *oidc.IDToken
)
