package models

// Page is ...
type Page struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	Content string `json:"content,omitempty"`
}
