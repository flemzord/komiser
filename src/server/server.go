package server

import (
	"embed"
	"fmt"
	. "github.com/flemzord/komiser/src/handlers/aws"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed dashboard/*
var distFiles embed.FS

func clientHandler() http.Handler {
	fsys := fs.FS(distFiles)
	contentStatic, _ := fs.Sub(fsys, "dashboard")
	return http.FileServer(http.FS(contentStatic))

}

func Server(cmd *cobra.Command, args []string) {
	port := 3010

	r := mux.NewRouter()

	awsHandler := NewAWSHandler()

	r.HandleFunc("/aws/cost/current", awsHandler.CurrentCostHandler)
	r.HandleFunc("/aws/cost/history", awsHandler.CostAndUsageHandler)
	r.HandleFunc("/aws/support/history", awsHandler.SupportTicketsInLastSixMonthsHandlers)
	r.HandleFunc("/aws/support/open", awsHandler.SupportOpenTicketsHandler)

	allowedHeaders := handlers.AllowedHeaders([]string{})
	loggedRouter := handlers.LoggingHandler(os.Stdout, handlers.CORS(allowedHeaders)(r))
	r.Handle("/", clientHandler())

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), loggedRouter)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Server started on port %d", port)
	}
}
