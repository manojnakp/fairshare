package internal

// dbKey is used to create a value for context key.
type dbKey struct{}

// Context Keys used across fairshare.
var (
	DBKey dbKey // db conn context key
	_     dbKey
)
