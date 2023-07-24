package models

type Install struct {
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=6"`
	Domain       string `json:"domain" validate:"required"`
	StripeSecret string `json:"stripe_secret" validate:"required,min=100"`
}
