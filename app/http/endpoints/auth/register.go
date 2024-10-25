package auth

import (
	"crspy2/licenses/app/http/utils"
	"crspy2/licenses/database"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func RegisterRoute(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if len(username) < 3 {
		return c.Status(http.StatusNotAcceptable).
			JSON(utils.InternalResponse{
				Success: false,
				Error:   "Username must be at least 3 characters in length",
			})
	}

	if len(password) < 8 {
		return c.Status(http.StatusNotAcceptable).
			JSON(utils.InternalResponse{
				Success: false,
				Error:   "Password must be at least 8 characters in length",
			})
	}

	staff, _ := database.Client.Staff.GetByName(username)
	if staff != nil {
		return c.Status(http.StatusConflict).
			JSON(utils.InternalResponse{
				Success: false,
				Error:   "This username is already in use, please choose another one",
			})
	}

	hashedPassword, _ := utils.HashPassword(password)

	staff, err := database.Client.Staff.Create(username, hashedPassword)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(utils.InternalResponse{
				Success: false,
				Error:   err.Error(),
			})
	}

	return c.Status(http.StatusOK).
		JSON(utils.InternalResponse{
			Success: false,
			Message: "Successfully created staff account. Waiting for administration approval",
			Data:    *staff,
		})
}
