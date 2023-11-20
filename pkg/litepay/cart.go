package litepay

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Cart struct {
	ID       string `json:"id"`
	Currency string `json:"currency"`
	Items    []Item `json:"items"`
}

type Item struct {
	PriceData Price `json:"price"`
	Quantity  int   `json:"quantity"`
}

type Price struct {
	UnitAmount int     `json:"init_amount"`
	Product    Product `json:"product"`
}

type Product struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Images      []string `json:"images"`
}

type Payment struct {
	PaymentSystem PaymentSystem `json:"provider"`
	MerchantID    string        `json:"merchant_id"`
	CartID        string        `json:"cart_id"`
	AmountTotal   int           `json:"amount_total"`
	Currency      string        `json:"currency"`
	Status        Status        `json:"status"`
	URL           string        `json:"url,omitempty"`
	Coin          *Coin         `json:"coin,omitempty"`
}

// Validate is ...
func (v Payment) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.CartID, validation.Length(15, 15)),
	)
}

type Coin struct {
	AmountTotal float64 `json:"amount_total"`
	Currency    string  `json:"currency"`
}
