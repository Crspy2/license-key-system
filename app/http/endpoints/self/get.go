package self

import (
	"github.com/crspy2/license-panel/database"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func SelfRoute(c *fiber.Ctx) error {
	session := c.Locals("session").(*database.SessionModal)

	if session == nil {
		return c.Status(http.StatusUnauthorized).
			JSON(fiber.Map{
				"status": http.StatusUnauthorized,
				"error":  "No session was found",
			})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": "User session retrieved",
			"data":    session,
		})
}
