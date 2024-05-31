package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/hbeserra/ramengo/internal/api"
	"github.com/hbeserra/ramengo/internal/broths"
	"github.com/hbeserra/ramengo/internal/orders"
	"github.com/hbeserra/ramengo/internal/proteins"
)

// handleListBroths is a handler function that returns a list of broths
func handleListBroths(w http.ResponseWriter, _ *http.Request) {
	api.WriteJson(w, http.StatusOK, broths.List())
}

// handleListProteins is a handler function that returns a list of proteins
func handleListProteins(w http.ResponseWriter, _ *http.Request) {
	api.WriteJson(w, http.StatusOK, proteins.List())
}

type OrderRequest struct {
	BrothId   string `json:"brothId"`
	ProteinId string `json:"proteinId"`
}

// handleOrder is a handler function that handles an order
func handleOrder(w http.ResponseWriter, r *http.Request) {

	// read request body
	var request OrderRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		slog.Error("Failed to read request body", "error", err)
		api.WriteJson(w, http.StatusInternalServerError, api.ErrorResponse{Error: "could not place order"})
		return
	}

	order, err := orders.Create(request.BrothId, request.ProteinId)
	if errors.Is(err, orders.ErrInvalidRequest) {
		api.WriteJson(w, http.StatusBadRequest, api.ErrorResponse{Error: "both brothId and proteinId are required"})
		return
	} else if err != nil {
		slog.Error("Failed to create order", "error", err)
		api.WriteJson(w, http.StatusInternalServerError, api.ErrorResponse{Error: "could not place order"})
		return
	}

	api.WriteJson(w, http.StatusOK, order)
}

func main() {

	mux := http.NewServeMux()

	// add the middleware to every request
	wrappedMux := api.UseMiddleware(mux, api.CorsMiddleware, api.ValidateXAPIKeyMiddleware)

	mux.Handle("GET /broths", http.HandlerFunc(handleListBroths))
	mux.Handle("GET /proteins", http.HandlerFunc(handleListProteins))
	mux.Handle("POST /order", http.HandlerFunc(handleOrder))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	slog.Info("Starting server", "port", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), wrappedMux))
}
