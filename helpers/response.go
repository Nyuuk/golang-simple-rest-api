package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Response(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"code":    status,
		"message": message,
		"data":    data,
	})
}

func ResponseErrorInternal(c *fiber.Ctx, err error) error {
	return Response(c, fiber.StatusInternalServerError, "Internal server error", err.Error())
}

func ResponseErrorBadRequest(c *fiber.Ctx, err error) error {
	log.Error("Bad request: ", err)
	return Response(c, fiber.StatusBadRequest, "Bad request", nil)
}

func ResponseErrorNotFound(c *fiber.Ctx, err error) error {
	return Response(c, fiber.StatusNotFound, "Not found", nil)
}