package routes

import (
	"github.com/gofiber/fiber/v2"

	handlers "github.com/shurco/litecart/internal/handlers/private"
	"github.com/shurco/litecart/internal/middleware"
)

// ApiPrivateRoutes is ...
func ApiPrivateRoutes(c *fiber.App) {
	c.Post("/api/install", handlers.Install)

	sign := c.Group("/api/sign")
	sign.Post("/in", handlers.SignIn)
	sign.Post("/out", middleware.JWTProtected(), handlers.SignOut)

	product := c.Group("/api/products", middleware.JWTProtected())
	product.Get("/", handlers.Products)
	product.Post("/", handlers.AddProduct)
	product.Get("/:id<len(15)>", handlers.Product)
	product.Patch("/:id<len(15)>", handlers.UpdateProduct)
	product.Delete("/:id<len(15)>", handlers.DeleteProduct)

	product.Patch("/active/:id<len(15)>", handlers.UpdateProductActive)

	// stripe section
	product.Post("/stripe/:id<len(15)>", handlers.AddStripeProduct)
	product.Delete("/stripe/:id<len(15)>", handlers.DeleteStripeProduct)

	product.Get("/stripe/:id<len(15)>/check", handlers.CheckStripeProduct)
}
