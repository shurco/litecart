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

// AddPage is ...
// [post] /api/_/page/
func AddPage(c *fiber.Ctx) error {
	db := queries.DB()
	request := new(models.Page)

	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	page, err := db.AddPage(request)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Page added", page)
}

// UpdatePage is ...
// [patch] /api/_/pages/:page_id
func UpdatePage(c *fiber.Ctx) error {
	pageID := c.Params("page_id")
	db := queries.DB()
	request := new(models.Page)
	request.ID = pageID

	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	if err := db.UpdatePage(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Page updated", nil)
}

// DeletePage is ...
// [delete] /api/_/pages/:page_id
func DeletePage(c *fiber.Ctx) error {
	pageID := c.Params("page_id")
	db := queries.DB()

	if err := db.DeletePage(pageID); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Page deleted", nil)
}

// UpdatePageContent is ...
// [get] /api/_/page/:page_id/content
func UpdatePageContent(c *fiber.Ctx) error {
	db := queries.DB()
	pageID := c.Params("page_id")
	request := &models.Page{
		Core: models.Core{
			ID: pageID,
		},
	}

	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	if err := db.UpdatePageContent(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Page content updated", nil)
}

// UpdatePageActive is ...
// [patch] /api/_/page/:page_id/active
func UpdatePageActive(c *fiber.Ctx) error {
	pageID := c.Params("page_id")
	db := queries.DB()

	if err := db.UpdatePageActive(pageID); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Page active updated", nil)
}
