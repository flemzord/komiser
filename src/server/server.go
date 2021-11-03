package server

import (
	"embed"
	. "github.com/flemzord/komiser/src/handlers/aws"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"io/fs"
	"log"
	"net/http"
)

//go:embed dashboard/*
var distFiles embed.FS

func clientHandler() http.Handler {
	fsys := fs.FS(distFiles)
	contentStatic, _ := fs.Sub(fsys, "dashboard")
	return http.FileServer(http.FS(contentStatic))
}

func Server(cmd *cobra.Command, args []string) {
	port := ":3010"

	awsHandler := NewAWSHandler()

	app := fiber.New()
	app.Static("/", "./dashboard")

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
