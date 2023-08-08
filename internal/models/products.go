package models

// Products is ...
type Products struct {
	Total    int       `json:"total"`
	Products []Product `json:"products"`
}

// Product is ...
type Product struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Images      []Images          `json:"images"`
	URL         string            `json:"url"`
	Metadata    map[string]string `json:"metadata"`
	Attributes  []string          `json:"attributes"`
	Stripe      Stripe            `json:"stripe"`
	Active      bool              `json:"active"`
	Created     int64             `json:"created"`
	Updated     int64             `json:"updated"`
}

// Images is ...
type Images struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Ext  string `json:"ext"`
}

// Stripe is ...
type Stripe struct {
	Product StripeProduct `json:"product"`
	Price   StripePrice   `json:"price"`
}

// StripeProduct is ...
type StripeProduct struct {
	ID    string `json:"id"`
	Valid bool   `json:"valid"`
}

// StripePrice is ...
type StripePrice struct {
	ID       string `json:"id"`
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

// StripeInfo is ...
type StripeInfo struct {
	ProductID string `json:"product_id"`
	PriceID   string `json:"price_id"`
}
