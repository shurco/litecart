package queries

import (
	"database/sql"
	"strconv"

	"github.com/shurco/litecart/pkg/jwtutil"
	"github.com/stripe/stripe-go/v74/client"
)

type SettingQueries struct {
	*sql.DB
}

type Stripe struct {
	SecretKey  string
	WebhookKey string
	Client     *client.API
}

// IsInstalled is ...
func (q *SettingQueries) IsInstalled() bool {
	var installed bool
	q.DB.QueryRow(`SELECT "value" FROM "setting" WHERE "key" = 'installed'`).Scan(&installed)
	return installed
}

// CheckSubdomain is ...
func (q *SettingQueries) CheckSubdomain(name string) bool {
	var id int
	err := q.DB.QueryRow(`SELECT "id" FROM "domain" WHERE "name" = ?`, name).Scan(&id)
	return err == nil
}

// GetSession is ...
func (q *SettingQueries) GetSession(key string) (string, error) {
	var value string
	err := q.DB.QueryRow(`SELECT "value" FROM "session" WHERE "key" = ?`, key).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}

// AddSession is ...
func (q *SettingQueries) AddSession(key, value string, expires int64) error {
	_, err := q.DB.Exec(`INSERT INTO "session" ("key", "value", "expires") VALUES (?, ?, ?)`, key, value, expires)
	if err != nil {
		return err
	}
	return nil
}

// UpdateSession is ...
func (q *SettingQueries) UpdateSession(key, value string, expires int64) error {
	_, err := q.DB.Exec(`UPDATE "session" SET "value" = ?, "expires" = ? WHERE "key" = ? `, value, expires, key)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSession is ...
func (q *SettingQueries) DeleteSession(key string) error {
	_, err := q.DB.Exec(`DELETE FROM "session" WHERE "key" = ?`, key)
	if err != nil {
		return err
	}
	return nil
}

// SettingJWT is settings
func (q *SettingQueries) SettingJWT() (*jwtutil.Setting, error) {
	settings := &jwtutil.Setting{}

	query := `SELECT "key", "value" FROM "setting" WHERE "key" IN (?, ?)`
	rows, err := q.DB.Query(query, "jwt_secret", "jwt_secret_expire_hours")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var key, value string
		err := rows.Scan(&key, &value)
		if err != nil {
			return nil, err
		}

		switch key {
		case "jwt_secret":
			settings.Secret = value
		case "jwt_secret_expire_hours":
			settings.SecretExpireHours, _ = strconv.Atoi(value)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return settings, nil
}

// SettingStripe is ...
func (q *SettingQueries) SettingStripe() (*Stripe, error) {
	settings := &Stripe{}

	query := `SELECT "key", "value" FROM "setting" WHERE "key" IN (?, ?)`
	rows, err := q.DB.Query(query, "stripe_secret_key", "stripe_webhook_secret_key")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var key, value string
		err := rows.Scan(&key, &value)
		if err != nil {
			return nil, err
		}

		switch key {
		case "stripe_secret_key":
			settings.SecretKey = value
		case "stripe_webhook_secret_key":
			settings.WebhookKey = value
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return settings, nil
}
