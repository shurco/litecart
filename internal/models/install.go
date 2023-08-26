package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Install is ...
type Install struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Domain       string `json:"domain"`
	StripeSecret string `json:"stripe_secret"`
}

// Validate is ...
func (v Install) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Email, validation.Required, is.Email),
		validation.Field(&v.Password, validation.Required, validation.Length(6, 50)),
		validation.Field(&v.StripeSecret, validation.Length(100, 120)),
	)
}
