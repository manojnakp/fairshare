package db

import (
	"context"
	"database/sql"
)

// TxnMgr is a transaction manager that can begin new transaction.
type TxnMgr interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

// Handler is an SQL handler that can query and execute sql on database.
type Handler interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

// Scannable represents any source that can be scanned into an object.
// Typically, it is [sql.Row] or [sql.Rows].
type Scannable interface {
	Scan(dest ...any) error
}
