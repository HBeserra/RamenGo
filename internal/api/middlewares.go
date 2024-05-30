package api

import "net/http"

// ValidateXAPIKeyMiddleware checks if the provided API key is valid
func ValidateXAPIKeyMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			WriteJson(w, http.StatusForbidden, ErrorResponse{Error: "x-api-key header missing"})
			return
		}

		// ToDo: RedVentures define a API key ?
		if apiKey != "12345" {
			WriteJson(w, http.StatusForbidden, ErrorResponse{Error: "x-api-key header invalid"})
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

	for _, mw := range middlewares {
		mwHandler = mw(mwHandler)
	}

	return mwHandler
}
