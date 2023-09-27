package models

// Core is ...
type Core struct {
	ID      string `json:"id"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated,omitempty"`
}

// UpdateClause is ...
type UpdateClause struct {
	Field string
	Value string
}

// Seo is ...
type Seo struct {
	Title       string `json:"title"`
	Keywords    string `json:"keywords"`
	Description string `json:"description"`
}
