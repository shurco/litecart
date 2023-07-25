package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shurco/litecart/pkg/webutil"
)

// SiteRoutes is ...
func SiteRoutes(c *fiber.App) {
	route := c.Group("/")

	route.Get("/", func(c *fiber.Ctx) error {
		return c.Render("site/index", fiber.Map{
			"Title": "Hello, World!",
		}, "site/layouts/main")
	})

	route.Get("/success", func(c *fiber.Ctx) error {
		return c.Render("site/success", fiber.Map{
			"Title": "Hello, World!",
		}, "site/layouts/clear")
	})

	route.Get("/webhook", func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "Webhook", "ok")
	})
}
