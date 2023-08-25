package routes

import (
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/shurco/litecart/web"
)

// AdminRoutes is ...
func AdminRoutes(c *fiber.App) {
	embedAdmin, _ := fs.Sub(web.EmbedAdmin(), "admin/dist")
	c.Use("/_", filesystem.New(filesystem.Config{
		Root:         http.FS(embedAdmin),
		Index:        "index.html",
		NotFoundFile: "index.html",
		MaxAge:       3600,
	}))
}
