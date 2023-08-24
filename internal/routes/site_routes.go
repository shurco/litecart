package routes

import (
	"github.com/gofiber/fiber/v2"

	handlers "github.com/shurco/litecart/internal/handlers/public"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/webutil"
)

// SiteRoutes is ...
func SiteRoutes(c *fiber.App) {
	c.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil, "layouts/main")
	})

	c.Get("/terms", func(c *fiber.Ctx) error {
		return c.Render("text", nil, "layouts/main")
	})

	c.Get("/privacy", func(c *fiber.Ctx) error {
		return c.Render("text", nil, "layouts/main")
	})

	c.Get("/cookies", func(c *fiber.Ctx) error {
		return c.Render("text", nil, "layouts/main")
	})

	// catalog section
	c.Get("/products/:product_url", func(c *fiber.Ctx) error {
		productURL := c.Params("product_url")
		db := queries.DB()

		if !db.IsProduct(productURL) {
			return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{}, "layouts/clear")
		}

		return c.Render("product", fiber.Map{
			"ProductUrl": c.Params("product_url"),
		}, "layouts/main")
	})

	// cart section
	cart := c.Group("/cart")

	cart.Get("/", func(c *fiber.Ctx) error {
		return c.Render("cart", nil, "layouts/main")
	})

	cart.Post("/checkout", handlers.Checkout)
	cart.Get("/success/:cart_id<len(15)>/:session_id", handlers.CheckoutSuccess)
	cart.Get("/cancel/:cart_id<len(15)>", handlers.CheckoutCancel)

	c.Get("/webhook", func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "Webhook", nil)
	})
}
