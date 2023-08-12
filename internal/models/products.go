package models

// Products is ...
type Products struct {
	Total    int       `json:"total"`
	Currency string    `json:"currency"`
	Products []Product `json:"products"`
}

// Product is ...
type Product struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Images      []Images          `json:"images,omitempty"`
	URL         string            `json:"url"`
	Amount      int               `json:"amount"`
	Currency    string            `json:"currency,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	Attributes  []string          `json:"attributes,omitempty"`
	Active      bool              `json:"active"`
	Created     int64             `json:"created"`
	Updated     int64             `json:"updated,omitempty"`
}

// Images is ...
type Images struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Ext  string `json:"ext"`
}
