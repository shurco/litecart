package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/webutil"
)

// Products is ...
// [get] /api/products
func Products(c *fiber.Ctx) error {
	db := queries.DB()

	products, err := db.ListProducts(c.Context(), false)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Products", products)
}

// GetProduct is ...
// [get] /api/products/:product_id
func Product(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	db := queries.DB()

	product, err := db.Product(c.Context(), false, productID)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product info", product)
}
