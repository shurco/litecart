package middleware

import (
	"strings"

	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/webutil"
)

// JWTProtected is ...
func JWTProtected() func(*fiber.Ctx) error {
	db := queries.DB()
	settingJWT, _ := db.SettingJWT()

	config := jwtMiddleware.Config{
		SigningKey:   jwtMiddleware.SigningKey{Key: []byte(settingJWT.Secret)},
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
		TokenLookup:  "cookie:token",
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	path := strings.Split(c.Path(), "/")[1]
	if path == "api" {
		if err.Error() == "Missing or malformed token" {
			return webutil.Response(c, fiber.StatusBadRequest, "Bad request", err.Error())
		}
		return webutil.Response(c, fiber.StatusUnauthorized, "Unauthorized", err.Error())
	}

	return c.Redirect("/_/signin")
}
