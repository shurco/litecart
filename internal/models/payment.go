package models

type Checkout struct {
	ID            string             `json:"id"`
	Email         string             `json:"email"`
	Name          string             `json:"name"`
	Cart          []CheckoutLineItem `json:"cart"`
	AmountTotal   int64              `json:"amount_total"`
	Currency      string             `json:"currency"`
	PaymentID     string             `json:"payment_id"`
	PaymentStatus string             `json:"payment_status"`
	Created       int64              `json:"created"`
	Updated       int64              `json:"updated,omitempty"`
}

type CheckoutLineItem struct {
	Price    string `json:"price"`
	Quantity int    `json:"quantity"`
}
