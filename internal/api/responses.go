package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// ErrorResponse is a struct that represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// WriteJson returns a response using encoding/json
func WriteJson(w http.ResponseWriter, code int, obj interface{}) {
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(obj)
	if err != nil {
		slog.Error("Failed to encode response", "error", err)
	}

	w.Header().Set("Content-Type", "application/json")
}
