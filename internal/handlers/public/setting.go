package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/models"
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

	settingMain, err := db.GetSetting(c.Context(), &models.Main{})
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}
	_settingMain := settingMain.(*models.Main)

	settingSocial, err := db.GetSetting(c.Context(), &models.Social{})
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	pages, err := db.ListPages(c.Context(), false)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Settings", map[string]any{
		"main": map[string]string{
			"site_name": _settingMain.SiteName,
			"domain":    _settingMain.Domain,
			"currency":  _settingMain.Currency,
		},
		"socials": settingSocial.(*models.Social),
		"pages":   pages,
	})
}
