package queries

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/shurco/litecart/internal/mailer"
	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/pkg/errors"
)

// CartQueries is ...
type CartQueries struct {
	*sql.DB
}

// Carts is ...
func (q *CartQueries) Carts() ([]*models.Cart, error) {
	carts := []*models.Cart{}

	query := `
	SELECT 
		id, 
		email, 
		name, 
		amount_total,
		currency,
		payment_id,
		payment_status,
		strftime('%s', created),
		strftime('%s', updated)
	FROM cart
	`

	rows, err := q.DB.QueryContext(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var email, name, paymentID sql.NullString
		var updated sql.NullInt64
		cart := &models.Cart{}

		err := rows.Scan(
			&cart.ID,
			&email,
			&name,
			&cart.AmountTotal,
			&cart.Currency,
			&paymentID,
			&cart.PaymentStatus,
			&cart.Created,
			&updated,
		)
		if err != nil {
			return nil, err
		}

		if email.Valid {
			cart.Email = email.String
		}

		if name.Valid {
			cart.Name = name.String
		}

		if paymentID.Valid {
			cart.PaymentID = paymentID.String
		}

		if updated.Valid {
			cart.Updated = updated.Int64
		}

		carts = append(carts, cart)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return carts, nil
}

// Carts is ...
func (q *CartQueries) Cart(cartId string) (*models.Cart, error) {

	query := `
	SELECT 
    id, 
    email, 
    name, 
    amount_total,
    currency,
    payment_id,
    payment_status,
    strftime('%s', created),
    strftime('%s', updated)
	FROM cart
	WHERE id = ?
	`

	rows, err := q.DB.QueryContext(context.TODO(), query, cartId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var email, name, paymentID sql.NullString
	var updated sql.NullInt64
	cart := &models.Cart{}

	err = rows.Scan(
		&cart.ID,
		&email,
		&name,
		&cart.AmountTotal,
		&cart.Currency,
		&paymentID,
		&cart.PaymentStatus,
		&cart.Created,
		&updated,
	)
	if err != nil {
		return nil, err
	}

	if email.Valid {
		cart.Email = email.String
	}

	if name.Valid {
		cart.Name = name.String
	}

	if paymentID.Valid {
		cart.PaymentID = paymentID.String
	}

	if updated.Valid {
		cart.Updated = updated.Int64
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cart, nil
}

// AddCart is ...
func (q *CartQueries) AddCart(cart *models.Cart) error {
	byteCart, err := json.Marshal(cart.Cart)
	if err != nil {
		return err
	}

	_, err = q.DB.ExecContext(context.TODO(), `INSERT INTO cart (id, cart, amount_total, currency, payment_status) VALUES (?, ?, ?, ?, ?)`, cart.ID, string(byteCart), cart.AmountTotal, cart.Currency, cart.PaymentStatus)
	if err != nil {
		return err
	}

	return nil
}

// UpdateCart is ...
func (q *CartQueries) UpdateCart(cart *models.Cart) error {
	var (
		args []interface{}
		sql  strings.Builder
	)

	sql.WriteString("UPDATE cart SET ")

	if cart.Email != "" {
		sql.WriteString("email = ?, ")
		args = append(args, cart.Email)
	}

	if cart.Name != "" {
		sql.WriteString("name = ?, ")
		args = append(args, cart.Name)
	}

	if cart.PaymentID != "" {
		sql.WriteString("payment_id = ?, ")
		args = append(args, cart.PaymentID)
	}

	if cart.PaymentStatus != "" {
		sql.WriteString("payment_status = ?, ")
		args = append(args, cart.PaymentStatus)
	}

	sql.WriteString("updated = datetime('now') WHERE id = ?")
	args = append(args, cart.ID)

	if _, err := q.DB.ExecContext(context.TODO(), sql.String(), args...); err != nil {
		return err
	}

	if cart.Email != "" && cart.Name != "" && cart.PaymentStatus == "paid" {
		if err := q.CartSendMail(cart.ID); err != nil {
			return err
		}
	}

	return nil
}

// CartSendMail
func (q *CartQueries) CartSendMail(cartID string) error {
	mail := &models.Mail{}
	products := []models.CartProduct{}
	keys := []models.Data{}
	var name, cartJSON, letter string

	err := q.DB.QueryRowContext(context.TODO(), `SELECT email, name, cart FROM cart WHERE payment_status = 'paid' AND id = ?`, cartID).Scan(&mail.To, &name, &cartJSON)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.ErrPageNotFound
		}
		return err
	}

	if err := json.Unmarshal([]byte(cartJSON), &products); err != nil {
		return err
	}

	tx, err := q.DB.BeginTx(context.TODO(), nil)
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
		err := tx.QueryRowContext(context.TODO(), `SELECT digital FROM product WHERE id = ?`, cart.ProductID).Scan(&digitalType)
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.ErrPageNotFound
			}
			return err
		}

		if digitalType == "file" {
			rows, err := tx.QueryContext(context.TODO(), `SELECT id, name, ext, orig_name FROM digital_file WHERE product_id = ?`, cart.ProductID)
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
			err := tx.QueryRowContext(context.TODO(), `SELECT id, content FROM digital_data WHERE cart_id = ?`, cartID).Scan(&key.ID, &key.Content)
			if err != nil && err != sql.ErrNoRows {
				return err
			}
			if err == sql.ErrNoRows {
				err = tx.QueryRowContext(context.TODO(), `SELECT id, content FROM digital_data WHERE cart_id IS NULL AND product_id = ? LIMIT 1`, cart.ProductID).Scan(&key.ID, &key.Content)
				if err != nil {
					if err == sql.ErrNoRows {
						return errors.ErrPageNotFound
					}
					return err
				}
				if _, err := tx.ExecContext(context.TODO(), `UPDATE digital_data SET cart_id = ? WHERE id = ?`, cartID, key.ID); err != nil {
					return err
				}
			}

			keys = append(keys, key)
		}
	}

	if err := tx.QueryRowContext(context.TODO(), `SELECT value FROM setting WHERE key = 'email'`).Scan(&mail.From); err != nil {
		return err
	}

	if err := tx.QueryRowContext(context.TODO(), `SELECT value FROM setting WHERE key = 'mail_letter_purchase'`).Scan(&letter); err != nil {
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
		"Customer_Name": name,
		"Purchases":     purchases.String(),
		"Admin_Email":   mail.From,
	}

	setting := SettingQueries{q.DB}
	smtpSetting, err := setting.SettingMail()
	if err != nil {
		return err
	}

	if err := mailer.SendMail(smtpSetting, mail); err != nil {
		return err
	}

	return nil
}
