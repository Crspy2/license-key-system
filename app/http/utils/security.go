package utils

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func GetCSRFFromHeader(c *fiber.Ctx) (string, error) {
	headers := c.GetReqHeaders()
	csrfToken := headers["X-Csrf-Token"]
	if csrfToken[0] == "" {
		return "", errors.New("invalid csrf token")
	}

	return csrfToken[0], nil
}
