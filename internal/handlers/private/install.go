package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/models"
	"github.com/shurco/litecart/internal/queries"
	"github.com/shurco/litecart/pkg/validator"
	"github.com/shurco/litecart/pkg/webutil"
)

// Install is ...
func Install(c *fiber.Ctx) error {
	db := queries.DB()
	request := new(models.Install)

	if err := c.BodyParser(request); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	if err := validator.Struct(request); err != nil {
		return webutil.StatusBadRequest(c, err)
	}

	if err := db.Install(request); err != nil {
		return webutil.StatusBadRequest(c, err.Error())
	}

	return webutil.Response(c, fiber.StatusOK, "Cart installed", nil)
}
