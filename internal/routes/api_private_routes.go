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

	settings := c.Group("/api/_/settings", middleware.JWTProtected())
	settings.Get("/", handlers.Settings)
	settings.Patch("/", handlers.UpdateSettings)

	pages := c.Group("/api/_/pages", middleware.JWTProtected())
	pages.Get("/", handlers.Pages)
	pages.Post("/", handlers.AddPage)
	pages.Patch("/:page_id<len(15)>", handlers.UpdatePage)
	pages.Delete("/:page_id<len(15)>", handlers.DeletePage)
	pages.Patch("/:page_id<len(15)>/content", handlers.UpdatePageContent)
	pages.Patch("/:page_id<len(15)>/active", handlers.UpdatePageActive)

	product := c.Group("/api/_/products", middleware.JWTProtected())
	product.Get("/", handlers.Products)
	product.Post("/", handlers.AddProduct)
	product.Get("/:product_id<len(15)>", handlers.Product)
	product.Patch("/:product_id<len(15)>", handlers.UpdateProduct)
	product.Delete("/:product_id<len(15)>", handlers.DeleteProduct)
	product.Patch("/:product_id<len(15)>/active", handlers.UpdateProductActive)
	product.Get("/:product_id<len(15)>/image", handlers.ProductImages)
	product.Post("/:product_id<len(15)>/image", handlers.AddProductImage)
	product.Delete("/:product_id<len(15)>/image/:image_id<len(15)>", handlers.DeleteProductImage)

	// checkouts
	c.Get("/api/_/checkouts", handlers.Checkouts)
}
