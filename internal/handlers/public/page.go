package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/queries"
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
// [get] /api/page/:page_url
func Page(c *fiber.Ctx) error {
	pageUrl := c.Params("page_url")
	db := queries.DB()

	page, err := db.Page(pageUrl)
	if err != nil {
		if err.Error() == "page not found" {
			return webutil.StatusNotFound(c)
		}
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Page content", page)
}
