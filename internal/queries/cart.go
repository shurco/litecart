package queries

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/pkg/errors"
)

// CartQueries is a struct that embeds a pointer to an sql.DB.
// This allows for direct access to all the methods of sql.DB through CartQueries.
type CartQueries struct {
	*sql.DB
}

// PaymentList retrieves the status of different payment methods from the database.
func (q *CartQueries) PaymentList(ctx context.Context) (map[string]bool, error) {
	payments := map[string]bool{}
	keys := []any{
		"stripe_active", "paypal_active", "spectrocoin_active",
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

// Carts retrieves a list of carts from the database.
func (q *CartQueries) Carts(ctx context.Context) ([]*models.Cart, error) {
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

	rows, err := q.DB.QueryContext(ctx, query)
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

		cart.Email = email.String
		cart.PaymentID = paymentID.String
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

// Cart retrieves a cart from the database using the provided cartId.
func (q *CartQueries) Cart(ctx context.Context, cartId string) (*models.Cart, error) {
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

	var email, paymentID sql.NullString
	var created, updated sql.NullInt64
	cart := &models.Cart{}

	err := q.DB.QueryRowContext(ctx, query, cartId).
		Scan(
			&cart.ID,
			&email,
			&cart.AmountTotal,
			&cart.Currency,
			&paymentID,
			&cart.PaymentStatus,
			&created,
			&updated,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrProductNotFound
		}
		return nil, err
	}

	cart.Email = email.String
	cart.PaymentID = paymentID.String
	if created.Valid {
		cart.Created = created.Int64
	}

	if updated.Valid {
		cart.Updated = updated.Int64
	}

	return cart, nil
}

// AddCart inserts a new cart into the database.
func (q *CartQueries) AddCart(ctx context.Context, cart *models.Cart) error {
	byteCart, err := json.Marshal(cart.Cart)
	if err != nil {
		return err
	}

	query := `INSERT INTO cart (id, email, cart, amount_total, currency, payment_status, payment_system) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err = q.DB.ExecContext(ctx, query, cart.ID, cart.Email, string(byteCart), cart.AmountTotal, cart.Currency, cart.PaymentStatus, cart.PaymentSystem)
	return err
}

// UpdateCart updates the cart details in the database.
func (q *CartQueries) UpdateCart(ctx context.Context, cart *models.Cart) error {
	var (
		args []interface{}
		sql  strings.Builder
	)

	sql.WriteString("UPDATE cart SET ")

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

	_, err := q.DB.ExecContext(ctx, sql.String(), args...)
	return err
}
