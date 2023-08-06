package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shurco/litecart/pkg/webutil"
)

// SiteRoutes is ...
func SiteRoutes(c *fiber.App) {
	c.Get("/", func(c *fiber.Ctx) error {
		return c.Render("site/index", fiber.Map{
			"Title": "Hello, World!",
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
