package mailer

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/errors"
	"github.com/shurco/litecart/pkg/litepay"
)

// SendTestLetter is ...
func SendTestLetter(letterName string) error {
	db := queries.DB().SettingQueries

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_smtpSetting, err := db.GetSetting(ctx, &models.SMTP{})
	if err != nil {
		return err
	}
	smtpSetting := _smtpSetting.(*models.SMTP)

	settingEmail, err := db.GetSettingByKey(ctx, "email")
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
			"Payment_URL":    "https://payment.com/order/1234567890",
			"Admin_Email":    "admin@mail.com",
			"Site_Name":      "Site name",
			"Amount_Payment": "21.00 USD",
		},
	}

	if letterName != "smtp" {
		settingLetter, err := db.GetSettingByKey(ctx, letterName)
		if err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(settingLetter.Value.(string)), &letter.Letter); err != nil {
			return err
		}
	}

	if err := SendMail(smtpSetting, letter); err != nil {
		return err
	}

	return nil
}

// SendPrepaymentLetter is ...
func SendPrepaymentLetter(email, amountPayment, paymentURL string) error {
	db := queries.DB().SettingQueries

	mail := &models.Mail{
		To: email,
		Data: map[string]string{
			"Payment_URL":    paymentURL,
			"Site_Name":      "",
			"Amount_Payment": amountPayment,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.DB.QueryContext(ctx, `SELECT key, value FROM setting WHERE key IN ('site_name', 'email','mail_letter_payment')`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var key, value string
		err := rows.Scan(&key, &value)
		if err != nil {
			return err
		}

		switch key {
		case "site_name":
			mail.Data["Site_Name"] = value
		case "email":
			mail.From = value
		case "mail_letter_payment":
			if err := json.Unmarshal([]byte(value), &mail.Letter); err != nil {
				return err
			}
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	_smtpSetting, err := db.GetSetting(ctx, &models.SMTP{})
	if err != nil {
		return err
	}
	smtpSetting := _smtpSetting.(*models.SMTP)

	if err := SendMail(smtpSetting, mail); err != nil {
		return err
	}

	return nil
}

// SendCartLetter
func SendCartLetter(cartID string) error {
	db := queries.DB().SettingQueries

	mail := &models.Mail{}
	products := []models.CartProduct{}
	keys := []models.Data{}
	var cartJSON, letter string

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := db.QueryRowContext(ctx, `SELECT email, cart FROM cart WHERE payment_status = ? AND id = ?`, litepay.PAID, cartID).Scan(&mail.To, &cartJSON)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.ErrPageNotFound
		}
		return err
	}

	if err := json.Unmarshal([]byte(cartJSON), &products); err != nil {
		return err
	}

	tx, err := db.DB.BeginTx(ctx, nil)
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

	for _, cart := range products {
		var digitalType string
		err := tx.QueryRowContext(ctx, `SELECT digital FROM product WHERE id = ?`, cart.ProductID).Scan(&digitalType)
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.ErrPageNotFound
			}
			return err
		}

		if digitalType == "file" {
			rows, err := tx.QueryContext(ctx, `SELECT id, name, ext, orig_name FROM digital_file WHERE product_id = ?`, cart.ProductID)
			if err != nil {
				return err
			}
			defer rows.Close()

			for rows.Next() {
				file := models.File{}
				err := rows.Scan(
					&file.ID,
					&file.Name,
					&file.Ext,
					&file.OrigName,
				)
				if err != nil {
					return err
				}
				mail.Files = append(mail.Files, file)
			}
		}

		if digitalType == "data" {
			key := models.Data{}
			err := tx.QueryRowContext(ctx, `SELECT id, content FROM digital_data WHERE cart_id = ?`, cartID).Scan(&key.ID, &key.Content)
			if err != nil && err != sql.ErrNoRows {
				return err
			}
			if err == sql.ErrNoRows {
				err = tx.QueryRowContext(ctx, `SELECT id, content FROM digital_data WHERE cart_id IS NULL AND product_id = ? LIMIT 1`, cart.ProductID).Scan(&key.ID, &key.Content)
				if err != nil {
					if err == sql.ErrNoRows {
						return errors.ErrPageNotFound
					}
					return err
				}
				if _, err := tx.ExecContext(ctx, `UPDATE digital_data SET cart_id = ? WHERE id = ?`, cartID, key.ID); err != nil {
					return err
				}
			}

			keys = append(keys, key)
		}
	}

	if err := tx.QueryRowContext(ctx, `SELECT value FROM setting WHERE key = 'email'`).Scan(&mail.From); err != nil {
		return err
	}

	if err := tx.QueryRowContext(ctx, `SELECT value FROM setting WHERE key = 'mail_letter_purchase'`).Scan(&letter); err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(letter), &mail.Letter); err != nil {
		return err
	}

	var purchases strings.Builder
	count := 1

	if len(keys) > 0 {
		purchases.WriteString("Keys:\n")
		for _, key := range keys {
			purchases.WriteString(fmt.Sprintf("%v: %s\n", count, key.Content))
			count++
		}
	}

	if len(mail.Files) > 0 {
		purchases.WriteString("Files:\n")
		for _, file := range mail.Files {
			purchases.WriteString(fmt.Sprintf("%v: %s\n", count, file.OrigName))
			count++
		}
	}

	mail.Data = map[string]string{
		"Purchases":   purchases.String(),
		"Admin_Email": mail.From,
	}

	_smtpSetting, err := db.GetSetting(ctx, &models.SMTP{})
	if err != nil {
		return err
	}
	smtpSetting := _smtpSetting.(*models.SMTP)

	if err := SendMail(smtpSetting, mail); err != nil {
		return err
	}

	return nil
}
