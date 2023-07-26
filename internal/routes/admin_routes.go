package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/middleware"
	"github.com/shurco/litecart/internal/queries"
)

// AdminRoutes is ...
func AdminRoutes(c *fiber.App) {
	route := c.Group("/_")

	route.Get("/install", func(c *fiber.Ctx) error {
		return c.Render("admin/install", nil, "admin/layouts/clear")
	})

	route.Get("/signin", func(c *fiber.Ctx) error {
		return c.Render("admin/signin", nil, "admin/layouts/clear")
	})

	// product section
	route.Get("/products", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		db := queries.DB()
		products, err := db.ListProducts()
		if err != nil {
			return err
		}

		return c.Render("admin/products", fiber.Map{
			"Products": products,
			"Menu":     "products",
		}, "admin/layouts/main")
	})

	// setting section
	route.Get("/settings", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return c.Render("admin/settings", fiber.Map{
			"Menu": "settings",
		}, "admin/layouts/main")
	})
}
