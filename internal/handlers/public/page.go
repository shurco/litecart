package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/errors"
	"github.com/shurco/litecart/pkg/webutil"
)

// Pages is ...
// [get] /api/pages
func Pages(c *fiber.Ctx) error {
	db := queries.DB()

	pages, err := db.ListPages(false)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Pages", pages)
}

// Page is ...
// [get] /api/page/:page_slug
func Page(c *fiber.Ctx) error {
	pageSlug := c.Params("page_slug")
	db := queries.DB()

	page, err := db.Page(pageSlug)
	if err != nil {
		if err == errors.ErrPageNotFound {
			return webutil.StatusNotFound(c)
		}
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Page content", page)
}
