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
		validation.Field(&v.Password, validation.Required, validation.Length(3, 20)),
	)
}

// JWT iss ...
type JWT struct {
	Secret      string `json:"secret"`
	ExpireHours int    `json:"expire_hours"`
}

// Validate is ...
func (v JWT) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Secret, validation.Length(30, 100)),
		validation.Field(&v.ExpireHours, validation.Length(30, 100)),
	)
}
