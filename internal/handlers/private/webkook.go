package handlers

import (
	"github.com/gofiber/fiber/v2"
	"encoding/json"


	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/webhook"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/webutil"
)


// StripeWebHookListen is a handler for stripe webhook events 
// [post] /api/_/stripe-webhook
func StripeWebHookListen(c *fiber.Ctx) error {
	db := queries.DB()
	settings, err := db.Settings(true)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	if (settings.Stripe.WebhookSecretKey == "") {
		return webutil.Response(c, fiber.StatusUnauthorized, "unable to verify", map[string]string{"status": "failed"})			
	}

	event := stripe.Event{}

	signatureHeader := c.Get("Stripe-Signature")
		
	// settings.Stripe.WebhookSecretKey
	event, err = webhook.ConstructEvent(c.BodyRaw(), signatureHeader, settings.Stripe.WebhookSecretKey)

	if err != nil {
		return webutil.Response(c, fiber.StatusUnauthorized, "unable to verify", map[string]string{"status": "failed"})			
	}

	switch event.Type {
		case "checkout.session.completed":
			// TODO: display somwhere i guess
			var chackoutThings stripe.CheckoutSession
			_ = json.Unmarshal(event.Data.Raw, &chackoutThings)
			

		case "payment_intent.succeeded":
			// TODO: display somwhere i guess
			var paymentIntent stripe.PaymentIntent
			_ = json.Unmarshal(event.Data.Raw, &paymentIntent)

			
		case "payment_intent.payment_failed":
			// TODO: display somwhere i guess
			var paymentIntent stripe.PaymentIntent
			_ = json.Unmarshal(event.Data.Raw, &paymentIntent)

			
	}

	return webutil.Response(c, fiber.StatusOK, "webhook", map[string]string{"status": "success"})
}


