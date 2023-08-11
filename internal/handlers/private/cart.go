package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/webutil"
)

// Checkouts is ...
// [get] /api/_/checkouts
func Checkouts(c *fiber.Ctx) error {
	db := queries.DB()

	products, err := db.Checkouts()
	if err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Checkouts", products)
}
