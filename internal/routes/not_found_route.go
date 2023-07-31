package routes

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shurco/litecart/pkg/webutil"
)

// NotFoundRoute func for describe 404 Error route.
func NotFoundRoute(a *fiber.App) {
	a.Use(func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.Path(), "/api") {
			return webutil.Response(c, fiber.StatusNotFound, "Not Found", nil)
		}
		if strings.HasPrefix(c.Path(), "/_") {
			return c.Status(fiber.StatusNotFound).Render("admin/404", fiber.Map{}, "admin/layouts/clear")
		}
		return c.Status(fiber.StatusNotFound).Render("site/404", fiber.Map{}, "site/layouts/clear")
	})
}
