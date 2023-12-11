package webutil

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

// HTTPResponse represents response body of API
type HTTPResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  any    `json:"result,omitempty"`
}

// Response is a takes in a Fiber context object, an HTTP status code, a message string and some data.
func Response(c *fiber.Ctx, status int, message string, data any) error {
	if len(message) > 0 {
		return c.Status(status).JSON(HTTPResponse{
			Success: status == fiber.StatusOK,
			Message: message,
			Result:  data,
		})
	}

	return c.Status(status).JSON(data)
}

// StatusOK is ...
func StatusOK(c *fiber.Ctx, message string, data any) error {
	return Response(c, fiber.StatusOK, message, data)
}

// StatusNotFound is ...
func StatusNotFound(c *fiber.Ctx) error {
	return Response(c, fiber.StatusNotFound, utils.StatusMessage(fiber.StatusNotFound), nil)
}

// StatusBadRequest is ...
func StatusBadRequest(c *fiber.Ctx, data any) error {
	return Response(c, fiber.StatusBadRequest, utils.StatusMessage(fiber.StatusBadRequest), data)
}

// StatusInternalServerError is ...
func StatusInternalServerError(c *fiber.Ctx) error {
	return Response(c, fiber.StatusInternalServerError, utils.StatusMessage(fiber.StatusInternalServerError), nil)
}
