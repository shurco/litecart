package queries

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/pkg/errors"
	"github.com/shurco/litecart/pkg/security"
)

// SettingQueries wraps a sql.DB connection allowing for easy querying and interaction
// with the database related to application settings.
type SettingQueries struct {
	*sql.DB
}

// createFieldMap generates a map of fields based on the type of settings.
func (q *SettingQueries) createFieldMap(settings any) map[string]any {
	switch s := settings.(type) {
	case *models.Main:
		return map[string]any{
			"site_name": &s.SiteName,
			"domain":    &s.Domain,
			"email":     &s.Email, // TODO: move to authorization section
			"currency":  &s.Currency,
		}
	case *models.JWT:
		return map[string]any{
			"jwt_secret":              &s.Secret,
			"jwt_secret_expire_hours": &s.ExpireHours,
		}
	case *models.Social:
		return map[string]any{
			"social_facebook":  &s.Facebook,
			"social_instagram": &s.Instagram,
			"social_twitter":   &s.Twitter,
			"social_dribbble":  &s.Dribbble,
			"social_github":    &s.Github,
		}
	case *models.Stripe:
		return map[string]any{
			"stripe_secret_key": &s.SecretKey,
			"stripe_active":     &s.Active,
		}
	case *models.Paypal:
		return map[string]any{
			"paypal_client_id":  &s.ClientID,
			"paypal_secret_key": &s.SecretKey,
			"paypal_active":     &s.Active,
		}
	case *models.Spectrocoin:
		return map[string]any{
			"spectrocoin_merchant_id": &s.MerchantID,
			"spectrocoin_project_id":  &s.ProjectID,
			"spectrocoin_private_key": &s.PrivateKey,
			"spectrocoin_active":      &s.Active,
		}
	case *models.Webhook:
		return map[string]any{
			"webhook_url": &s.Url,
		}
	case *models.SMTP:
		return map[string]any{
			"smtp_host":       &s.Host,
			"smtp_port":       &s.Port,
			"smtp_username":   &s.Username,
			"smtp_password":   &s.Password,
			"smtp_encryption": &s.Encryption,
		}
	default:
		return nil
	}
}

// GetSetting retrieves settings based on the provided `settings` struct, populating it with values from the database.
func (q *SettingQueries) GetSetting(ctx context.Context, settings any) (any, error) {
	fieldMap := q.createFieldMap(settings)

	keys := make([]any, 0, len(fieldMap))
	for k := range fieldMap {
		keys = append(keys, k)
	}

	query := fmt.Sprintf("SELECT key, value FROM setting WHERE key IN (%s)", strings.Repeat("?, ", len(keys)-1)+"?")
	rows, err := q.DB.QueryContext(ctx, query, keys...)
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

		if fieldPtr, ok := fieldMap[key]; ok {
			switch ptr := fieldPtr.(type) {
			case *string:
				*ptr = value
			case *bool:
				bValue, err := strconv.ParseBool(value)
				if err != nil {
					return nil, err
				}
				*ptr = bValue
			case *int:
				iValue, err := strconv.Atoi(value)
				if err != nil {
					return nil, err
				}
				*ptr = iValue
			}
		}
	}

	return settings, nil
}

// UpdateSetting updates the settings in the database using a transaction.
// It takes a context and a settings object of any type as arguments.
func (q *SettingQueries) UpdateSetting(ctx context.Context, settings any) error {
	fieldMap := q.createFieldMap(settings)

	tx, err := q.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `UPDATE setting SET value = ? WHERE key = ?`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for key, value := range fieldMap {
		if _, err = stmt.ExecContext(ctx, value, key); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// UpdatePassword updates the current user's password in the database.
func (q *SettingQueries) UpdatePassword(ctx context.Context, password *models.Password) error {
	var passwordHash string
	query := `SELECT value FROM setting WHERE key = 'password'`
	if err := q.DB.QueryRowContext(ctx, query).Scan(&passwordHash); err != nil {
		return errors.ErrUserNotFound
	}
	compareUserPassword := security.ComparePasswords(passwordHash, password.Old)
	if !compareUserPassword {
		return errors.ErrWrongPassword
	}

	query = `UPDATE setting SET value = ? WHERE key = 'password'`
	_, err := q.DB.ExecContext(ctx, query, security.GeneratePassword(password.New))
	return err
}

// GetSettingByKey retrieves a setting by its key from the database.
// It accepts a context for cancellation and a string representing the key of the setting.
// Returns a pointer to a SettingName model if found, or an error if not found or any other issue occurs.
func (q *SettingQueries) GetSettingByKey(ctx context.Context, key string) (*models.SettingName, error) {
	setting := &models.SettingName{
		Key: key,
	}

	query := `SELECT id, value FROM setting WHERE key = ?`
	err := q.DB.QueryRowContext(ctx, query, key).Scan(&setting.ID, &setting.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrSettingNotFound
		}
		return nil, err
	}
	return setting, nil
}

// UpdateSettingByKey updates the value of a setting in the database based on the provided key.
func (q *SettingQueries) UpdateSettingByKey(ctx context.Context, setting *models.SettingName) error {
	query := `UPDATE setting SET value = ? WHERE key = ? `
	_, err := q.DB.ExecContext(ctx, query, setting.Value, setting.Key)
	return err
}
