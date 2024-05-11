package middleware

import (
	"net/http"
)

// CorsMiddleware is a middleware function that adds CORS headers to the response.
func CorsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Set the CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5173")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE ")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle preflight requests (OPTIONS)
        if r.Method == http.MethodOptions {
            // Respond with the CORS headers and a 204 status (No Content)
            w.WriteHeader(http.StatusNoContent)
            return
        }

        // Otherwise, continue to the next handler
        next.ServeHTTP(w, r)
    })
}
