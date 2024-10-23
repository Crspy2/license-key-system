package auth

import (
	"github.com/crspy2/license-panel/app/internal-http/utils"
	"github.com/crspy2/license-panel/database"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func RegisterRoute(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if len(username) < 3 {
		return c.Status(http.StatusNotAcceptable).
			JSON(fiber.Map{
				"status": http.StatusNotAcceptable,
				"error":  "Username must be at least 3 characters in length",
			})
	}

	if len(password) < 8 {
		return c.Status(http.StatusNotAcceptable).
			JSON(fiber.Map{
				"status": http.StatusNotAcceptable,
				"error":  "Password must be at least 8 characters in length",
			})
	}

	staff, _ := database.Client.Staff.GetByName(username)
	if staff != nil {
		return c.Status(http.StatusConflict).
			JSON(fiber.Map{
				"status": http.StatusConflict,
				"error":  "This username is already in use, please choose another one",
			})
	}

	hashedPassword, _ := utils.HashPassword(password)

	staff, err := database.Client.Staff.Create(username, hashedPassword)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"status": http.StatusInternalServerError,
				"error":  err.Error(),
			})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": "Successfully created staff account. Waiting for administration approval",
			"data":    *staff,
		})
}
