package queries

import (
	"database/sql"
	"errors"
)

type AuthQueries struct {
	*sql.DB
}

func (q *AuthQueries) GetPasswordByEmail(email string) (string, error) {
	var id, value string

	if err := q.DB.QueryRow(`SELECT "id" FROM "setting" WHERE "key"='email' AND "value"=?`, email).Scan(&id); err != nil {
		return "", errors.New("user not found")
	}
	if id == "" {
		return "", errors.New("user with the given email is not found")
	}

	if err := q.DB.QueryRow(`SELECT "value" FROM "setting" WHERE "key"='password'`).Scan(&value); err != nil {
		return "", errors.New("not found user password")
	}
	if value == "" {
		return "", errors.New("user with the given email is not found")
	}

	return value, nil
}
