package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/webutil"
)

// Socials is ...
// [get] /api/socials
func Socials(c *fiber.Ctx) error {
	db := queries.DB()

	socials, err := db.ListSocials()
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Socials", socials)
}

// Ping is ...
// [get] /ping
func Ping(c *fiber.Ctx) error {
	return webutil.Response(c, fiber.StatusOK, "Pong", nil)
}
