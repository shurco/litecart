package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/pkg/webutil"
)

// ListProduct is ...
func ListProduct(c *fiber.Ctx) error {
	return webutil.Response(c, fiber.StatusOK, "ListProduct", nil)
}

// GetProduct is ...
func GetProduct(c *fiber.Ctx) error {
	return webutil.Response(c, fiber.StatusOK, "GetProduct", nil)
}

// AddProduct is ...
func AddProduct(c *fiber.Ctx) error {
	return webutil.Response(c, fiber.StatusOK, "AddProduct", nil)
}

// UpdateProduct is ...
func UpdateProduct(c *fiber.Ctx) error {
	return webutil.Response(c, fiber.StatusOK, "UpdateProduct", nil)
}

// DeleteProduct is ...
func DeleteProduct(c *fiber.Ctx) error {
	return webutil.Response(c, fiber.StatusOK, "DeleteProduct", nil)
}
