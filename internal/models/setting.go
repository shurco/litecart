package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type Setting struct {
	Main     Main     `json:"main,omitempty"`
	Password Password `json:"password,omitempty"`
	Stripe   Stripe   `json:"stripe,omitempty"`
	Social   Social   `json:"social,omitempty"`
	Mail     Mail     `json:"mail,omitempty"`
}

// Validate is ...
func (v Setting) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Main),
		validation.Field(&v.Stripe),
		validation.Field(&v.Social),
		validation.Field(&v.Mail),
	)
}

type Main struct {
	Domain   string `json:"domain"`
	Email    string `json:"email"`
	Currency string `json:"currency"`
	JWT      JWT    `json:"jwt"`
}

// Validate is ...
func (v Main) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Domain, is.Domain),
		validation.Field(&v.Email, is.Email),
		validation.Field(&v.Currency, is.CurrencyCode),
		validation.Field(&v.JWT),
	)
}

type Password struct {
	Old string `json:"old"`
	New string `json:"new"`
}

// Validate is ...
func (v Password) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Old, validation.Length(6, 30)),
		validation.Field(&v.New, validation.Length(6, 30)),
	)
}

type JWT struct {
	Secret      string `json:"secret"`
	ExpireHours string `json:"expire_hours"`
}

// Validate is ...
func (v JWT) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Secret, validation.Length(30, 100)),
		validation.Field(&v.ExpireHours, validation.Length(30, 100)),
	)
}

type Stripe struct {
	SecretKey        string `json:"secret_key"`
	WebhookSecretKey string `json:"webhook_secret_key"`
}

// Validate is ...
func (v Stripe) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.SecretKey, validation.Length(100, 130)),
		validation.Field(&v.WebhookSecretKey, validation.Length(100, 130)),
	)
}

type Social struct {
	Facebook  string `json:"facebook,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	Twitter   string `json:"twitter,omitempty"`
	Dribbble  string `json:"dribbble,omitempty"`
	Github    string `json:"github,omitempty"`
}

// Validate is ...
func (v Social) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Facebook, validation.Length(3, 20)),
		validation.Field(&v.Instagram, validation.Length(3, 20)),
		validation.Field(&v.Twitter, validation.Length(3, 20)),
		validation.Field(&v.Github, validation.Length(3, 20)),
	)
}

type Mail struct {
	Host       string `json:"smtp_host,omitempty"`
	Port       int    `json:"smtp_port,omitempty"`
	Encryption string `json:"smtp_encryption,omitempty"`
	Username   string `json:"smtp_username,omitempty"`
	Password   string `json:"smtp_password,omitempty"`
}

// Validate is ...
func (v Mail) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.Host, is.Host),
		validation.Field(&v.Port, is.Port),
		//validation.Field(&v.Encryption),
		validation.Field(&v.Username, validation.Length(3, 20)),
		validation.Field(&v.Password, validation.Length(3, 20)),
	)
}
