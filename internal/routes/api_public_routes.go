package routes

import (
	"github.com/gofiber/fiber/v2"

	handlers "github.com/shurco/litecart/internal/handlers/public"
	"github.com/shurco/litecart/pkg/webutil"
)

// ApiPublicRoutes is ...
func ApiPublicRoutes(c *fiber.App) {
	product := c.Group("/api/products")
	product.Get("/", handlers.Products)
	product.Get("/:product_id", handlers.Product)

	c.Get("/api/cart", func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "Cart", "ok")
	})

	c.Post("/api/checkout-session", func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "Checkout Session", "ok")
	})
}
