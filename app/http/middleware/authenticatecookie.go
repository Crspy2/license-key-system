package middleware

import (
	"crspy2/licenses/database"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

func AuthenticateCookie(c *fiber.Ctx) error {
	sessionId := c.Cookies("session_token")
	if sessionId == "" {
		return c.Status(http.StatusUnauthorized).
			JSON(fiber.Map{
				"status": http.StatusUnauthorized,
				"error":  "No session token could be found",
			})
	}

	s, err := database.Client.Session.Get(sessionId, c.IP())
	if err != nil {
		c.Cookie(&fiber.Cookie{
			Name:    "csrf_token",
			Value:   "",
			Expires: time.Now().Add(-time.Hour),
		})

		return c.Status(http.StatusNotFound).
			JSON(fiber.Map{
				"status": http.StatusNotFound,
				"error":  "Session not found",
			})
	}

	c.Locals("session", s)
	return c.Next()
}
