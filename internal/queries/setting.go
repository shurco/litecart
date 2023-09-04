package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/pkg/jwtutil"
	"github.com/shurco/litecart/pkg/security"
)

// SettingQueries is ...
type SettingQueries struct {
	*sql.DB
}

// Settings is ...
func (q *SettingQueries) Settings() (*models.Setting, error) {
	settings := &models.Setting{}

	keys := []any{
		"domain", "email", "currency", // 3
		"jwt_secret", "jwt_secret_expire_hours", // 2
		"stripe_secret_key", "stripe_webhook_secret_key", // 2
		"social_facebook", "social_instagram", "social_twitter", "social_dribbble", "social_github", // 5
		"smtp_host", "smtp_port", "smtp_username", "smtp_password", "smtp_encryption", // 5
	}

	query := fmt.Sprintf("SELECT key, value FROM setting WHERE key IN (%s)", strings.Repeat("?, ", len(keys)-1)+"?")
	rows, err := q.DB.Query(query, keys...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fieldMap := map[string]interface{}{
		"domain":                    &settings.Main.Domain,
		"email":                     &settings.Main.Email,
		"currency":                  &settings.Main.Currency,
		"jwt_secret":                &settings.Main.JWT.Secret,
		"jwt_secret_expire_hours":   &settings.Main.JWT.ExpireHours,
		"stripe_secret_key":         &settings.Stripe.SecretKey,
		"stripe_webhook_secret_key": &settings.Stripe.WebhookSecretKey,
		"social_facebook":           &settings.Social.Facebook,
		"social_instagram":          &settings.Social.Instagram,
		"social_twitter":            &settings.Social.Twitter,
		"social_dribbble":           &settings.Social.Dribbble,
		"social_github":             &settings.Social.Github,
		"smtp_host":                 &settings.Mail.Host,
		"smtp_port":                 &settings.Mail.Port,
		"smtp_username":             &settings.Mail.Username,
		"smtp_password":             &settings.Mail.Password,
		"smtp_encryption":           &settings.Mail.Encryption,
	}

	for rows.Next() {
		var key, value string
		err := rows.Scan(&key, &value)
		if err != nil {
			return nil, err
		}

		fieldPtr, ok := fieldMap[key]
		if !ok {
			continue
		}

		switch v := fieldPtr.(type) {
		case *string:
			*v = value
		case *int:
			vInt, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			*v = vInt
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return settings, nil
}

// UpdateSettings is ...
func (q *SettingQueries) UpdateSettings(settings *models.Setting, section string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tx, err := q.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `UPDATE setting SET value = ? WHERE key = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	sectionSettings := make(map[string]any)

	switch section {
	case "main":
		sectionSettings = map[string]any{
			"domain":                  settings.Main.Domain,
			"email":                   settings.Main.Email,
			"currency":                settings.Main.Currency,
			"jwt_secret":              settings.Main.JWT.Secret,
			"jwt_secret_expire_hours": settings.Main.JWT.ExpireHours,
		}
	case "password":
		var passwordHash string
		if err := q.DB.QueryRow(`SELECT value FROM setting WHERE key = 'password'`).Scan(&passwordHash); err != nil {
			return errors.New("not found user")
		}
		compareUserPassword := security.ComparePasswords(passwordHash, settings.Password.Old)
		if !compareUserPassword {
			return errors.New("wrong password")
		}

		sectionSettings = map[string]any{
			"password": security.GeneratePassword(settings.Password.New),
		}

	case "stripe":
		sectionSettings = map[string]any{
			"stripe_secret_key":         settings.Stripe.SecretKey,
			"stripe_webhook_secret_key": settings.Stripe.WebhookSecretKey,
		}
	case "social":
		sectionSettings = map[string]any{
			"social_facebook":  settings.Social.Facebook,
			"social_instagram": settings.Social.Instagram,
			"social_twitter":   settings.Social.Twitter,
			"social_dribbble":  settings.Social.Dribbble,
			"social_github":    settings.Social.Github,
		}
	case "mail":
		sectionSettings = map[string]any{
			"smtp_host":       settings.Mail.Host,
			"smtp_port":       settings.Mail.Port,
			"smtp_username":   settings.Mail.Username,
			"smtp_password":   settings.Mail.Password,
			"smtp_encryption": settings.Mail.Encryption,
		}

	default:
		return errors.New("not found")
	}

	for key, value := range sectionSettings {
		if _, err := stmt.ExecContext(ctx, value, key); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// IsInstalled is ...
func (q *SettingQueries) IsInstalled() bool {
	var installed bool
	q.DB.QueryRow(`SELECT value FROM setting WHERE key = 'installed'`).Scan(&installed)
	return installed
}

// GetDomain is ...
func (q *SettingQueries) GetDomain() string {
	var domain string
	q.DB.QueryRow(`SELECT value FROM setting WHERE key = 'domain'`).Scan(&domain)
	return domain
}

// GetCurrency is ...
func (q *SettingQueries) GetCurrency() string {
	var currency string
	q.DB.QueryRow(`SELECT value FROM setting WHERE key = 'currency'`).Scan(&currency)
	return currency
}

// CheckSubdomain is ...
func (q *SettingQueries) CheckSubdomain(name string) bool {
	var id int
	err := q.DB.QueryRow(`SELECT id FROM domain WHERE name = ?`, name).Scan(&id)
	return err == nil
}

// GetSession is ...
func (q *SettingQueries) GetSession(key string) (string, error) {
	var value string
	err := q.DB.QueryRow(`SELECT value FROM session WHERE key = ?`, key).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}

// AddSession is ...
func (q *SettingQueries) AddSession(key, value string, expires int64) error {
	_, err := q.DB.Exec(`INSERT INTO session (key, value, expires) VALUES (?, ?, ?)`, key, value, expires)
	if err != nil {
		return err
	}
	return nil
}

// UpdateSession is ...
func (q *SettingQueries) UpdateSession(key, value string, expires int64) error {
	_, err := q.DB.Exec(`UPDATE session SET value = ?, expires = ? WHERE key = ? `, value, expires, key)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSession is ...
func (q *SettingQueries) DeleteSession(key string) error {
	_, err := q.DB.Exec(`DELETE FROM session WHERE key = ?`, key)
	if err != nil {
		return err
	}
	return nil
}

// SettingJWT is settings
func (q *SettingQueries) SettingJWT() (*jwtutil.Setting, error) {
	settings := &jwtutil.Setting{}

	query := `SELECT key, value FROM setting WHERE key IN (?, ?)`
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
func (q *SettingQueries) SettingStripe() (*models.Setting, error) {
	settings := &models.Setting{}

	query := `SELECT key, value FROM setting WHERE key IN (?, ?, ?)`
	rows, err := q.DB.Query(query, "stripe_secret_key", "stripe_webhook_secret_key", "domain")
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
			settings.Stripe.SecretKey = value
		case "stripe_webhook_secret_key":
			settings.Stripe.WebhookSecretKey = value
		case "domain":
			settings.Main.Domain = fmt.Sprintf("https://%s", value)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return settings, nil
}

// ListSocials is ...
func (q *SettingQueries) ListSocials() (*models.Social, error) {
	socials := &models.Social{}

	keys := []any{"social_facebook", "social_instagram", "social_twitter", "social_dribbble", "social_github"}

	query := fmt.Sprintf("SELECT key, value FROM setting WHERE key IN (%s)", strings.Repeat("?, ", len(keys)-1)+"?")
	rows, err := q.DB.Query(query, keys...)
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

		if value != "" {
			switch key {
			case "social_facebook":
				socials.Facebook = "https://facebook.com/" + value
			case "social_instagram":
				socials.Instagram = "https://instagram.com/" + value
			case "social_twitter":
				socials.Twitter = "https://twitter.com/@" + value
			case "social_dribbble":
				socials.Dribbble = "https://dribbble.com/" + value
			case "social_github":
				socials.Github = "https://github.com/" + value
			}
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return socials, nil
}
