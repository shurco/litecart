package litepay

import (
	"fmt"
	"strings"
)

type dummy struct {
	Cfg
	successURL string
}

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
