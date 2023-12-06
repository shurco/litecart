package routes

import (
	"github.com/gofiber/fiber/v2"

	handlers "github.com/shurco/litecart/internal/handlers/public"
	"github.com/shurco/litecart/internal/queries"
)

// SiteRoutes is ...
func SiteRoutes(c *fiber.App) {
	c.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil, "layouts/main")
	})

	// catalog section
	c.Get("/products/:product_slug", func(c *fiber.Ctx) error {
		productSlug := c.Params("product_slug")
		db := queries.DB()

		if !db.IsProduct(c.Context(), productSlug) {
			return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{}, "layouts/clear")
		}

		return c.Render("product", fiber.Map{
			"ProductSlug": productSlug,
		}, "layouts/main")
	})

	// cart section
	c.Get("/cart", func(c *fiber.Ctx) error {
		return c.Render("cart", nil, "layouts/main")
	})

	payment := c.Group("/cart/payment")
	payment.Post("/", handlers.Payment)
	payment.Post("/callback", handlers.PaymentCallback)
	payment.Get("/success", handlers.PaymentSuccess)
	payment.Get("/cancel", handlers.PaymentCancel)
}
