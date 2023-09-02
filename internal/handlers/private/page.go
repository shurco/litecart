package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/webutil"
)

// Pages is ...
// [get] /api/_/pages
func Pages(c *fiber.Ctx) error {
	db := queries.DB()

	pages, err := db.ListPages(true)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Pages", pages)
}

// UpdatePage is ...
// [get] /api/_/page/:page_id
func UpdatePage(c *fiber.Ctx) error {
	db := queries.DB()
	pageID := c.Params("page_id")
	request := &models.Page{
		ID: pageID,
	}

	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	if err := db.UpdatePage(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Page updated", nil)
}
