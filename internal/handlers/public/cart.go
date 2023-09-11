package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/security"
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

		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
				UnitAmount: stripe.Int64(int64(item.Amount)),
				Currency:   stripe.String(currency),
				ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
					Name:        stripe.String(item.Name),
					Description: stripe.String(item.Description),
					Images:      stripe.StringSlice(images),
				},
			},
			Quantity: stripe.Int64(1),
		})
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
		ID:            cartID,
		Cart:          *items,
		AmountTotal:   stripeSession.AmountTotal,
		Currency:      string(stripeSession.Currency),
		PaymentStatus: string(stripeSession.PaymentStatus),
	})

	return webutil.Response(c, fiber.StatusOK, "Checkout url", stripeSession.URL)
}

// CheckoutSuccess is ...
// [get] /cart/success/:cart_id/:session_id
func CheckoutSuccess(c *fiber.Ctx) error {
	db := queries.DB()

	cartID := c.Params("cart_id")
	sessionID := c.Params("session_id")

	sessionStripe, err := session.Get(sessionID, nil)
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	err = db.UpdateCart(&models.Cart{
		ID:            cartID,
		Email:         sessionStripe.CustomerDetails.Email,
		Name:          sessionStripe.CustomerDetails.Name,
		PaymentID:     sessionStripe.PaymentIntent.ID,
		PaymentStatus: string(sessionStripe.PaymentStatus),
	})
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	return c.Render("success", nil, "layouts/main")
}

// CheckoutCancel is ...
// [get] /cart/cancel/:cart_id
func CheckoutCancel(c *fiber.Ctx) error {
	cartID := c.Params("cart_id")
	db := queries.DB()
	err := db.UpdateCart(&models.Cart{
		ID:            cartID,
		PaymentStatus: "cancel",
	})
	if err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	return c.Render("cancel", nil, "layouts/main")
}
