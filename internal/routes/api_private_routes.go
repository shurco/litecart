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
	product.Get("/:id<len(15)>", handlers.Product)
	product.Post("/", handlers.AddProduct)
	product.Patch("/", handlers.UpdateProduct)
	product.Delete("/", handlers.DeleteProduct)
}
