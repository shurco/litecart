package litepay

import validation "github.com/go-ozzo/ozzo-validation/v4"

// Cart represents a shopping cart with items to be purchased.
type Cart struct {
	ID       string `json:"id"`       // Unique cart identifier (15 characters)
	Currency string `json:"currency"` // ISO currency code (e.g., "USD", "EUR")
	Items    []Item `json:"items"`    // List of items in the cart
}

// Item represents a single item in the shopping cart.
type Item struct {
	PriceData Price `json:"price"`    // Price information for this item
	Quantity  int   `json:"quantity"` // Number of items
}

// Price contains pricing information for a product.
type Price struct {
	UnitAmount int     `json:"init_amount"` // Price in smallest currency unit (cents/kopeks)
	Product    Product `json:"product"`     // Product details
}

// Product represents the product being purchased.
type Product struct {
	Name        string   `json:"name"`                  // Product name
	Description string   `json:"description,omitempty"` // Optional product description
	Images      []string `json:"images"`                // List of product image URLs
}

// Payment represents a payment transaction.
type Payment struct {
	PaymentSystem PaymentSystem `json:"provider"`     // Payment provider used
	MerchantID    string        `json:"merchant_id"`  // Transaction ID from the provider
	CartID        string        `json:"cart_id"`      // Associated cart ID
	AmountTotal   int           `json:"amount_total"` // Total amount in smallest currency unit
	Currency      string        `json:"currency"`     // ISO currency code
	Status        Status        `json:"status"`       // Current payment status
	URL           string        `json:"url,omitempty"` // Checkout URL to redirect user (if applicable)
	Coin          *Coin         `json:"coin,omitempty"` // Cryptocurrency payment details (if applicable)
}

// Validate validates the Payment structure.
// It ensures that the CartID has exactly 15 characters.
func (v Payment) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.CartID, validation.Length(15, 15)),
	)
}

// Coin represents cryptocurrency payment information.
// Used primarily with SpectroCoin provider.
type Coin struct {
	AmountTotal float64 `json:"amount_total"` // Amount in cryptocurrency
	Currency    string  `json:"currency"`     // Cryptocurrency code (e.g., "BTC", "ETH")
}
