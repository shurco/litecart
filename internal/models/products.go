package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Products is ...
type Products struct {
	Total    int       `json:"total"`
	Currency string    `json:"currency"`
	Products []Product `json:"products"`
}

// Product is ...
type Product struct {
	Core
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	Images      []File     `json:"images,omitempty"`
	Slug        string     `json:"slug"`
	Amount      int        `json:"amount"`
	Metadata    []Metadata `json:"metadata,omitempty"`
	Attributes  []string   `json:"attributes,omitempty"`
	Digital     Digital    `json:"digital,omitempty"`
	Active      bool       `json:"active"`
	Seo         *Seo       `json:"seo,omitempty"`
}

// Validate is ...
func (v Product) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.ID, validation.Length(15, 15)),
		validation.Field(&v.Name, validation.Length(3, 50)),
		validation.Field(&v.Description, validation.NotNil),
		validation.Field(&v.Images),
		validation.Field(&v.Slug, validation.Required, validation.Length(3, 20)),
		validation.Field(&v.Amount, validation.Required, validation.Min(0)),
		validation.Field(&v.Metadata),
		validation.Field(&v.Attributes, validation.Each(validation.Length(3, 254))),
		validation.Field(&v.Digital),
		validation.Field(&v.Seo),
	)
}

// Metadata is ...
type Metadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Validate is ...
func (v Metadata) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Key, validation.Required, validation.Length(1, 20)),
		validation.Field(&v.Value, validation.Required, validation.Min(0)),
	)
}

// Digital is ...
type Digital struct {
	Type   string `json:"type"`
	Filled bool   `json:"filled,omitempty"`
	Files  []File `json:"files,omitempty"`
	Data   []Data `json:"data,omitempty"`
}

// Validate is ...
func (v Digital) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Type, validation.In("file", "data", "api")),
		validation.Field(&v.Files),
		validation.Field(&v.Data, validation.Each(validation.Length(1, 254))),
	)
}

// File is ...
type File struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Ext      string `json:"ext"`
	OrigName string `json:"orig_name,omitempty"`
}

// Validate is ...
func (v File) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.ID, validation.Length(15, 15)),
		validation.Field(&v.Name, is.UUIDv4),
		// validation.Field(&v.Ext, validation.In("jpeg", "png")),
	)
}

// Data is ...
type Data struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	CartID  string `json:"cart_id"`
}

// Validate is ...
func (v Data) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.ID, validation.Length(15, 15)),
		validation.Field(&v.Content, validation.Length(1, 254)),
		// validation.Field(&v.Ext, validation.In("jpeg", "png")),
	)
}
