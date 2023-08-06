package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/middleware"
)

// AdminRoutes is ...
func AdminRoutes(c *fiber.App) {
	admin := c.Group("/_")

	admin.Get("/install", func(c *fiber.Ctx) error {
		return c.Render("admin/install", nil, "admin/layouts/clear")
	})

	admin.Get("/signin", func(c *fiber.Ctx) error {
		return c.Render("admin/signin", nil, "admin/layouts/clear")
	})

	admin.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/_/products")
	})

	// product section
	admin.Get("/products", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return c.Render("admin/products", fiber.Map{
			"Menu": "products",
		}, "admin/layouts/main")
	})

	// setting section
	admin.Get("/settings", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return c.Render("admin/settings", fiber.Map{
			"Menu": "settings",
		}, "admin/layouts/main")
	})
}
