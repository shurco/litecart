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
	embedAdminFolder, _ := fs.Sub(web.EmbedAdmin(), "admin/dist")
	c.Use("/_", filesystem.New(filesystem.Config{
		Root: http.FS(embedAdminFolder),
	}))

	embedAdminIndex, _ := fs.Sub(web.EmbedAdminIndex(), "admin/dist")
	c.Use("/_/*", filesystem.New(filesystem.Config{
		Root: http.FS(embedAdminIndex),
	}))
}
