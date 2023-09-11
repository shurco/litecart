package handlers

import (
	"fmt"

	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/errors"
	"github.com/shurco/litecart/pkg/fsutil"
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

// AddProduct is ...
// [post] /api/_/products
func AddProduct(c *fiber.Ctx) error {
	db := queries.DB()
	request := &models.Product{}

	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	product, err := db.AddProduct(request)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product added", product)
}

// GetProduct is ...
// [get] /api/_/products/:product_id
func Product(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	db := queries.DB()

	product, err := db.Product(true, productID)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product info", product)
}

// UpdateProduct is ...
// [patch] /api/_/products/:product_id
func UpdateProduct(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	db := queries.DB()
	request := new(models.Product)
	request.ID = productID

	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	if err := db.UpdateProduct(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product updated", nil)
}

// DeleteProduct is ...
// [delete] /api/_/products/:product_id
func DeleteProduct(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	db := queries.DB()

	if err := db.DeleteProduct(productID); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product deleted", nil)
}

// UpdateProductActive is ...
// [patch] /api/_/products/:product_id/active
func UpdateProductActive(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	db := queries.DB()

	if err := db.UpdateActive(productID); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product active updated", nil)
}

// ProductImages
// [get] /api/_/products/:product_id/image
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
// [post] /api/_/products/:product_id/image
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
	filePath := fmt.Sprintf("./lc_uploads/%s", fileName)
	fileOrigName := file.Filename

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
		err := imaging.Save(resizedImage, fmt.Sprintf("./lc_uploads/%s_%s.%s", fileUUID, s.size, fileExt))
		if err != nil {
			return webutil.StatusBadRequest(c, err.Error())
		}
	}

	addedImage, err := db.AddImage(productID, fileUUID, fileExt, fileOrigName)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Image added", addedImage)
}

// DeleteProductImage is ...
// [delete] /api/_/products/:product_id/image/:image_id
func DeleteProductImage(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	imageID := c.Params("image_id")
	db := queries.DB()

	if err := db.DeleteImage(productID, imageID); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Image deleted", nil)
}

// ProductDigital
// [get] /api/_/products/:product_id/digital
func ProductDigital(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	db := queries.DB()

	digital, err := db.ProductDigital(productID)
	if err != nil {
		if err == errors.ErrProductNotFound {
			return webutil.StatusNotFound(c)
		}
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Product digital", digital)
}

// AddProductDigital is ...
// [post] /api/_/products/:product_id/digital
func AddProductDigital(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	db := queries.DB()

	fileTmp, _ := c.FormFile("document")
	if fileTmp != nil {
		fileUUID := uuid.New().String()
		fileExt := fsutil.ExtName(fileTmp.Filename)
		fileName := fmt.Sprintf("%s.%s", fileUUID, fileExt)
		filePath := fmt.Sprintf("./lc_digitals/%s", fileName)
		fileOrigName := fileTmp.Filename

		c.SaveFile(fileTmp, filePath)

		file, err := db.AddDigitalFile(productID, fileUUID, fileExt, fileOrigName)
		if err != nil {
			return webutil.StatusBadRequest(c, err.Error())
		}

		return webutil.Response(c, fiber.StatusOK, "Digital added", file)
	}

	data, err := db.AddDigitalData(productID, "")
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Digital added", data)
}

// UpdateProductDigital is ...
// [patch] /api/_/products/:product_id/digital/:digital_id
func UpdateProductDigital(c *fiber.Ctx) error {
	request := new(models.Data)
	request.ID = c.Params("digital_id")
	//request.Content = c.Params("digital_id")
	db := queries.DB()

	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	if err := db.UpdateDigital(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Digital updated", nil)
}

// DeleteProductDigital is ...
// [delete] /api/_/products/:product_id/digital/:digital_id
func DeleteProductDigital(c *fiber.Ctx) error {
	productID := c.Params("product_id")
	digitalID := c.Params("digital_id")
	db := queries.DB()

	if err := db.DeleteDigital(productID, digitalID); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Digital deleted", nil)
}
