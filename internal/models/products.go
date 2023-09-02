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
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Images      []Images   `json:"images,omitempty"`
	Url         string     `json:"url"`
	Amount      int        `json:"amount"`
	Metadata    []Metadata `json:"metadata,omitempty"`
	Attributes  []string   `json:"attributes,omitempty"`
	Active      bool       `json:"active"`
	Created     int64      `json:"created"`
	Updated     int64      `json:"updated,omitempty"`
}

// Validate is ...
func (v Product) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.ID, validation.Length(15, 15)),
		validation.Field(&v.Name, validation.Length(3, 50)),
		validation.Field(&v.Description, validation.NotNil),
		validation.Field(&v.Images),
		validation.Field(&v.Url, validation.Required, validation.Length(1, 20)),
		validation.Field(&v.Amount, validation.Required, validation.Min(0)),
		validation.Field(&v.Metadata),
		validation.Field(&v.Attributes, validation.Each(validation.Length(3, 254))),
		validation.Field(&v.Active),
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

// Images is ...
type Images struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Ext  string `json:"ext"`
}

// Validate is ...
func (v Images) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.ID, validation.Length(15, 15)),
		validation.Field(&v.Name, is.UUIDv4),
		validation.Field(&v.Ext, validation.In("jpeg", "png")),
	)
}
