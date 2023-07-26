package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/webutil"
)

// GetProduct is ...
// [get] /api/products/:id
func Product(c *fiber.Ctx) error {
	id := c.Params("id")
	db := queries.DB()

	product, err := db.Product(id)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product information", product)
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
