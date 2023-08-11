package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/webutil"
)

// SiteRoutes is ...
func SiteRoutes(c *fiber.App) {
	c.Get("/", func(c *fiber.Ctx) error {
		return c.Render("site/index", nil, "site/layouts/main")
	})

	c.Get("/terms", func(c *fiber.Ctx) error {
		return c.Render("site/text", nil, "site/layouts/main")
	})

	c.Get("/privacy", func(c *fiber.Ctx) error {
		return c.Render("site/text", nil, "site/layouts/main")
	})

	c.Get("/cookies", func(c *fiber.Ctx) error {
		return c.Render("site/text", nil, "site/layouts/main")
	})

	c.Get("/cart", func(c *fiber.Ctx) error {
		return c.Render("site/cart", nil, "site/layouts/main")
	})

	c.Get("/products/:product_url", func(c *fiber.Ctx) error {
		productURL := c.Params("product_url")
		db := queries.DB()

		if !db.IsProduct(productURL) {
			return c.Status(fiber.StatusNotFound).Render("site/404", fiber.Map{}, "site/layouts/clear")
		}

		return c.Render("site/product", fiber.Map{
			"ProductUrl": c.Params("product_url"),
		}, "site/layouts/main")
	})

	c.Get("/success", func(c *fiber.Ctx) error {
		return c.Render("site/success", fiber.Map{
			"Title": "Hello, World!",
		}, "site/layouts/clear")
	})

	c.Get("/webhook", func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "Webhook", nil)
	})
}
