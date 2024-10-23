package auth

import (
	"github.com/crspy2/license-panel/database"
	"github.com/gofiber/fiber/v2"
	"go.jetify.com/typeid"
	"net/http"
	"time"
)

func LoginRoute(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	staff, err := database.Client.Staff.Authenticate(username, password)
	if err != nil {
		return c.Status(http.StatusNotFound).
			JSON(fiber.Map{
				"status": http.StatusNotFound,
				"error":  "No user found with the specified username",
			})
	}

	sessionToken := typeid.Must(typeid.WithPrefix("sess")).String()

	sessionInfo := database.SessionModal{
		Id:        sessionToken,
		StaffId:   staff.Id,
		IpAddress: c.IP(),
		UserAgent: c.Get("User-Agent"),
		ExpiresAt: time.Now().Add(5 * time.Hour),
	}

	_ = database.Client.Session.DeleteByIP(c.IP())
	err = database.Client.Session.Create(&sessionInfo)
	if err != nil {
		return c.Status(http.StatusInternalServerError).
			JSON(fiber.Map{
				"status": http.StatusInternalServerError,
				"error":  err.Error(),
			})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  sessionInfo.ExpiresAt,
		Secure:   true,
		HTTPOnly: true,
	})

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": "User session created",
			"data":    sessionInfo,
		})
}
