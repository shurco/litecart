package queries

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/shurco/litecart/internal/models"
)

// CartQueries is ...
type CartQueries struct {
	*sql.DB
}

// PaymentList is ...
func (q *CartQueries) PaymentList() (map[string]bool, error) {
	payments := map[string]bool{}
	keys := []any{
		"stripe_active", "spectrocoin_active",
	}

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

		vBool, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		name := strings.ReplaceAll(key, "_active", "")
		payments[name] = vBool
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return payments, nil
}

// Carts is ...
func (q *CartQueries) Carts() ([]*models.Cart, error) {
	carts := []*models.Cart{}

	query := `
	SELECT 
		id, 
		email, 
		amount_total,
		currency,
		payment_id,
		payment_status,
		payment_system,
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
		var email, paymentID sql.NullString
		var updated sql.NullInt64
		cart := &models.Cart{}

		err := rows.Scan(
			&cart.ID,
			&email,
			&cart.AmountTotal,
			&cart.Currency,
			&paymentID,
			&cart.PaymentStatus,
			&cart.PaymentSystem,
			&cart.Created,
			&updated,
		)
		if err != nil {
			return nil, err
		}

		if email.Valid {
			cart.Email = email.String
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

	var email, paymentID sql.NullString
	var updated sql.NullInt64
	cart := &models.Cart{}

	err = rows.Scan(
		&cart.ID,
		&email,
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

	_, err = q.DB.ExecContext(context.TODO(), `INSERT INTO cart (id, email, cart, amount_total, currency, payment_status, payment_system) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		cart.ID,
		cart.Email,
		string(byteCart),
		cart.AmountTotal,
		cart.Currency,
		cart.PaymentStatus,
		cart.PaymentSystem,
	)
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

	return nil
}
