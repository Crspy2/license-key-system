package http

import (
	"fmt"
	authRoutes "github.com/crspy2/license-panel/app/http/endpoints/auth"
	selfRoutes "github.com/crspy2/license-panel/app/http/endpoints/self"
	"github.com/crspy2/license-panel/app/http/middleware"
	"github.com/crspy2/license-panel/app/http/utils"
	"github.com/crspy2/license-panel/config"
	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/storage/redis/v2"
	"log"
	"net/http"
	"time"
)

func StartServer() {
	// TODO: Add cookie encryption

	app := fiber.New()

	fmt.Println("Connecting to Redis instance...")
	redisClient := redis.New(redis.Config{
		Host:     config.Conf.Redis.Host,
		Port:     config.Conf.Redis.Port,
		Password: config.Conf.Redis.Password,
		PoolSize: config.Conf.Redis.Threads,
	})
	fmt.Println("Connected to Redis")

	// Sentry configuration
	app.Use(fibersentry.New(fibersentry.Config{
		Repanic:         true,
		WaitForDelivery: true,
	}))

	// Encrypt all cookies except csrf_token
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key:    config.Conf.SessionEncryptionKey,
		Except: []string{"csrf_token"},
	}))

	app.Use(recover.New())
	app.Get("/metrics", monitor.New())
	app.Use(logger.New())
	app.Use(helmet.New())

	// Set Anti-CSRF token protection
	app.Use(csrf.New(csrf.Config{
		KeyLookup:         "header:" + csrf.HeaderName,
		CookieName:        "csrf_token",
		CookieSameSite:    "Lax",
		CookieSecure:      true,
		CookieSessionOnly: true,
		CookieHTTPOnly:    false,
		Expiration:        1 * time.Hour,
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			return c.Status(http.StatusForbidden).JSON(utils.InternalResponse{
				Success: false,
				Error:   "Forbidden — Invalid CSRF token provided",
			})
		},
		Storage:    redisClient,
		Extractor:  utils.GetCSRFFromHeader,
		SessionKey: "session_token",
	}))

	// Configure ratelimit settings
	app.Use(limiter.New(limiter.Config{
		Max:               config.Conf.Ratelimit.Max,
		Expiration:        time.Duration(config.Conf.Ratelimit.Window) * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
		LimitReached: func(c *fiber.Ctx) error {
			retryAfter := c.GetRespHeader("Retry-After")
			return c.Status(http.StatusTooManyRequests).JSON(utils.InternalResponse{
				Success: false,
				Error:   fmt.Sprintf("You have been ratelimited, please try in %s", retryAfter),
			})
		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello!")
	})

	app.Get("/clear-cookies", func(c *fiber.Ctx) error {
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

		return c.JSON(fiber.Map{
			"success": "cookies have been reset",
		})
	})

	auth := app.Group("/auth")
	auth.Post("/register", authRoutes.RegisterRoute)
	auth.Post("/login", authRoutes.LoginRoute)
	auth.Post("/logout", middleware.AuthenticateCookie, authRoutes.LogoutRoute)

	self := app.Group("/self", middleware.AuthenticateCookie)

	self.Get("/", selfRoutes.SelfRoute)

	log.Fatal(app.Listen(":8080"))
}
