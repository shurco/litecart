package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/errors"
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
// [patch] /api/_/settings
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

// SettingByKey is ...
// [get] /api/_/settings/:setting_key
func SettingByKey(c *fiber.Ctx) error {
	db := queries.DB()
	settingKey := c.Params("setting_key")

	if settingKey == "password" {
		return webutil.StatusNotFound(c)
	}

	setting, err := db.SettingValueByKey(settingKey)
	if err != nil {
		if err == errors.ErrSettingNotFound {
			return webutil.StatusNotFound(c)
		}
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Setting", setting)
}

// UpdateSettingByKey is ...
// [patch] /api/_/settings/:setting_key
func UpdateSettingByKey(c *fiber.Ctx) error {
	db := queries.DB()
	request := &models.SettingName{
		Key: c.Params("setting_key"),
	}

	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	if err := db.UpdateSettingValueByKey(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Setting updated", nil)
}

// TestLetter is ...
// [get] /api/_/settings/test/:letter_name
func TestLetter(c *fiber.Ctx) error {
	db := queries.DB()
	letter := c.Params("letter_name")

	if err := db.SettingTestLetter(letter); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Test mail", "Message sent to your mailbox")
}
