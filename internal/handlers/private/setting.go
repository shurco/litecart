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

	session, err := db.GetSession(c.Context(), "update")
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

		if err := db.DeleteSession(c.Context(), "update"); err != nil {
			return webutil.StatusBadRequest(c, err.Error())
		}

		json, _ := json.Marshal(version)
		expires := time.Now().Add(24 * time.Hour).Unix()
		if err := db.AddSession(c.Context(), "update", string(json), expires); err != nil {
			return webutil.StatusBadRequest(c, err.Error())
		}
	}

	if session != "" {
		version = new(update.Version)
		json.Unmarshal([]byte(session), version)
	}

	return webutil.Response(c, fiber.StatusOK, "Version", version)
}

// GetSetting is ...
// [get] /api/_/settings/:setting_key
func GetSetting(c *fiber.Ctx) error {
	db := queries.DB()
	settingKey := c.Params("setting_key")

	var section any
	var err error

	switch settingKey {
	case "password":
		return webutil.StatusNotFound(c)
	case "main":
		section, err = db.GetSettingByGroup(c.Context(), &models.Main{})
	case "social":
		section, err = db.GetSettingByGroup(c.Context(), &models.Social{})
	case "auth":
		section, err = db.GetSettingByGroup(c.Context(), &models.Auth{})
	case "jwt":
		section, err = db.GetSettingByGroup(c.Context(), &models.JWT{})
	case "webhook":
		section, err = db.GetSettingByGroup(c.Context(), &models.Webhook{})
	case "payment":
		section, err = db.GetSettingByGroup(c.Context(), &models.Payment{})
	case "stripe":
		section, err = db.GetSettingByGroup(c.Context(), &models.Stripe{})
	case "paypal":
		section, err = db.GetSettingByGroup(c.Context(), &models.Paypal{})
	case "spectrocoin":
		section, err = db.GetSettingByGroup(c.Context(), &models.Spectrocoin{})
	case "mail":
		section, err = db.GetSettingByGroup(c.Context(), &models.Mail{})
	default:
		section, err = db.GetSettingByKey(c.Context(), settingKey)
	}

	if err != nil {
		if err == errors.ErrSettingNotFound {
			return webutil.StatusNotFound(c)
		}
		return webutil.StatusBadRequest(c, err.Error())
	}
	return webutil.Response(c, fiber.StatusOK, "Setting", section)
}

// UpdateSetting is ...
// [patch] /api/_/settings/:setting_key
func UpdateSetting(c *fiber.Ctx) error {
	db := queries.DB()
	settingKey := c.Params("setting_key")
	var request any

	switch settingKey {
	case "password":
		request = &models.Password{}
	case "main":
		request = &models.Main{}
	case "auth":
		request = &models.Auth{}
	case "jwt":
		request = &models.JWT{}
	case "social":
		request = &models.Social{}
	case "payment":
		request = &models.Payment{}
	case "stripe":
		request = &models.Stripe{}
	case "paypal":
		request = &models.Paypal{}
	case "spectrocoin":
		request = &models.Spectrocoin{}
	case "webhook":
		request = &models.Webhook{}
	case "mail":
		request = &models.Mail{}
	default:
		request = &models.SettingName{Key: settingKey}
	}

	// Parse the request body into the appropriate struct
	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	// Handle the password update separately if that's the case
	if settingKey == "password" {
		password := request.(*models.Password)
		if err := db.UpdatePassword(c.Context(), password); err != nil {
			return webutil.StatusBadRequest(c, err.Error())
		}
		return webutil.Response(c, fiber.StatusOK, "Password updated", nil)
	}

	// For default case where setting key doesn't match any predefined keys
	if _, ok := request.(*models.SettingName); ok {
		if err := db.UpdateSettingByKey(c.Context(), request.(*models.SettingName)); err != nil {
			return webutil.StatusBadRequest(c, err.Error())
		}
		return webutil.Response(c, fiber.StatusOK, "Setting key updated", nil)
	}

	// Update setting for all other cases
	if err := db.UpdateSettingByGroup(c.Context(), request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Setting group updated", nil)
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
