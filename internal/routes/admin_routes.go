package routes

import (
	"github.com/gofiber/fiber/v2"
)

// AdminRoutes is ...
func AdminRoutes(c *fiber.App) {
	c.Static("/_", "../web/admin/dist", fiber.Static{
		Compress: true,
	})

	c.Get("/_/*", func(ctx *fiber.Ctx) error {
		return ctx.SendFile("../web/admin/dist/index.html", true)
	})
}
