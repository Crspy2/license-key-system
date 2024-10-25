package auth

import (
	"crspy2/licenses/app/http/utils"
	"crspy2/licenses/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

func LogoutRoute(c *fiber.Ctx) error {
	fmt.Println("TEST")
	sessionId := c.Cookies("session_token")
	if sessionId == "" {
		return c.Status(http.StatusUnauthorized).
			JSON(utils.InternalResponse{
				Success: false,
				Error:   "No session token could be found",
			})
	}

	fmt.Println("TEST")
	session, err := database.Client.Session.Get(sessionId, c.IP())
	if err != nil {
		return c.Status(http.StatusForbidden).
			JSON(utils.InternalResponse{
				Success: false,
				Error:   "No valid session token could be found",
			})
	}

	c.Cookie(&fiber.Cookie{
		Name:    "csrf_token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})

	fmt.Println("TEST")
	c.Cookie(&fiber.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	})

	err = database.Client.Session.Delete(session.Id)
	if err != nil {
		return c.Status(http.StatusForbidden).
			JSON(utils.InternalResponse{
				Success: false,
				Error:   "No valid session token could be found",
			})
	}

	fmt.Println("TEST")
	return c.Status(http.StatusOK).
		JSON(utils.InternalResponse{
			Success: true,
			Message: "User has been signed out, and their session has been deleted",
		})
}
