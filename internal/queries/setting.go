package queries

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/shurco/litecart/internal/mailer"
	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/pkg/errors"
	"github.com/shurco/litecart/pkg/jwtutil"
	"github.com/shurco/litecart/pkg/security"
)

// SettingQueries is ...
type SettingQueries struct {
	*sql.DB
}

// Settings is ...
func (q *SettingQueries) Settings(private bool) (*models.Setting, error) {
	settings := new(models.Setting)
	keys := []any{
		"site_name", "domain", "email", "currency", // 3
		"social_facebook", "social_instagram", "social_twitter", "social_dribbble", "social_github", // 5
	}

	if private {
		keys = append(keys,
			"jwt_secret", "jwt_secret_expire_hours", // 2
			"stripe_secret_key", "stripe_webhook_secret_key", // 2
			"payment_webhook_url", // 1
			"smtp_host", "smtp_port", "smtp_username", "smtp_password", "smtp_encryption", // 5
		)
	}

	query := fmt.Sprintf("SELECT key, value FROM setting WHERE key IN (%s)", strings.Repeat("?, ", len(keys)-1)+"?")
	rows, err := q.DB.QueryContext(context.TODO(), query, keys...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fieldMap := map[string]interface{}{
		"site_name":                 &settings.Main.SiteName,
		"domain":                    &settings.Main.Domain,
		"email":                     &settings.Main.Email,
		"currency":                  &settings.Main.Currency,
		"jwt_secret":                &settings.Main.JWT.Secret,
		"jwt_secret_expire_hours":   &settings.Main.JWT.ExpireHours,
		"stripe_secret_key":         &settings.Stripe.SecretKey,
		"stripe_webhook_secret_key": &settings.Stripe.WebhookSecretKey,
		"payment_webhook_url":       &settings.Payment.WebhookUrl,
		"social_facebook":           &settings.Social.Facebook,
		"social_instagram":          &settings.Social.Instagram,
		"social_twitter":            &settings.Social.Twitter,
		"social_dribbble":           &settings.Social.Dribbble,
		"social_github":             &settings.Social.Github,
		"smtp_host":                 &settings.SMTP.Host,
		"smtp_port":                 &settings.SMTP.Port,
		"smtp_username":             &settings.SMTP.Username,
		"smtp_password":             &settings.SMTP.Password,
		"smtp_encryption":           &settings.SMTP.Encryption,
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
	defer func() {
		if p := recover(); p != nil || err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	stmt, err := tx.PrepareContext(ctx, `UPDATE setting SET value = ? WHERE key = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var sectionSettings map[string]any
	switch section {
	case "main":
		sectionSettings = map[string]any{
			"site_name":               settings.Main.SiteName,
			"domain":                  settings.Main.Domain,
			"email":                   settings.Main.Email,
			"currency":                settings.Main.Currency,
			"jwt_secret":              settings.Main.JWT.Secret,
			"jwt_secret_expire_hours": settings.Main.JWT.ExpireHours,
		}
	case "password":
		var passwordHash string
		if err := q.DB.QueryRowContext(context.TODO(), `SELECT value FROM setting WHERE key = 'password'`).Scan(&passwordHash); err != nil {
			return errors.ErrUserNotFound
		}
		compareUserPassword := security.ComparePasswords(passwordHash, settings.Password.Old)
		if !compareUserPassword {
			return errors.ErrWrongPassword
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
	case "payment":
		sectionSettings = map[string]any{
			"payment_webhook_url":  settings.Payment.WebhookUrl,
		}
	case "smtp":
		sectionSettings = map[string]any{
			"smtp_host":       settings.SMTP.Host,
			"smtp_port":       settings.SMTP.Port,
			"smtp_username":   settings.SMTP.Username,
			"smtp_password":   settings.SMTP.Password,
			"smtp_encryption": settings.SMTP.Encryption,
		}

	default:
		return errors.ErrSettingNotFound
	}

	for key, value := range sectionSettings {
		if _, err := stmt.ExecContext(ctx, value, key); err != nil {
			return err
		}
	}

	return nil
}

// SettingValueByKey is ...
func (q *SettingQueries) SettingValueByKey(key string) (*models.SettingName, error) {
	setting := &models.SettingName{
		Key: key,
	}
	err := q.DB.QueryRowContext(context.TODO(), `SELECT id, value FROM setting WHERE key = ?`, key).Scan(&setting.ID, &setting.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrSettingNotFound
		}
		return nil, err
	}
	return setting, nil
}

// UpdateSettingValueByKey is ...
func (q *SettingQueries) UpdateSettingValueByKey(setting *models.SettingName) error {
	_, err := q.DB.ExecContext(context.TODO(), `UPDATE setting SET value = ? WHERE key = ? `, setting.Value, setting.Key)
	if err != nil {
		return err
	}
	return nil
}

// IsInstalled is ...
func (q *SettingQueries) IsInstalled() bool {
	var installed bool
	q.DB.QueryRowContext(context.TODO(), `SELECT value FROM setting WHERE key = 'installed'`).Scan(&installed)
	return installed
}

// GetDomain is ...
func (q *SettingQueries) GetDomain() string {
	var domain string
	q.DB.QueryRowContext(context.TODO(), `SELECT value FROM setting WHERE key = 'domain'`).Scan(&domain)
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
	err := q.DB.QueryRowContext(context.TODO(), `SELECT id FROM domain WHERE name = ?`, name).Scan(&id)
	return err == nil
}

// GetSession is ...
func (q *SettingQueries) GetSession(key string) (string, error) {
	var value string
	expires := time.Now().Unix()
	err := q.DB.QueryRowContext(context.TODO(), `SELECT value FROM session WHERE key = ? AND expires > ?`, key, expires).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}

// AddSession is ...
func (q *SettingQueries) AddSession(key, value string, expires int64) error {
	_, err := q.DB.ExecContext(context.TODO(), `INSERT INTO session (key, value, expires) VALUES (?, ?, ?)`, key, value, expires)
	if err != nil {
		return err
	}
	return nil
}

// UpdateSession is ...
func (q *SettingQueries) UpdateSession(key, value string, expires int64) error {
	_, err := q.DB.ExecContext(context.TODO(), `UPDATE session SET value = ?, expires = ? WHERE key = ? `, value, expires, key)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSession is ...
func (q *SettingQueries) DeleteSession(key string) error {
	_, err := q.DB.ExecContext(context.TODO(), `DELETE FROM session WHERE key = ?`, key)
	if err != nil {
		return err
	}
	return nil
}

// SettingJWT is settings
func (q *SettingQueries) SettingJWT() (*jwtutil.Setting, error) {
	settings := &jwtutil.Setting{}

	query := `SELECT key, value FROM setting WHERE key IN (?, ?)`
	rows, err := q.DB.QueryContext(context.TODO(), query, "jwt_secret", "jwt_secret_expire_hours")
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

	query := `SELECT key, value FROM setting WHERE key IN (?, ?, ?, ?)`
	rows, err := q.DB.QueryContext(context.TODO(), query, "stripe_secret_key", "stripe_webhook_secret_key", "domain", "payment_webhook_url")
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
		case "payment_webhook_url":
			settings.Payment.WebhookUrl = value
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
	rows, err := q.DB.QueryContext(context.TODO(), query, keys...)
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

// SettingTestSMTP is ...
func (q *SettingQueries) SettingMail() (*models.SMTP, error) {
	setting := new(models.SMTP)
	keys := []any{"smtp_host", "smtp_port", "smtp_username", "smtp_password", "smtp_encryption"}

	query := fmt.Sprintf("SELECT key, value FROM setting WHERE key IN (%s)", strings.Repeat("?, ", len(keys)-1)+"?")
	rows, err := q.DB.QueryContext(context.TODO(), query, keys...)
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
			case "smtp_host":
				setting.Host = value
			case "smtp_port":
				SMTPPort, _ := strconv.Atoi(value)
				setting.Port = SMTPPort
			case "smtp_username":
				setting.Username = value
			case "smtp_password":
				setting.Password = value
			case "smtp_encryption":
				setting.Encryption = value
			}
		}
	}

	return setting, nil
}

// SettingTestLetter is ...
func (q *SettingQueries) SettingTestLetter(letterName string) error {
	smtpSetting, err := q.SettingMail()
	if err != nil {
		return err
	}

	settingEmail, err := q.SettingValueByKey("email")
	if err != nil {
		return err
	}
	emailAdmin := settingEmail.Value.(string)

	letter := &models.Mail{
		From: emailAdmin,
		To:   emailAdmin,
		Letter: models.Letter{
			Subject: "litecart test smtp settings",
			Text:    "test message",
		},
		Data: map[string]string{
			"Customer_Name": "Customer Name",
			"Admin_Email":   "admin@mail.com",
		},
	}

	if letterName != "smtp" {
		settingLetter, err := q.SettingValueByKey(letterName)
		if err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(settingLetter.Value.(string)), &letter.Letter); err != nil {
			return err
		}
	}

	if err := mailer.SendMail(smtpSetting, letter); err != nil {
		return err
	}

	return nil
}
