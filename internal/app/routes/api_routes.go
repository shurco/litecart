package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/app/controllers"
	"github.com/shurco/litecart/internal/app/queries"
	"github.com/shurco/litecart/internal/core"
	"github.com/shurco/litecart/internal/core/middleware"
	"github.com/shurco/litecart/pkg/webutil"
)

// ApiRoutes is ...
func ApiRoutes(c *core.Core, q *queries.Base) {
	route := c.Fiber.Group("/api")
	route.Post("/install", controllers.Install(q))

	sign := c.Fiber.Group("/api/sign")
	sign.Post("/in", controllers.SignIn(q, c.JWT))
	sign.Post("/out", middleware.JWTProtected(c.JWT.Secret), controllers.SignOut(q, c.JWT.Secret))

	product := c.Fiber.Group("/product", middleware.JWTProtected(c.JWT.Secret))
	product.Get("/", controllers.ListProduct(q))
	product.Get("/:id", controllers.GetProduct(q))
	product.Post("/", controllers.AddProduct(q))
	product.Patch("/", controllers.UpdateProduct(q))
	product.Delete("/", controllers.DeleteProduct(q))

	route.Get("/cart", func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "Cart", "ok")
	})

	route.Post("/checkout-session", func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "Checkout Session", "ok")
	})
}
