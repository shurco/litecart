package routes

import (
	"io/fs"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shurco/litecart/web"
)

// AdminRoutes sets up routes for the admin panel SPA.
func AdminRoutes(c *fiber.App) {
	embedAdmin, _ := fs.Sub(web.EmbedAdmin(), web.AdminBuildPath)

	// Handle admin routes with path prefix removal
	c.Use("/_", func(c *fiber.Ctx) error {
		originalPath := c.Path()
		normalizedPath := strings.TrimPrefix(originalPath, "/_")
		c.Path(normalizedPath)
		// Admin routes don't skip any paths - all paths under /_ are handled
		handler := setupSPAHandler(embedAdmin, func(string) bool { return false })
		err := handler(c)
		c.Path(originalPath)
		return err
	})
}
