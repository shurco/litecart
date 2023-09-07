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

	c.Get("/terms", func(c *fiber.Ctx) error {
		return c.Render("pages", fiber.Map{
			"PageUrl": "terms",
		}, "layouts/main")
	})

	c.Get("/privacy", func(c *fiber.Ctx) error {
		return c.Render("pages", fiber.Map{
			"PageUrl": "privacy",
		}, "layouts/main")
	})

	c.Get("/cookies", func(c *fiber.Ctx) error {
		return c.Render("pages", fiber.Map{
			"PageUrl": "cookies",
		}, "layouts/main")
	})

	// catalog section
	c.Get("/products/:product_slug", func(c *fiber.Ctx) error {
		productSlug := c.Params("product_slug")
		db := queries.DB()

		if !db.IsProduct(productSlug) {
			return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{}, "layouts/clear")
		}

		return c.Render("product", fiber.Map{
			"ProductSlug": productSlug,
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

}
