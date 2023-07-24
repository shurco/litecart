package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/app/queries"
	"github.com/shurco/litecart/pkg/webutil"
)

// ListProduct is ...
func ListProduct(q *queries.Base) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "ListProduct", nil)
	}
}

// GetProduct is ...
func GetProduct(q *queries.Base) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "GetProduct", nil)
	}
}

// AddProduct is ...
func AddProduct(q *queries.Base) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "AddProduct", nil)
	}
}

// UpdateProduct is ...
func UpdateProduct(q *queries.Base) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "UpdateProduct", nil)
	}
}

// DeleteProduct is ...
func DeleteProduct(q *queries.Base) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return webutil.Response(c, fiber.StatusOK, "DeleteProduct", nil)
	}
}
