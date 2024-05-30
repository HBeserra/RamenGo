package main

import (
	"log"
	"net/http"

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
	wrappedMux := api.UseMiddleware(mux, api.ValidateXAPIKeyMiddleware)

	log.Fatal(http.ListenAndServe(":3000", wrappedMux))
}
