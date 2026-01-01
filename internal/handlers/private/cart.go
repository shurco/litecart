package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shurco/litecart/internal/mailer"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/errors"
	"github.com/shurco/litecart/pkg/logging"
	"github.com/shurco/litecart/pkg/webutil"
)

// Carts returns a list of all carts.
// [get] /api/_/carts
func Carts(c *fiber.Ctx) error {
	db := queries.DB()
	log := logging.New()

	products, err := db.Carts(c.Context())
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	return webutil.Response(c, fiber.StatusOK, "Carts", products)
}

// Cart returns detailed cart information by cart_id.
// [get] /api/_/carts/:cart_id
func Cart(c *fiber.Ctx) error {
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
	var cartItems []map[string]interface{}
	if len(cart.Cart) > 0 {
		products, err := db.ListProducts(c.Context(), false, cart.Cart...)
		if err != nil {
			log.ErrorStack(err)
			return webutil.StatusInternalServerError(c)
		}
		cartItems = queries.BuildCartItems(cart, products)
	}

	return webutil.Response(c, fiber.StatusOK, "Cart", map[string]interface{}{
		"id":             cart.ID,
		"email":          cart.Email,
		"amount_total":   cart.AmountTotal,
		"currency":       cart.Currency,
		"payment_status": cart.PaymentStatus,
		"payment_system": cart.PaymentSystem,
		"payment_id":     cart.PaymentID,
		"created":        cart.Created,
		"updated":        cart.Updated,
		"items":          cartItems,
	})
}

// CartSendMail sends an email notification for a cart.
// [post] /api/_/carts/:cart_id/mail
func CartSendMail(c *fiber.Ctx) error {
	cartID := c.Params("cart_id")
	log := logging.New()

	if err := mailer.SendCartLetter(cartID); err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	return webutil.Response(c, fiber.StatusOK, "Mail sended", nil)
}
