package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shurco/litecart/internal/mailer"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/webutil"
)

// Carts is ...
// [get] /api/_/carts
func Carts(c *fiber.Ctx) error {
	db := queries.DB()

	products, err := db.Carts()
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Carts", products)
}

// CartSendMail
// [post] /api/_/carts/:cart_id/mail
func CartSendMail(c *fiber.Ctx) error {
	cartID := c.Params("cart_id")

	if err := mailer.SendCartLetter(cartID); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Mail sended", nil)
}
