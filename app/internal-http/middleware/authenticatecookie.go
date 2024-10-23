package middleware

import (
	"fmt"
	"github.com/crspy2/license-panel/database"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func AuthenticateCookie(c *fiber.Ctx) error {
	if c.Method() != "GET" {
		headers := c.GetReqHeaders()
		fmt.Println(headers)
		csrfToken := headers["X-Csrf-Token"]
		if len(csrfToken) == 0 {
			return c.Status(http.StatusForbidden).
				JSON(fiber.Map{
					"status": http.StatusForbidden,
					"error":  "Forbidden — Invalid CSRF token provided",
				})
		}
	}

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
		return c.Status(http.StatusNotFound).
			JSON(fiber.Map{
				"status": http.StatusNotFound,
				"error":  "Session not found",
			})
	}

	c.Locals("session", s)
	return c.Next()
}
