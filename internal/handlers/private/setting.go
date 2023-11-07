package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/errors"
	"github.com/shurco/litecart/pkg/update"
	"github.com/shurco/litecart/pkg/webutil"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/webhookendpoint"
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

	prevSettings, err := db.Settings(true)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	request := new(models.Setting)

	request.Stripe.WebhookId = prevSettings.Stripe.WebhookId
	request.Stripe.WebhookSecretKey = prevSettings.Stripe.SecretKey

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

	if prevSettings.Stripe.WebhookUrl != request.Stripe.WebhookUrl {
		if prevSettings.Stripe.SecretKey == "" && request.Stripe.SecretKey == "" {
			return webutil.Response(c, fiber.StatusBadRequest, "missing parameter", nil)
		}

		if prevSettings.Stripe.SecretKey == "" {
			stripe.Key = request.Stripe.SecretKey
		} else {
			stripe.Key = prevSettings.Stripe.SecretKey
		}

		if request.Stripe.WebhookUrl == "" { // if url removed, delete from stripe account
			we, _ := webhookendpoint.Del(
				prevSettings.Stripe.WebhookId,
				nil,
			)
			if we.Deleted {
				request.Stripe.WebhookId = ""
				request.Stripe.WebhookSecretKey = ""
			}
		} else { // update values in stripe account
			params := &stripe.WebhookEndpointParams{
				EnabledEvents: []*string{
					stripe.String("checkout.session.completed"),
					stripe.String("payment_intent.succeeded"),
					stripe.String("payment_intent.payment_failed"),
				},
				URL: stripe.String(request.Stripe.WebhookUrl),
			}
			if prevSettings.Stripe.WebhookUrl == "" { // if there wasn't any previous webhook url
				we, _ := webhookendpoint.New(params)

				if we.LastResponse.Status == "200 OK" {
					var res map[string]interface{}

					err := json.Unmarshal(we.LastResponse.RawJSON, &res)
					if err != nil {
						return webutil.Response(c, fiber.StatusBadRequest, "internal error", nil)
					}
					request.Stripe.WebhookUrl = res["url"].(string)
					request.Stripe.WebhookSecretKey = res["secret"].(string)
					request.Stripe.WebhookId = res["id"].(string)
				}
			} else {
				webhookendpoint.Update(
					prevSettings.Stripe.WebhookId,
					params,
				)
			}
		}
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
