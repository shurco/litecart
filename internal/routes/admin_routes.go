package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/middleware"
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

	route.Get("/", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return c.Render("admin/index", fiber.Map{
			"Title": "Hello, World!",
		}, "admin/layouts/main")
	})

	// product section
	product := route.Group("/products", middleware.JWTProtected())
	product.Get("/", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return c.Render("admin/products", fiber.Map{
			"Menu": "products",
		}, "admin/layouts/main")
	})
	product.Get("/add", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return c.Render("admin/products_add", fiber.Map{
			"Menu": "products",
		}, "admin/layouts/main")
	})
	product.Get("/update", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return c.Render("admin/products_update", fiber.Map{
			"Menu": "products",
		}, "admin/layouts/main")
	})

	// invoice section
	route.Get("/invoices", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return c.Render("admin/invoices", fiber.Map{
			"Menu": "invoices",
		}, "admin/layouts/main")
	})

	// setting section
	route.Get("/settings", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return c.Render("admin/settings", fiber.Map{
			"Menu": "settings",
		}, "admin/layouts/main")
	})

}
