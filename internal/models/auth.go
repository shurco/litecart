package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// SignIn is ...
type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate is ...
func (v SignIn) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Email, validation.Required, is.Email),
		validation.Field(&v.Password, validation.Required, validation.Length(6, 50)),
	)
}

// Renew is ...
type Renew struct {
	RefreshToken string `json:"refresh_token"`
}
