package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/mailer"
	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/internal/webhook"
	"github.com/shurco/litecart/pkg/litepay"
	"github.com/shurco/litecart/pkg/security"
	"github.com/shurco/litecart/pkg/webutil"
)

// Payment is ...
// [get] /cart/payment
func PaymentList(c *fiber.Ctx) error {
	db := queries.DB()
	paymentList, err := db.PaymentList()
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Payment list", paymentList)
}

// Payment is ...
// [post] /cart/payment
func Payment(c *fiber.Ctx) error {
	payment := new(models.Payment)

	if err := c.BodyParser(payment); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	db := queries.DB()

	domain := db.GetDomain()
	products, err := db.ListProducts(false, payment.Products...)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
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
		Currency: db.GetCurrency(),
		Items:    items,
	}

	settingPayment, err := db.SettingBySection(string(payment.Provider))
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	callbackURL := fmt.Sprintf("https://%s/cart/payment/callback", domain)
	successURL := fmt.Sprintf("https://%s/cart/payment/success", domain)
	cancelURL := fmt.Sprintf("https://%s/cart/payment/cancel", domain)
	pay := litepay.New(callbackURL, successURL, cancelURL)

	paymentURL := fmt.Sprintf("https://%s/cart", domain)
	paymentSystem := payment.Provider
	switch paymentSystem {
	case litepay.STRIPE:
		setting := settingPayment.(models.Stripe)
		if !setting.Active {
			return webutil.Response(c, fiber.StatusOK, "Payment url", paymentURL)
		}
		session := pay.Stripe(setting.SecretKey)
		response, err := session.Pay(cart)
		if err != nil {
			return webutil.StatusBadRequest(c, err)
		}
		paymentURL = response.URL

	case litepay.SPECTROCOIN:
		setting := settingPayment.(models.Spectrocoin)
		if !setting.Active {
			return webutil.Response(c, fiber.StatusOK, "Payment url", paymentURL)
		}
		session := pay.Spectrocoin(setting.MerchantID, setting.ProjectID, setting.PrivateKey)
		response, err := session.Pay(cart)
		if err != nil {
			return webutil.StatusBadRequest(c, err)
		}
		paymentURL = response.URL
	}

	var amountTotal int
	for _, s := range cart.Items {
		amountTotal += s.PriceData.UnitAmount * s.Quantity
	}

	db.AddCart(&models.Cart{
		Core: models.Core{
			ID: cart.ID,
		},
		Email:         payment.Email,
		Cart:          payment.Products,
		AmountTotal:   amountTotal,
		Currency:      cart.Currency,
		PaymentStatus: litepay.NEW,
		PaymentSystem: paymentSystem,
	})

	// send email
	if err := mailer.SendPrepaymentLetter(payment.Email, fmt.Sprintf("%.2f %s", float64(amountTotal)/100, cart.Currency), paymentURL); err != nil {
		return webutil.StatusBadRequest(c, err)
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
		return webutil.StatusBadRequest(c, err)
	}

	return webutil.Response(c, fiber.StatusOK, "Payment url", paymentURL)
}

// PaymentCallback is ...
// [get] /cart/payment/callback
func PaymentCallback(c *fiber.Ctx) error {
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
			return webutil.StatusBadRequest(c, err)
		}
		payment.Status = litepay.StatusPayment(string(rune(response.Status)))
		payment.MerchantID = response.MerchantApiID
		payment.Coin = &litepay.Coin{
			AmountTotal: response.ReceiveAmount,
			Currency:    response.ReceiveCurrency,
		}
	}

	// send email
	if payment.Status == litepay.PAY {
		if err := mailer.SendCartLetter(payment.CartID); err != nil {
			return err
		}
	}

	// send hook
	hook := &webhook.Payment{
		Event:     webhook.PAYMENT_CALLBACK,
		TimeStamp: time.Now().Unix(),
		Data: webhook.Data{
			PaymentSystem: payment.PaymentSystem,
			PaymentStatus: payment.Status,
			CartID:        payment.CartID,
		},
	}
	if err := webhook.SendPaymentHook(hook); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	return c.Status(fiber.StatusOK).SendString("*ok*")
}

// PaymentSuccess is ...
// [get] /cart/payment/success
func PaymentSuccess(c *fiber.Ctx) error {
	payment := &litepay.Payment{
		CartID:        c.Query("cart_id"),
		PaymentSystem: litepay.PaymentSystem(c.Query("payment_system")),
	}

	if err := payment.Validate(); err != nil {
		return c.Redirect("/")
	}

	db := queries.DB()
	settingPayment, err := db.SettingBySection(string(payment.PaymentSystem))
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	switch payment.PaymentSystem {
	case litepay.STRIPE:
		sessionStripe := c.Query("session")

		setting := settingPayment.(models.Stripe)
		if !setting.Active {
			return webutil.StatusBadRequest(c, err)
		}

		response, err := litepay.New("", "", "").Stripe(setting.SecretKey).Checkout(payment, sessionStripe)
		if err != nil {
			return webutil.StatusBadRequest(c, err)
		}
		payment.Status = response.Status

	case litepay.SPECTROCOIN:
		fmt.Print(payment)
	}

	err = db.UpdateCart(&models.Cart{
		Core: models.Core{
			ID: payment.CartID,
		},
		PaymentStatus: payment.Status,
		PaymentSystem: payment.PaymentSystem,
	})
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	// send email
	if payment.Status == litepay.PAY {
		if err := mailer.SendCartLetter(payment.CartID); err != nil {
			return webutil.StatusBadRequest(c, err)
		}
	}

	// send hook
	hook := &webhook.Payment{
		Event:     webhook.PAYMENT_SUCCESS,
		TimeStamp: time.Now().Unix(),
		Data: webhook.Data{
			PaymentSystem: payment.PaymentSystem,
			PaymentStatus: payment.Status,
			CartID:        payment.CartID,
		},
	}
	if err := webhook.SendPaymentHook(hook); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	return c.Render("success", nil, "layouts/main")
}

// PaymentCancel is ...
// [get] /cart/payment/cancel
func PaymentCancel(c *fiber.Ctx) error {
	payment := &litepay.Payment{
		CartID:        c.Query("cart_id"),
		PaymentSystem: litepay.PaymentSystem(c.Query("payment_system")),
	}

	db := queries.DB()
	err := db.UpdateCart(&models.Cart{
		Core: models.Core{
			ID: payment.CartID,
		},
		PaymentStatus: litepay.CANCELED,
		PaymentSystem: payment.PaymentSystem,
	})
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	// send hook
	hook := &webhook.Payment{
		Event:     webhook.PAYMENT_CANCEL,
		TimeStamp: time.Now().Unix(),
		Data: webhook.Data{
			PaymentSystem: payment.PaymentSystem,
			PaymentStatus: litepay.CANCELED,
			CartID:        payment.CartID,
		},
	}
	if err := webhook.SendPaymentHook(hook); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	return c.Render("cancel", nil, "layouts/main")
}
