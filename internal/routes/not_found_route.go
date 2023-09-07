package routes

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/webutil"
)

// NotFoundRoute func for describe 404 Error route.
func NotFoundRoute(a *fiber.App, noSite bool) {
	a.Use(func(c *fiber.Ctx) error {
		db := queries.DB()
		if db.IsPage(c.Path()[1:]) {
			return c.Render("pages", nil, "layouts/main")
		}

		if strings.HasPrefix(c.Path(), "/api") {
			return webutil.StatusNotFound(c)
		}
		if strings.HasPrefix(c.Path(), "/_") {
			return c.Next()
		}

		if noSite {
			return c.Next()
		}
		return c.Status(fiber.StatusNotFound).Render("404", nil, "layouts/clear")
	})
}
