package api

import (
	"encoding/json"
	"errors"
	"github.com/hbeserra/ramengo/internal/broths"
	"github.com/hbeserra/ramengo/internal/orders"
	"github.com/hbeserra/ramengo/internal/proteins"
	"log/slog"
	"net/http"
)

// handleListBroths is a handler function that returns a list of broths
func HandleListBroths(w http.ResponseWriter, _ *http.Request) {
	WriteJson(w, http.StatusOK, broths.List())
}

// handleListProteins is a handler function that returns a list of proteins
func HandleListProteins(w http.ResponseWriter, _ *http.Request) {
	WriteJson(w, http.StatusOK, proteins.List())
}

type OrderRequest struct {
	BrothId   string `json:"brothId"`
	ProteinId string `json:"proteinId"`
}

// handleOrder is a handler function that handles an order
func HandleOrder(w http.ResponseWriter, r *http.Request) {

	// read request body
	var request OrderRequest
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		slog.Error("Failed to read request body", "error", err)
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "could not place order"})
		return
	}

	order, err := orders.Create(request.BrothId, request.ProteinId)
	if errors.Is(err, orders.ErrInvalidRequest) {
		WriteJson(w, http.StatusBadRequest, ErrorResponse{Error: "both brothId and proteinId are required"})
		return
	} else if err != nil {
		slog.Error("Failed to create order", "error", err)
		WriteJson(w, http.StatusInternalServerError, ErrorResponse{Error: "could not place order"})
		return
	}

	WriteJson(w, http.StatusOK, order)
}
