package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/mailer"
	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/internal/webhook"
	"github.com/shurco/litecart/pkg/errors"
	"github.com/shurco/litecart/pkg/litepay"
	"github.com/shurco/litecart/pkg/logging"
	"github.com/shurco/litecart/pkg/security"
	"github.com/shurco/litecart/pkg/webutil"
)

// sendPaymentWebhook sends a payment webhook notification.
// If blockOnError is true, returns error on webhook failure (for API endpoints).
// If blockOnError is false, logs error but doesn't block (for user-facing pages).
func sendPaymentWebhook(event webhook.Event, paymentSystem litepay.PaymentSystem, paymentStatus litepay.Status, cartID string, log *logging.Log, blockOnError bool) error {
	hook := &webhook.Payment{
		Event:     event,
		TimeStamp: time.Now().Unix(),
		Data: webhook.Data{
			PaymentSystem: paymentSystem,
			PaymentStatus: paymentStatus,
			CartID:        cartID,
		},
	}

	if err := webhook.SendPaymentHook(hook); err != nil {
		log.ErrorStack(err)
		if blockOnError {
			return err
		}
	}
	return nil
}

// PaymentList returns a list of available payment systems.
// [get] /api/cart/payment
func PaymentList(c *fiber.Ctx) error {
	db := queries.DB()
	log := logging.New()
	paymentList, err := db.PaymentList(c.Context())
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Payment list", paymentList)
}

// GetCart returns cart information by cart_id.
// [get] /api/cart/:cart_id
func GetCart(c *fiber.Ctx) error {
	db := queries.DB()
	log := logging.New()
	cartID := c.Params("cart_id")

	if cartID == "" {
		return webutil.StatusBadRequest(c, "cart_id is required")
	}

	cart, err := db.Cart(c.Context(), cartID)
	if err != nil {
		log.ErrorStack(err)
		if err == errors.ErrProductNotFound {
			return webutil.StatusNotFound(c)
		}
		return webutil.StatusInternalServerError(c)
	}

	// Load full product information for cart items
	if len(cart.Cart) > 0 {
		productIDs := make([]string, len(cart.Cart))
		for i, item := range cart.Cart {
			productIDs[i] = item.ProductID
		}

		products, err := db.ListProducts(c.Context(), false, cart.Cart...)
		if err != nil {
			log.ErrorStack(err)
			return webutil.StatusInternalServerError(c)
		}

		// Create a map for quick lookup
		productMap := make(map[string]*models.Product)
		for i := range products.Products {
			productMap[products.Products[i].ID] = &products.Products[i]
		}

		// Build cart items with full product information
		cartItems := make([]map[string]interface{}, 0, len(cart.Cart))
		for _, cartItem := range cart.Cart {
			if product, ok := productMap[cartItem.ProductID]; ok {
				var image interface{}
				if len(product.Images) > 0 {
					image = product.Images[0]
				}
				cartItems = append(cartItems, map[string]interface{}{
					"id":       product.ID,
					"name":     product.Name,
					"slug":     product.Slug,
					"amount":   product.Amount,
					"quantity": cartItem.Quantity,
					"image":    image,
				})
			}
		}

		// Return cart with full product information
		return webutil.Response(c, fiber.StatusOK, "Cart", map[string]interface{}{
			"id":             cart.ID,
			"email":          cart.Email,
			"amount_total":   cart.AmountTotal,
			"currency":       cart.Currency,
			"payment_status": cart.PaymentStatus,
			"payment_system": cart.PaymentSystem,
			"items":          cartItems,
		})
	}

	return webutil.Response(c, fiber.StatusOK, "Cart", cart)
}

// Payment initiates a payment process for a cart.
// [post] /cart/payment
func Payment(c *fiber.Ctx) error {
	db := queries.DB()
	log := logging.New()
	payment := new(models.CartPayment)

	if err := c.BodyParser(payment); err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}

	setting, err := db.GetSettingByKey(c.Context(), "domain", "currency")
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}
	domain := setting["domain"].Value.(string)
	currency := setting["currency"].Value.(string)

	products, err := db.ListProducts(c.Context(), false, payment.Products...)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	items := make([]litepay.Item, len(products.Products))
	for i, product := range products.Products {
		images := []string{}
		for _, image := range product.Images {
			path := fmt.Sprintf("https://%s/uploads/%s_md.%s", domain, image.Name, image.Ext)
			images = append(images, path)
		}

		quantity := 1
		for _, cartProduct := range payment.Products {
			if cartProduct.ProductID == product.ID {
				quantity = cartProduct.Quantity
			}
		}

		items[i] = litepay.Item{
			PriceData: litepay.Price{
				UnitAmount: product.Amount,
				Product: litepay.Product{
					Name:   product.Name,
					Images: images,
				},
			},
			Quantity: quantity,
		}

		if product.Description != "" {
			items[i].PriceData.Product.Description = product.Description
		}
	}

	cart := litepay.Cart{
		ID:       security.RandomString(),
		Currency: currency,
		Items:    items,
	}

	callbackURL := fmt.Sprintf("https://%s/cart/payment/callback", domain)
	successURL := fmt.Sprintf("https://%s/cart/payment/success", domain)
	cancelURL := fmt.Sprintf("https://%s/cart/payment/cancel", domain)
	pay := litepay.New(callbackURL, successURL, cancelURL)

	paymentURL := fmt.Sprintf("https://%s/cart", domain)
	paymentSystem := payment.Provider
	switch paymentSystem {
	case litepay.STRIPE:
		setting, err := queries.GetSettingByGroup[models.Stripe](c.Context(), db)
		if err != nil {
			log.ErrorStack(err)
			return webutil.StatusInternalServerError(c)
		}

		if !setting.Active {
			return webutil.Response(c, fiber.StatusOK, "Payment url", paymentURL)
		}
		session := pay.Stripe(setting.SecretKey)
		response, err := session.Pay(cart)
		if err != nil {
			log.ErrorStack(err)
			return webutil.StatusInternalServerError(c)
		}
		paymentURL = response.URL

	case litepay.PAYPAL:
		setting, err := queries.GetSettingByGroup[models.Paypal](c.Context(), db)
		if err != nil {
			log.ErrorStack(err)
			return webutil.StatusInternalServerError(c)
		}

		if !setting.Active {
			return webutil.Response(c, fiber.StatusOK, "Payment url", paymentURL)
		}
		session := pay.Paypal(setting.ClientID, setting.SecretKey)
		response, err := session.Pay(cart)
		if err != nil {
			log.ErrorStack(err)
			return webutil.StatusInternalServerError(c)
		}
		paymentURL = response.URL

	case litepay.SPECTROCOIN:
		setting, err := queries.GetSettingByGroup[models.Spectrocoin](c.Context(), db)
		if err != nil {
			log.ErrorStack(err)
			return webutil.StatusInternalServerError(c)
		}

		if !setting.Active {
			return webutil.Response(c, fiber.StatusOK, "Payment url", paymentURL)
		}
		session := pay.Spectrocoin(setting.MerchantID, setting.ProjectID, setting.PrivateKey)
		response, err := session.Pay(cart)
		if err != nil {
			log.ErrorStack(err)
			return webutil.StatusInternalServerError(c)
		}
		paymentURL = response.URL
	}

	var amountTotal int
	for _, s := range cart.Items {
		amountTotal += s.PriceData.UnitAmount * s.Quantity
	}

	if err := db.AddCart(c.Context(), &models.Cart{
		Core: models.Core{
			ID: cart.ID,
		},
		Email:         payment.Email,
		Cart:          payment.Products,
		AmountTotal:   amountTotal,
		Currency:      cart.Currency,
		PaymentStatus: litepay.NEW,
		PaymentSystem: paymentSystem,
	}); err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	// send email
	if err := mailer.SendPrepaymentLetter(payment.Email, fmt.Sprintf("%.2f %s", float64(amountTotal)/100, cart.Currency), paymentURL); err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	// send hook
	hook := &webhook.Payment{
		Event:     webhook.PAYMENT_INITIATION,
		TimeStamp: time.Now().Unix(),
		Data: webhook.Data{
			PaymentSystem: paymentSystem,
			PaymentStatus: litepay.NEW,
			CartID:        cart.ID,
			TotalAmount:   amountTotal,
			Currency:      cart.Currency,
			CartItems:     items,
		},
	}
	if err := webhook.SendPaymentHook(hook); err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	return webutil.Response(c, fiber.StatusOK, "Payment url", map[string]string{"url": paymentURL})
}

// PaymentCallback handles payment callback from payment providers.
// [post] /cart/payment/callback
func PaymentCallback(c *fiber.Ctx) error {
	log := logging.New()
	payment := &litepay.Payment{
		CartID:        c.Query("cart_id"),
		PaymentSystem: litepay.PaymentSystem(c.Query("payment_system")),
	}

	switch payment.PaymentSystem {
	// case litepay.STRIPE:
	//	return webutil.Response(c, fiber.StatusOK, "Callback", payment)
	case litepay.SPECTROCOIN:
		response := new(litepay.CallbackSpectrocoin)
		if err := c.BodyParser(response); err != nil {
			log.ErrorStack(err)
			return webutil.StatusBadRequest(c, err.Error())
		}
		payment.Status = litepay.StatusPayment(litepay.SPECTROCOIN, string(rune(response.Status)))
		payment.MerchantID = response.MerchantApiID
		payment.Coin = &litepay.Coin{
			AmountTotal: response.ReceiveAmount,
			Currency:    response.ReceiveCurrency,
		}
	}

	db := queries.DB()
	err := db.UpdateCart(c.Context(), &models.Cart{
		Core: models.Core{
			ID: payment.CartID,
		},
		PaymentID:     payment.MerchantID,
		PaymentStatus: payment.Status,
		PaymentSystem: payment.PaymentSystem,
	})
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	// send email
	if payment.Status == litepay.PAID {
		if err := mailer.SendCartLetter(payment.CartID); err != nil {
			log.ErrorStack(err)
			return webutil.StatusInternalServerError(c)
		}
	}

	// send hook
	if err := sendPaymentWebhook(webhook.PAYMENT_CALLBACK, payment.PaymentSystem, payment.Status, payment.CartID, log, true); err != nil {
		return webutil.StatusInternalServerError(c)
	}

	return c.Status(fiber.StatusOK).SendString("*ok*")
}

// PaymentSuccess handles successful payment redirects.
// [get] /cart/payment/success
func PaymentSuccess(c *fiber.Ctx) error {
	// Only process GET requests
	if c.Method() != fiber.MethodGet {
		return c.Next()
	}

	log := logging.New()
	if c.Query("cart_id") == "" {
		return webutil.StatusBadRequest(c, nil)
	}

	payment := &litepay.Payment{
		CartID:        c.Query("cart_id"),
		PaymentSystem: litepay.PaymentSystem(c.Query("payment_system")),
	}

	if err := payment.Validate(); err != nil {
		return c.Redirect("/")
	}

	db := queries.DB()
	cartInfo, err := db.Cart(c.Context(), c.Query("cart_id"))
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	// If already paid, pass control to SPA handler
	if cartInfo.PaymentStatus == "paid" {
		return c.Next()
	}

	switch payment.PaymentSystem {
	case litepay.STRIPE:
		sessionStripe := c.Query("session")
		setting, err := queries.GetSettingByGroup[models.Stripe](c.Context(), db)
		if err != nil {
			log.ErrorStack(err)
			return webutil.StatusInternalServerError(c)
		}

		if !setting.Active {
			return webutil.StatusNotFound(c)
		}
		response, err := litepay.New("", "", "").Stripe(setting.SecretKey).Checkout(payment, sessionStripe)
		if err != nil {
			log.ErrorStack(err)
			return webutil.StatusInternalServerError(c)
		}
		payment.MerchantID = response.MerchantID
		payment.Status = response.Status

	case litepay.PAYPAL:
		tokenPaypal := c.Query("token")
		setting, err := queries.GetSettingByGroup[models.Paypal](c.Context(), db)
		if err != nil {
			log.ErrorStack(err)
			return webutil.StatusInternalServerError(c)
		}

		if !setting.Active {
			return webutil.StatusNotFound(c)
		}
		response, err := litepay.New("", "", "").Paypal(setting.ClientID, setting.SecretKey).Checkout(payment, tokenPaypal)
		if err != nil {
			log.ErrorStack(err)
			return webutil.StatusInternalServerError(c)
		}
		payment.MerchantID = response.MerchantID
		payment.Status = response.Status

	case litepay.SPECTROCOIN:
		// Spectrocoin payment processing handled in callback
	}

	err = db.UpdateCart(c.Context(), &models.Cart{
		Core: models.Core{
			ID: payment.CartID,
		},
		PaymentID:     payment.MerchantID,
		PaymentStatus: payment.Status,
		PaymentSystem: payment.PaymentSystem,
	})
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	// send email
	if payment.Status == litepay.PAID {
		if err := mailer.SendCartLetter(payment.CartID); err != nil {
			log.ErrorStack(err)
			return webutil.StatusInternalServerError(c)
		}
	}

	// send hook (не блокируем процесс при ошибке webhook)
	sendPaymentWebhook(webhook.PAYMENT_SUCCESS, payment.PaymentSystem, payment.Status, payment.CartID, log, false)

	// After processing payment, pass control to SPA handler
	// The SPA will display the success page with cart information
	return c.Next()
}

// PaymentCancel handles canceled payment redirects.
// [get] /cart/payment/cancel
func PaymentCancel(c *fiber.Ctx) error {
	// Only process GET requests
	if c.Method() != fiber.MethodGet {
		return c.Next()
	}

	log := logging.New()
	payment := &litepay.Payment{
		CartID:        c.Query("cart_id"),
		PaymentSystem: litepay.PaymentSystem(c.Query("payment_system")),
	}

	db := queries.DB()
	err := db.UpdateCart(c.Context(), &models.Cart{
		Core: models.Core{
			ID: payment.CartID,
		},
		PaymentStatus: litepay.CANCELED,
		PaymentSystem: payment.PaymentSystem,
	})
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	// send hook (не блокируем процесс при ошибке webhook)
	sendPaymentWebhook(webhook.PAYMENT_CANCEL, payment.PaymentSystem, litepay.CANCELED, payment.CartID, log, false)

	// Redirect to SPA cancel page with query parameters
	redirectURL := "/cart/payment/cancel"
	if payment.CartID != "" {
		redirectURL += "?cart_id=" + payment.CartID
		if string(payment.PaymentSystem) != "" {
			redirectURL += "&payment_system=" + string(payment.PaymentSystem)
		}
	}
	return c.Redirect(redirectURL)
}
