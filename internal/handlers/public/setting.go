package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/logging"
	"github.com/shurco/litecart/pkg/webutil"
)

// Ping returns a pong response for health checks.
// [get] /ping
func Ping(c *fiber.Ctx) error {
	return webutil.Response(c, fiber.StatusOK, "Pong", nil)
}

// Settings returns public settings including main, social, payment, and pages.
// [get] /api/settings
func Settings(c *fiber.Ctx) error {
	db := queries.DB()
	log := logging.New()

	settingMain, err := queries.GetSettingByGroup[models.Main](c.Context(), db)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	settingSocial, err := queries.GetSettingByGroup[models.Social](c.Context(), db)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	settingPayment, err := queries.GetSettingByGroup[models.Payment](c.Context(), db)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	pages, _, err := db.ListPages(c.Context(), false, 0, 0)
	if err != nil {
		log.ErrorStack(err)
		return webutil.StatusInternalServerError(c)
	}

	return webutil.Response(c, fiber.StatusOK, "Settings", map[string]any{
		"main": map[string]string{
			"site_name": settingMain.SiteName,
			"domain":    settingMain.Domain,
			"currency":  settingPayment.Currency,
		},
		"socials": settingSocial,
		"pages":   pages,
	})
}
