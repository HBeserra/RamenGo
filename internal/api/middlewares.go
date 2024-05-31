package api

import (
	"log/slog"
	"net/http"
	"os"
)

// ValidateXAPIKeyMiddleware checks if the provided API key is valid
func ValidateXAPIKeyMiddleware(next http.Handler) http.Handler {

	slog.Info("ValidateXAPIKeyMiddleware")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		slog.Info("ValidateXAPIKeyMiddleware")

		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			WriteJson(w, http.StatusForbidden, ErrorResponse{Error: "x-api-key header missing"})
			return
		}

		// ToDo: RedVentures define a API key ?
		if apiKey != "ZtVdh8XQ2U8pWI2gmZ7f796Vh8GllXoN7mr0djNf" {
			WriteJson(w, http.StatusForbidden, ErrorResponse{Error: "x-api-key header invalid"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

// CorsMiddleware adds CORS headers to the response
func CorsMiddleware(next http.Handler) http.Handler {

	allowOrigin := os.Getenv("ALLOWED_ORIGIN")
	if allowOrigin == "" {
		allowOrigin = "*"
	}

	slog.Info("CorsMiddleware", "allowOrigins", allowOrigin)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		slog.Info("CorsMiddleware")

		slog.Info("Request", "req.host", r.Host)

		w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, x-api-key")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// UseMiddleware applies a series of middleware functions to the given http.ServeMux.
//
// Parameters:
// - r: The http.ServeMux to which the middleware functions will be applied.
// - middlewares: A variadic parameter of middleware functions.
//
// Return type:
// - http.Handler: The resulting http.Handler after applying all the middleware functions.
func UseMiddleware(r *http.ServeMux, middlewares ...func(http.Handler) http.Handler) http.Handler {
	var mwHandler http.Handler = r

	for i := len(middlewares) - 1; i >= 0; i-- {
		mwHandler = middlewares[i](mwHandler)
	}

	return mwHandler
}
