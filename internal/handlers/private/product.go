package handlers

import (
	"fmt"

	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/fsutil"
	"github.com/shurco/litecart/pkg/validator"
	"github.com/shurco/litecart/pkg/webutil"
)

// Products is ...
// [get] /api/_/products
func Products(c *fiber.Ctx) error {
	db := queries.DB()

	products, err := db.ListProducts(true)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Products", products)
}

// GetProduct is ...
// [get] /api/_/products/:product_id
func Product(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	db := queries.DB()

	product, err := db.Product(productID)
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
// [patch] /api/:product_id/active
func UpdateProductActive(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	db := queries.DB()

	if err := db.UpdateActive(productID); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product active updated", nil)
}

// ProductImages
// [get] /api/:product_id/image
func ProductImages(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	db := queries.DB()

	images, err := db.ProductImages(productID)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product images", images)
}

// AddProductImage is ...
// [post] /api/:product_id/image
func AddProductImage(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	db := queries.DB()

	file, err := c.FormFile("document")
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	validMIME := false
	validMIMETypes := []string{"image/png", "image/jpeg"}
	for _, mime := range validMIMETypes {
		if mime == file.Header["Content-Type"][0] {
			validMIME = true
		}
	}
	if !validMIME {
		return webutil.StatusBadRequest(c, "file format not supported")
	}

	fileUUID := uuid.New().String()
	fileExt := fsutil.ExtName(file.Filename)
	fileName := fmt.Sprintf("%s.%s", fileUUID, fileExt)
	filePath := fmt.Sprintf("./uploads/%s", fileName)

	c.SaveFile(file, filePath)

	fileSource, err := imaging.Open(filePath)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	sizes := []struct {
		size string
		dim  int
	}{
		{"sm", 147},
		{"md", 400},
	}

	for _, s := range sizes {
		resizedImage := imaging.Fill(fileSource, s.dim, s.dim, imaging.Center, imaging.Lanczos)
		err := imaging.Save(resizedImage, fmt.Sprintf("./uploads/%s_%s.%s", fileUUID, s.size, fileExt))
		if err != nil {
			return webutil.StatusBadRequest(c, err.Error())
		}
	}

	addedImage, err := db.AddImage(productID, fileUUID, fileExt)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Image added", addedImage)
}

// DeleteProductImage is ...
// [delete] /api/:product_id/image/:image_id
func DeleteProductImage(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	imageID := c.Params("image_id")
	db := queries.DB()

	if err := db.DeleteImage(productID, imageID); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Image deleted", nil)
}

// DeleteProduct is ...
func DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	db := queries.DB()

	if err := db.DeleteProduct(productID); err != nil {
		if err == queries.StripeProductNotFound {
			return webutil.StatusNotFound(c)
		}
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product deleted", nil)
}

// AddStripeProduct is ...
func AddStripeProduct(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	db := queries.DB()

	stripeID, err := db.AddStripeProduct(productID)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	if err := db.UpdateStripeProduct(productID, stripeID); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Stripe id", stripeID)
}

// DeleteStripeProduct is ...
func DeleteStripeProduct(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	db := queries.DB()

	if err := db.DeleteStripeProduct(productID); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Stripe product deleted", nil)
}

// CheckStripeProduct is ...
func CheckStripeProduct(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	db := queries.DB()

	return webutil.Response(c, fiber.StatusOK, "Stripe product check", db.IsStripeProduct(productID))
}
