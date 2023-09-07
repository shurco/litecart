package queries

import (
	"database/sql"

	"github.com/shurco/litecart/pkg/errors"
)

// AuthQueries is ...
type AuthQueries struct {
	*sql.DB
}

// GetPasswordByEmail is ...
func (q *AuthQueries) GetPasswordByEmail(email string) (string, error) {
	var id, value string

	if err := q.DB.QueryRow(`SELECT id FROM setting WHERE key = 'email' AND value = ?`, email).Scan(&id); err != nil {
		return "", errors.ErrUserNotFound
	}
	if id == "" {
		return "", errors.ErrUserEmailNotFound
	}

	if err := q.DB.QueryRow(`SELECT value FROM setting WHERE key = 'password'`).Scan(&value); err != nil {
		return "", errors.ErrUserPasswordNotFound
	}
	if value == "" {
		return "", errors.ErrUserEmailNotFound
	}

	return value, nil
}
