package litepay

import (
	"fmt"
	"strings"
)

type dummy struct {
	Cfg
	successURL string
}

// Dummy initializes a dummy payment provider for testing purposes.
// This provider always returns successful payments immediately without
// any actual payment processing.
//
// WARNING: This provider should ONLY be used for:
//   - Testing and development
//   - Free items (amount = 0)
//
// The application should validate that the cart total is 0 before
// allowing the use of this provider in production.
//
// Returns:
//   - LitePay: A configured dummy payment provider
//
// Supported currencies: EUR, USD, GBP, AUD, CAD, JPY, CNY, SEK
//
// Example:
//
//	pay := litepay.New(callbackURL, successURL, cancelURL)
//	dummy := pay.Dummy()
//	payment, err := dummy.Pay(cart) // Always succeeds with status PAID
func (c Cfg) Dummy() LitePay {
	c.paymentSystem = DUMMY
	c.currency = []string{"EUR", "USD", "GBP", "AUD", "CAD", "JPY", "CNY", "SEK"}
	return &dummy{
		Cfg:        c,
		successURL: c.successURL,
	}
}

func (c *dummy) Pay(cart Cart) (*Payment, error) {
	var amountTotal int
	for _, item := range cart.Items {
		amountTotal += item.PriceData.UnitAmount * item.Quantity
	}

	checkout := &Payment{
		AmountTotal:   amountTotal,
		Currency:      strings.ToUpper(cart.Currency),
		Status:        PAID,
		URL:           fmt.Sprintf("%s/?payment_system=%s&cart_id=%s", c.successURL, c.paymentSystem, cart.ID),
		PaymentSystem: c.paymentSystem,
	}

	return checkout, nil
}

func (c *dummy) Checkout(payment *Payment, session string) (*Payment, error) {
	payment.Status = PAID
	payment.MerchantID = "dummy_" + payment.CartID
	return payment, nil
}
