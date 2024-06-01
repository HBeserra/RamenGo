package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/hbeserra/ramengo/internal/api"
)

func main() {

	mux := http.NewServeMux()

	// add the middleware to every request
	wrappedMux := api.UseMiddleware(mux, api.CorsMiddleware, api.ValidateXAPIKeyMiddleware)

	mux.Handle("GET /broths", http.HandlerFunc(api.HandleListBroths))
	mux.Handle("GET /proteins", http.HandlerFunc(api.HandleListProteins))
	mux.Handle("POST /order", http.HandlerFunc(api.HandleOrder))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	slog.Info("Starting server", "port", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), wrappedMux))
}
