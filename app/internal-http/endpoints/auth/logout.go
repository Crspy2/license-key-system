package auth

import (
	"github.com/crspy2/license-panel/database"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

func LogoutRoute(c *fiber.Ctx) error {
	sessionId := c.Cookies("session_token")
	if sessionId == "" {
		return c.Status(http.StatusUnauthorized).
			JSON(fiber.Map{
				"status": http.StatusUnauthorized,
				"error":  "No session token could be found",
			})
	}

	session, err := database.Client.Session.Get(sessionId, c.IP())
	if err != nil {
		return c.Status(http.StatusForbidden).
			JSON(fiber.Map{
				"status": http.StatusForbidden,
				"error":  "No valid session token could be found",
			})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "csrf_token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})

	c.Cookie(&fiber.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})

	err = database.Client.Session.Delete(*session)
	if err != nil {
		return c.Status(http.StatusForbidden).
			JSON(fiber.Map{
				"status": http.StatusForbidden,
				"error":  "No valid session token could be found",
			})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": "User has been signed out, and their session has been deleted",
		})
}
