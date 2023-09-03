package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/webutil"
)

// Settings is ...
// [get] /api/_/settings
func Settings(c *fiber.Ctx) error {
	db := queries.DB()

	settings, err := db.Settings()
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Settings", settings)
}

// UpdateSettings is ...
// [get] /api/_/settings
func UpdateSettings(c *fiber.Ctx) error {
	db := queries.DB()
	request := new(models.Setting)

	sectionTmp := map[string]any{}
	if err := json.Unmarshal(c.Body(), &sectionTmp); err != nil {
		return webutil.StatusBadRequest(c, err)
	}
	section := ""
	for key := range sectionTmp {
		section = key
		break
	}

	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	if err := db.UpdateSettings(request, section); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Settings updated", nil)
}
