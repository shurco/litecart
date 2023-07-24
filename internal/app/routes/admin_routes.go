package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/app/queries"
	"github.com/shurco/litecart/internal/core"
	"github.com/shurco/litecart/internal/core/middleware"
)

// AdminRoutes is ...
func AdminRoutes(c *core.Core, q *queries.Base) {
	route := c.Fiber.Group("/_")

	route.Get("/install", func(c *fiber.Ctx) error {
		return c.Render("admin/install", nil, "admin/layouts/clear")
	})

	route.Get("/signin", func(c *fiber.Ctx) error {
		return c.Render("admin/signin", nil, "admin/layouts/clear")
	})

	route.Get("/", middleware.JWTProtected(c.JWT.Secret), func(c *fiber.Ctx) error {
		return c.Render("admin/index", fiber.Map{
			"Title": "Hello, World!",
		}, "admin/layouts/main")
	})

	product := route.Group("/products", middleware.JWTProtected(c.JWT.Secret))
	product.Get("/", middleware.JWTProtected(c.JWT.Secret), func(c *fiber.Ctx) error {
		return c.Render("admin/products", fiber.Map{
			"Menu": "products",
		}, "admin/layouts/main")
	})
	product.Get("/add", middleware.JWTProtected(c.JWT.Secret), func(c *fiber.Ctx) error {
		return c.Render("admin/products_add", fiber.Map{
			"Menu": "products",
		}, "admin/layouts/main")
	})
	product.Get("/update", middleware.JWTProtected(c.JWT.Secret), func(c *fiber.Ctx) error {
		return c.Render("admin/products_update", fiber.Map{
			"Menu": "products",
		}, "admin/layouts/main")
	})

	route.Get("/invoices", middleware.JWTProtected(c.JWT.Secret), func(c *fiber.Ctx) error {
		return c.Render("admin/invoices", fiber.Map{
			"Menu": "invoices",
		}, "admin/layouts/main")
	})

	route.Get("/settings", middleware.JWTProtected(c.JWT.Secret), func(c *fiber.Ctx) error {
		return c.Render("admin/settings", fiber.Map{
			"Menu": "settings",
		}, "admin/layouts/main")
	})

}
