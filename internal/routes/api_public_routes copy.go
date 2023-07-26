package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/pkg/webutil"
)

// ApiPublicRoutes is ...
// route have path '/api'
func ApiPublicRoutes(route fiber.Router) {
	route.Get("/cart", func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "Cart", "ok")
	})

	route.Post("/checkout-session", func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "Checkout Session", "ok")
	})
}
