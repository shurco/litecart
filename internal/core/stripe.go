package core

import (
	"github.com/stripe/stripe-go/v74/client"
)

type Stripe struct {
	SecretKey  string
	WebhookKey string
	Client     *client.API
}

func InitStripeClient(secretKey string) *client.API {
	client := &client.API{}
	client.Init(secretKey, nil)
	return client
}
