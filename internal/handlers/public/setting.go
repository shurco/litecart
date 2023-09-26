package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/webutil"
)

// Ping is ...
// [get] /ping
func Ping(c *fiber.Ctx) error {
	return webutil.Response(c, fiber.StatusOK, "Pong", nil)
}

// Settings is ...
// [get] /api/settings
func Settings(c *fiber.Ctx) error {
	db := queries.DB()

	settings, err := db.Settings(false)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	pages, err := db.ListPages(false)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Settings", map[string]any{
		"main": map[string]string{
			"site_name": settings.Main.SiteName,
			"domain":    settings.Main.Domain,
			"currency":  settings.Main.Currency,
		},
		"socials": settings.Social,
		"pages":   pages,
	})
}
