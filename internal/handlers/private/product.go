package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/validator"
	"github.com/shurco/litecart/pkg/webutil"
)

// Products is ...
// [get] /api/products
func Products(c *fiber.Ctx) error {
	db := queries.DB()

	products, err := db.ListProducts()
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Products", products)
}

// GetProduct is ...
// [get] /api/products/:id
func Product(c *fiber.Ctx) error {
	id := c.Params("id")
	db := queries.DB()

	product, err := db.Product(id)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product info", product)
}

// AddProduct is ...
func AddProduct(c *fiber.Ctx) error {
	db := queries.DB()
	request := new(models.Product)

	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	if err := validator.Struct(request); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	if err := db.AddProduct(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product added", nil)
}

// UpdateProduct is ...
func UpdateProduct(c *fiber.Ctx) error {
	db := queries.DB()
	request := new(models.Product)

	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	if err := validator.Struct(request); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	if err := db.UpdateProduct(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product updated", nil)
}

// UpdateProductActive is ...
// [patch] /api/active/:id
func UpdateProductActive(c *fiber.Ctx) error {
	id := c.Params("id")
	db := queries.DB()

	if err := db.UpdateActive(id); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product active updated", nil)
}

// DeleteProduct is ...
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := queries.DB()

	if err := db.DeleteProduct(id); err != nil {
		if err == queries.StripeProductNotFound {
			return webutil.StatusNotFound(c)
		}
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product deleted", nil)
}

// AddStripeProduct is ...
func AddStripeProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := queries.DB()

	stripeID, err := db.AddStripeProduct(id)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	if err := db.UpdateStripeProduct(id, stripeID); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Stripe id", stripeID)
}

// DeleteStripeProduct is ...
func DeleteStripeProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := queries.DB()

	if err := db.DeleteStripeProduct(id); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Stripe product deleted", nil)
}

// CheckStripeProduct is ...
func CheckStripeProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := queries.DB()

	return webutil.Response(c, fiber.StatusOK, "Stripe product check", db.IsStripeProduct(id))
}
