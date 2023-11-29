package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/mailer"
	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/errors"
	"github.com/shurco/litecart/pkg/update"
	"github.com/shurco/litecart/pkg/webutil"
)

// Version is ...
// [get] /api/_/version
func Version(c *fiber.Ctx) error {
	db := queries.DB()

	session, err := db.GetSession("update")
	if err != nil && err != sql.ErrNoRows {
		return webutil.StatusBadRequest(c, err.Error())
	}

	version := (*update.Version)(nil)
	if err == sql.ErrNoRows {
		version = update.VersionInfo()

		release, err := update.FetchLatestRelease(context.Background(), "shurco", "litecart")
		if err != nil {
			return webutil.StatusBadRequest(c, err.Error())
		}

		if version.CurrentVersion != release.Name {
			version.NewVersion = release.Name
			version.ReleaseURL = release.GetUrl()
		}

		if err := db.DeleteSession("update"); err != nil {
			return webutil.StatusBadRequest(c, err.Error())
		}

		json, _ := json.Marshal(version)
		expires := time.Now().Add(24 * time.Hour).Unix()
		if err := db.AddSession("update", string(json), expires); err != nil {
			return webutil.StatusBadRequest(c, err.Error())
		}
	}

	if session != "" {
		version = new(update.Version)
		json.Unmarshal([]byte(session), version)
	}

	return webutil.Response(c, fiber.StatusOK, "Version", version)
}

// Settings is ...
// [get] /api/_/settings
func Settings(c *fiber.Ctx) error {
	db := queries.DB()

	settings, err := db.Settings(true)
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
		return webutil.StatusBadRequest(c, err.Error())
	}
	section := ""
	for key := range sectionTmp {
		section = key
		break
	}

	switch section {
	case "stripe", "paypal", "spectrocoin":
		if err := json.Unmarshal(c.Body(), &request.PaymentSystem); err != nil {
			return webutil.StatusBadRequest(c, err.Error())
		}
	}

	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
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

	switch settingKey {
	case "password":
		return webutil.StatusNotFound(c)
	case "main", "social", "jwt", "webhook", "smtp", "stripe", "paypal", "spectrocoin":
		section, err := db.SettingBySection(settingKey)
		if err != nil {
			if err == errors.ErrSettingNotFound {
				return webutil.StatusNotFound(c)
			}
			return webutil.StatusBadRequest(c, err.Error())
		}
		return webutil.Response(c, fiber.StatusOK, "Setting section", section)
	default:
		setting, err := db.SettingValueByKey(settingKey)
		if err != nil {
			if err == errors.ErrSettingNotFound {
				return webutil.StatusNotFound(c)
			}
			return webutil.StatusBadRequest(c, err.Error())
		}
		return webutil.Response(c, fiber.StatusOK, "Setting", setting)
	}
}

// UpdateSettingByKey is ...
// [patch] /api/_/settings/:setting_key
func UpdateSettingByKey(c *fiber.Ctx) error {
	db := queries.DB()
	request := &models.SettingName{
		Key: c.Params("setting_key"),
	}

	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	if err := db.UpdateSettingValueByKey(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Setting updated", nil)
}

// TestLetter is ...
// [get] /api/_/test/letter/:letter_name
func TestLetter(c *fiber.Ctx) error {
	letter := c.Params("letter_name")

	if err := mailer.SendTestLetter(letter); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Test letter", "Message sent to your mailbox")
}
