package models

type Cart struct {
	Core
	Email         string        `json:"email"`
	Name          string        `json:"name"`
	Cart          []CartProduct `json:"cart"`
	AmountTotal   int64         `json:"amount_total"`
	Currency      string        `json:"currency"`
	PaymentID     string        `json:"payment_id"`
	PaymentStatus string        `json:"payment_status"`
}

type CartProduct struct {
	ProductID string `json:"id"`
	Quantity  int    `json:"quantity"`
}
