package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/shurco/litecart/internal/app/models"
	"github.com/shurco/litecart/internal/app/queries"
	"github.com/shurco/litecart/pkg/validator"
	"github.com/shurco/litecart/pkg/webutil"
)

// Install is ...
func Install(q *queries.Base) fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(models.Install)

		if err := c.BodyParser(request); err != nil {
			return webutil.StatusBadRequest(c, err)
		}

		if err := validator.Struct(request); err != nil {
			return webutil.StatusBadRequest(c, err)
		}

		err := q.Install(request)
		if err != nil {
			return webutil.StatusBadRequest(c, err.Error())
		}

		return webutil.Response(c, fiber.StatusOK, "Cart installed", nil)
	}
}
