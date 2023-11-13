package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/security"
	"github.com/shurco/litecart/pkg/webhook"
	"github.com/shurco/litecart/pkg/webutil"
)

// Checkout is ...
// [post] /cart/checkout
func Checkout(c *fiber.Ctx) error {
	items := &[]models.CartProduct{}
	if err := c.BodyParser(items); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	idList := []string{}
	for _, item := range *items {
		idList = append(idList, item.ProductID)
	}

	db := queries.DB()
	domain := db.GetDomain()
	currency := db.GetCurrency()
	products, err := db.ListProducts(false, idList...)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	settingStripe, err := db.SettingStripe()
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	lineItems := []*stripe.CheckoutSessionLineItemParams{}
	for _, item := range products.Products {

		images := []string{}
		for _, image := range item.Images {
			path := fmt.Sprintf("https://%s/uploads/%s_md.%s", domain, image.Name, image.Ext)
			images = append(images, path)
		}

		itemCart := &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				UnitAmount: stripe.Int64(int64(item.Amount)),
				Currency:   stripe.String(currency),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name:   stripe.String(item.Name),
					Images: stripe.StringSlice(images),
				},
			},
			Quantity: stripe.Int64(1),
		}

		if item.Description != "" {
			itemCart.PriceData.ProductData.Description = stripe.String(item.Description)
		}

		lineItems = append(lineItems, itemCart)
	}

	cartID := security.RandomString()
	stripe.Key = settingStripe.Stripe.SecretKey
	params := &stripe.CheckoutSessionParams{
		LineItems: lineItems,
		//AutomaticTax: &stripe.CheckoutSessionAutomaticTaxParams{
		//	Enabled: stripe.Bool(true),
		//},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(settingStripe.Main.Domain + "/cart/success/" + cartID + "/{CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String(settingStripe.Main.Domain + "/cart/cancel/" + cartID),
	}

	stripeSession, err := session.New(params)
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	db.AddCart(&models.Cart{
		Core: models.Core{
			ID: cartID,
		},
		Cart:          *items,
		AmountTotal:   stripeSession.AmountTotal,
		Currency:      string(stripeSession.Currency),
		PaymentStatus: string(stripeSession.PaymentStatus),
	})

	if err = settingStripe.Payment.Validate(); err != nil {
		log.Printf("update payment webhook url", err)
	} else {
		resData := map[string]any{
			"event":     "payment_initiation",
			"timestamp": stripeSession.Created,
			"data": map[string]any{
				"payment_id":     stripeSession.ID,
				"total_amount":   stripeSession.AmountTotal,
				"currency":       stripeSession.Currency,
				"cart_items":     items,
			},
		}

		jsonData, err := json.Marshal(resData)
		if err != nil {
			log.Println("Error:", err)
		}

		go func() {
			res, err := webhook.SendHook(settingStripe.Payment.WebhookUrl, jsonData)

			if err != nil {
				log.Println(err)
			}
			if res.StatusCode != 200 {
				log.Print("An issue has been identified with the payment webhook URL. Please verify that it responds with a status code of 200 OK.")
			}
		}()

	}

	return webutil.Response(c, fiber.StatusOK, "Checkout url", stripeSession.URL)
}

// CheckoutSuccess is ...
// [get] /cart/success/:cart_id/:session_id
func CheckoutSuccess(c *fiber.Ctx) error {
	db := queries.DB()

	settingStripe, err := db.SettingStripe()

	cartID := c.Params("cart_id")
	sessionID := c.Params("session_id")

	sessionStripe, err := session.Get(sessionID, nil)
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	err = db.UpdateCart(&models.Cart{
		Core: models.Core{
			ID: cartID,
		},
		Email:         sessionStripe.CustomerDetails.Email,
		Name:          sessionStripe.CustomerDetails.Name,
		PaymentID:     sessionStripe.PaymentIntent.ID,
		PaymentStatus: string(sessionStripe.PaymentStatus),
	})
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	if err = settingStripe.Payment.Validate(); err != nil {
		log.Printf("update payment webhook url", err)
	} else {
		resData := map[string]any{
			"event":     "payment_success",
			"timestamp": sessionStripe.Created,
			"data": map[string]any{
				"payment_system": "stripe",
				"cart_id":        cartID,
				"payment_id":     sessionStripe.PaymentIntent.ID,
				"total_amount":   sessionStripe.PaymentIntent.Amount,
				"currency":       sessionStripe.PaymentIntent.Currency,
				"user_email":     sessionStripe.Customer.Email,
			},
		}

		jsonData, err := json.Marshal(resData)
		if err != nil {
			log.Println("Error:", err)
		}

		go func() {
		res, err := webhook.SendHook(settingStripe.Payment.WebhookUrl, jsonData)

		if err != nil {
			log.Println(err)
		}

		if res.StatusCode != 200 {
			log.Print("An issue has been identified with the payment webhook URL. Please verify that it responds with a status code of 200 OK.")
		}
			}()
	}

	return c.Render("success", nil, "layouts/main")
}

// CheckoutCancel is ...
// [get] /cart/cancel/:cart_id
func CheckoutCancel(c *fiber.Ctx) error {
	cartID := c.Params("cart_id")
	db := queries.DB()
	settingStripe, err := db.SettingStripe()
	cartStripe, err := db.Cart(cartID)

	err = db.UpdateCart(&models.Cart{
		Core: models.Core{
			ID: cartID,
		},
		PaymentStatus: "cancel",
	})

	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	currentTime := time.Now().UTC()

	if err = settingStripe.Payment.Validate(); err != nil {
		log.Print("update payment webhook url", err)
	} else {
		resData := map[string]any{
			"event":     "payment_error",
			"timestamp": currentTime.Unix(),
			"data": map[string]any{
				"payment_system": "stripe",
				"cart_id":        cartID,
				"payment_id":     cartStripe.PaymentID,
				"total_amount":   cartStripe.AmountTotal,
				"currency":       cartStripe.Currency,
				"user_email":     cartStripe.Email,
			},
		}

		jsonData, err := json.Marshal(resData)
		if err != nil {
			log.Println("Error:", err)
		}

		go func() {
		res, err := webhook.SendHook(settingStripe.Payment.WebhookUrl, jsonData)

		if err != nil {
			log.Println(err)
		}

		if res.StatusCode != 200 {
			log.Print("An issue has been identified with the payment webhook URL. Please verify that it responds with a status code of 200 OK.")
		}
		}()
	}

	return c.Render("cancel", nil, "layouts/main")
}
