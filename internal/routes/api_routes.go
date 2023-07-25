package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/handlers"
	"github.com/shurco/litecart/internal/middleware"
	"github.com/shurco/litecart/pkg/webutil"
)

// ApiRoutes is ...
func ApiRoutes(c *fiber.App) {
	route := c.Group("/api")
	route.Post("/install", handlers.Install)

	sign := c.Group("/api/sign")
	sign.Post("/in", handlers.SignIn)
	sign.Post("/out", middleware.JWTProtected(), handlers.SignOut)

	product := c.Group("/product", middleware.JWTProtected())
	product.Get("/", handlers.ListProduct)
	product.Get("/:id", handlers.GetProduct)
	product.Post("/", handlers.AddProduct)
	product.Patch("/", handlers.UpdateProduct)
	product.Delete("/", handlers.DeleteProduct)

	route.Get("/cart", func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "Cart", "ok")
	})

	route.Post("/checkout-session", func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "Checkout Session", "ok")
	})
}
