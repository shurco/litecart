package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/pkg/webutil"
)

// ApiPublicRoutes is ...
func ApiPublicRoutes(c *fiber.App) {
	c.Get("/api/cart", func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "Cart", "ok")
	})

	c.Post("/api/checkout-session", func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "Checkout Session", "ok")
	})
}
