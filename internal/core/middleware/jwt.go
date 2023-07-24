package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/shurco/litecart/pkg/webutil"

	jwtMiddleware "github.com/gofiber/contrib/jwt"
)

// JWTProtected is ...
func JWTProtected(secret string) func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:   jwtMiddleware.SigningKey{Key: []byte(secret)},
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
