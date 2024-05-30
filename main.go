package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/hbeserra/ramengo/internal/api"
	"github.com/hbeserra/ramengo/internal/broths"
)

// handleListBroths is a handler function that returns a list of broths
func handleListBroths(w http.ResponseWriter, r *http.Request) {
	api.WriteJson(w, http.StatusOK, broths.List())
}

func main() {

	mux := http.NewServeMux()

	mux.Handle("GET /broths", http.HandlerFunc(handleListBroths))

	// add the middleware to every request
	wrappedMux := api.UseMiddleware(mux, api.CorsMiddleware, api.ValidateXAPIKeyMiddleware)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	slog.Info("Starting server", "port", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), wrappedMux))
}
