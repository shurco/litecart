package queries

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/shurco/litecart/internal/models"
)

// CartQueries is ...
type CartQueries struct {
	*sql.DB
}

// Checkouts is ...
func (q *CartQueries) Checkouts() ([]*models.Checkout, error) {
	checkouts := []*models.Checkout{}

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
		checkout := &models.Checkout{}

		err := rows.Scan(
			&checkout.ID,
			&email,
			&name,
			&checkout.AmountTotal,
			&checkout.Currency,
			&paymentID,
			&checkout.PaymentStatus,
			&checkout.Created,
			&updated,
		)
		if err != nil {
			return nil, err
		}

		if email.Valid {
			checkout.Email = email.String
		}

		if name.Valid {
			checkout.Name = name.String
		}

		if paymentID.Valid {
			checkout.PaymentID = paymentID.String
		}

		if updated.Valid {
			checkout.Updated = updated.Int64
		}

		checkouts = append(checkouts, checkout)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return checkouts, nil
}

// AddCart is ...
func (q *CartQueries) AddCart(cart *models.Checkout) error {
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
func (q *CartQueries) UpdateCart(cart *models.Checkout) error {
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

	return nil
}
