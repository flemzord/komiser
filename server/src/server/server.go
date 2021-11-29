package server

import (
	"log"

	. "github.com/flemzord/komiser/src/handlers/aws"
	"github.com/flemzord/komiser/src/libs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/cobra"
)

func Server(cmd *cobra.Command, args []string) {
	port := ":" + libs.GetEnvVar("PORT", "3010")

	awsHandler := NewAWSHandler()

	app := fiber.New()
	app.Use(cors.New())
	if libs.GetEnvVar("APP_ENV", "local") != "local" {
		app.Use(cache.New())
	}

	app.Static("/", "./../dashboard/dist")

	aws := app.Group("/aws")

	cost := aws.Group("/cost")
	cost.Get("/current", awsHandler.CurrentCostHandler)
	cost.Get("/history", awsHandler.CostAndUsageHandler)

	support := aws.Group("/support")
	support.Get("history", awsHandler.SupportTicketsInLastSixMonthsHandlers)
	support.Get("open", awsHandler.SupportOpenTicketsHandler)

	log.Fatal(app.Listen(port))

	//r.HandleFunc("/aws/support/history", awsHandler.SupportTicketsInLastSixMonthsHandlers)
	//r.HandleFunc("/aws/support/open", awsHandler.SupportOpenTicketsHandler)
}
