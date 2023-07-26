package models

type Product struct {
	ID          string            `json:"id"`
	Name        string            `json:"name" validate:"required"`
	Description string            `json:"description" validate:"required"`
	Price       Price             `json:"price"`
	Images      []string          `json:"images"`
	URL         string            `json:"url"`
	Metadata    map[string]string `json:"metadata" validate:"required"`
	Attributes  []string          `json:"attributes" validate:"required"`
	Created     int64             `json:"created"`
	Updated     int64             `json:"updated"`
}

type Price struct {
	ID       string `json:"id"`
	StripeID string `json:"stripe_id" validate:"required"`
	Currency string `json:"currency" validate:"required"`
	Amount   int    `json:"amount" validate:"required"`
}
