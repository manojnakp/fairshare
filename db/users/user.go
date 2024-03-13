package users

import (
	"database/sql"
	"time"

	"github.com/manojnakp/fairshare/db"

	"github.com/google/uuid"
)

// User represents a row from `users` table.
type User struct {
	Uid     uuid.UUID
	VPA     sql.Null[string]
	LastAdd time.Time
	LastMod time.Time
}

// UserFrom constructs a user from a given scannable.
func UserFrom(row db.Scannable) (user User, err error) {
	var u User
	err = row.Scan(&u.Uid, &u.VPA, &u.LastAdd, &u.LastMod)
	if err != nil {
		return
	}
	return u, nil
}
