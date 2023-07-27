package routes

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shurco/litecart/pkg/webutil"
)

// NotFoundRoute func for describe 404 Error route.
func NotFoundRoute(a *fiber.App) {
	a.Use(func(c *fiber.Ctx) error {
		path := strings.Split(c.Path(), "/")[1]
		if path == "api" {
			return webutil.Response(c, fiber.StatusNotFound, "Not Found", nil)
		}
		return c.Status(fiber.StatusNotFound).Render("error/404", fiber.Map{}, "site/layouts/clear")
	})
}
