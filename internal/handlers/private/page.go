package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/errors"
	"github.com/shurco/litecart/pkg/logging"
	"github.com/shurco/litecart/pkg/webutil"
)

// Pages returns a list of all pages.
// [get] /api/_/pages
func Pages(c *fiber.Ctx) error {
	db := queries.DB()
	log := logging.New()

	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 20)
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	offset := (page - 1) * limit

	pages, total, err := db.ListPages(c.Context(), true, limit, offset)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	return webutil.Response(c, fiber.StatusOK, "Pages", map[string]any{
		"pages": pages,
		"total":  total,
		"page":   page,
		"limit":  limit,
	})
}

// GetPage returns a single page by ID.
// [get] /api/_/pages/:page_id
func GetPage(c *fiber.Ctx) error {
	pageID := c.Params("page_id")
	db := queries.DB()
	log := logging.New()

	page, err := db.PageByID(c.Context(), pageID)
	if err != nil {
		if err == errors.ErrPageNotFound {
			return webutil.StatusNotFound(c)
		}
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	return webutil.Response(c, fiber.StatusOK, "Page", page)
}

// AddPage creates a new page.
// [post] /api/_/pages
func AddPage(c *fiber.Ctx) error {
	db := queries.DB()
	log := logging.New()
	request := new(models.Page)

	if err := c.BodyParser(request); err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}

	page, err := db.AddPage(c.Context(), request)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	return webutil.Response(c, fiber.StatusOK, "Page added", page)
}

// UpdatePage updates an existing page.
// [patch] /api/_/pages/:page_id
func UpdatePage(c *fiber.Ctx) error {
	pageID := c.Params("page_id")
	db := queries.DB()
	log := logging.New()
	request := new(models.Page)
	request.ID = pageID

	if err := c.BodyParser(request); err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}

	if err := db.UpdatePage(c.Context(), request); err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	return webutil.Response(c, fiber.StatusOK, "Page updated", nil)
}

// DeletePage deletes a page by ID.
// [delete] /api/_/pages/:page_id
func DeletePage(c *fiber.Ctx) error {
	pageID := c.Params("page_id")
	db := queries.DB()
	log := logging.New()

	if err := db.DeletePage(c.Context(), pageID); err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	return webutil.Response(c, fiber.StatusOK, "Page deleted", nil)
}

// UpdatePageContent updates the content of a page.
// [patch] /api/_/pages/:page_id/content
func UpdatePageContent(c *fiber.Ctx) error {
	db := queries.DB()
	log := logging.New()
	pageID := c.Params("page_id")

	request := &models.Page{
		Core: models.Core{
			ID: pageID,
		},
	}

	if err := c.BodyParser(request); err != nil {
		log.ErrorStack(err)
		return webutil.StatusBadRequest(c, err.Error())
	}

	if err := db.UpdatePageContent(c.Context(), request); err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	return webutil.Response(c, fiber.StatusOK, "Page content updated", nil)
}

// UpdatePageActive updates the active status of a page.
// [patch] /api/_/pages/:page_id/active
func UpdatePageActive(c *fiber.Ctx) error {
	pageID := c.Params("page_id")
	db := queries.DB()
	log := logging.New()

	if err := db.UpdatePageActive(c.Context(), pageID); err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	// Get updated page to return with updated timestamp
	page, err := db.PageByID(c.Context(), pageID)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	return webutil.Response(c, fiber.StatusOK, "Page active updated", page)
}
