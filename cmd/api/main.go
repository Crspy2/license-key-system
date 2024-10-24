package main

import (
	"fmt"
	"github.com/crspy2/license-panel/app/http"
	"github.com/crspy2/license-panel/config"
	"github.com/crspy2/license-panel/database"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	fmt.Println("Loading environment variables...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.LoadConfig()

	//Initialize Sentry
	fmt.Println("Initializing Sentry")
	_ = sentry.Init(sentry.ClientOptions{
		Dsn: config.Conf.SentryDSN,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint.Context != nil {
				if c, ok := hint.Context.Value(sentry.RequestContextKey).(*fiber.Ctx); ok {
					// You have access to the original Context if it panicked
					fmt.Println(utils.CopyString(c.Hostname()))
				}
			}
			fmt.Println(event)
			return event
		},
		Debug:            true,
		AttachStacktrace: true,
	})

	fmt.Println("Connecting to database...")
	database.ConnectToDatabase()

	http.StartServer()
}
