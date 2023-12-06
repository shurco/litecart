package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/jwtutil"
	"github.com/shurco/litecart/pkg/security"
	"github.com/shurco/litecart/pkg/webutil"
)

// SignIn is ...
// [post] /api/sign/in
func SignIn(c *fiber.Ctx) error {
	db := queries.DB()
	request := new(models.SignIn)

	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	if err := request.Validate(); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	passwordHash, err := db.GetPasswordByEmail(c.Context(), request.Email)
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	compareUserPassword := security.ComparePasswords(passwordHash, request.Password)
	if !compareUserPassword {
		return webutil.StatusBadRequest(c, "wrong user email address or password")
	}

	// Generate a new pair of access and refresh tokens.
	_settingJWT, err := db.GetSetting(c.Context(), &models.JWT{})
	if err != nil {
		return err
	}
	settingJWT := _settingJWT.(*models.JWT)

	userID := uuid.New()
	expires := time.Now().Add(time.Hour * time.Duration(settingJWT.ExpireHours)).Unix()
	token, err := jwtutil.GenerateNewToken(settingJWT.Secret, userID.String(), expires, nil)
	if err != nil {
		return webutil.Response(c, fiber.StatusInternalServerError, "Internal server error", err.Error())
	}

	// Add session record
	if err := db.AddSession(c.Context(), userID.String(), "admin", expires); err != nil {
		return webutil.Response(c, fiber.StatusInternalServerError, "Failed to save token", err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: "lax",
	})

	return webutil.StatusOK(c, "Token", token)
}

// SignOut is ...
// [post] /api/sign/out
func SignOut(c *fiber.Ctx) error {
	db := queries.DB()
	_settingJWT, err := db.GetSetting(c.Context(), &models.JWT{})
	if err != nil {
		return err
	}
	settingJWT := _settingJWT.(*models.JWT)

	claims, err := jwtutil.ExtractTokenMetadata(c, settingJWT.Secret)
	if err != nil {
		return webutil.Response(c, fiber.StatusInternalServerError, "Internal server error", err.Error())
	}

	if err := db.DeleteSession(c.Context(), claims.ID); err != nil {
		return webutil.Response(c, fiber.StatusInternalServerError, "Failed to delete token", err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Expires: time.Now().Add(-(time.Hour * 2)),
		// HTTPOnly: true,
		SameSite: "lax",
	})

	return c.SendStatus(fiber.StatusNoContent)
}
