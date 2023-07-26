package routes

import (
	"github.com/gofiber/fiber/v2"

	handlers "github.com/shurco/litecart/internal/handlers/private"
	"github.com/shurco/litecart/internal/middleware"
)

// ApiPrivateRoutes is ...
// route have path '/api'
func ApiPrivateRoutes(route fiber.Router) {
	route.Post("/install", handlers.Install)

	sign := route.Group("/sign")
	sign.Post("/in", handlers.SignIn)
	sign.Post("/out", middleware.JWTProtected(), handlers.SignOut)

	product := route.Group("/products", middleware.JWTProtected())
	product.Get("/:id<len(15)>", handlers.Product)
	product.Post("/", handlers.AddProduct)
	product.Patch("/", handlers.UpdateProduct)
	product.Delete("/", handlers.DeleteProduct)
}
