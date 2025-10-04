package main

import (
	"fmt"
	"net/http"
)

// recoverPanic middleware recovers from panics and returns a 500 error
func (a *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// defer will be called when the stack unwinds
		defer func() {
			// recover() checks for panics
			err := recover()
			if err != nil {
				w.Header().Set("Connection", "close")
				a.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// enableCORS adds CORS headers to responses and handles preflight requests
func (a *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add Vary header to indicate response varies by Origin
		w.Header().Add("Vary", "Origin")
		w.Header().Add("Vary", "Access-Control-Request-Method")

		// Get the origin from the request
		origin := r.Header.Get("Origin")

		// Check if we have trusted origins configured
		if len(a.config.cors.trustedOrigins) != 0 {
			// Check if the origin is in our trusted list
			originAllowed := false
			for i := range a.config.cors.trustedOrigins {
				if origin == a.config.cors.trustedOrigins[i] {
					originAllowed = true
					break
				}
			}

			if originAllowed {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
		} else {
			// Development mode: allow all origins
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Max-Age", "86400") // Cache for 24 hours

			// Return 200 OK for preflight
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// rateLimit middleware (placeholder - implement based on your requirements)
func (a *application) rateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement rate limiting logic using golang.org/x/time/rate
		// For now, just pass through
		next.ServeHTTP(w, r)
	})
}

// authenticate middleware (placeholder - implement based on your requirements)
func (a *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement JWT or session-based authentication
		// Extract and validate token from Authorization header
		// For now, just pass through
		next.ServeHTTP(w, r)
	})
}

// requireAuthenticatedUser ensures user is authenticated
func (a *application) requireAuthenticatedUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Check if user is authenticated
		// If not, return 401 Unauthorized
		next.ServeHTTP(w, r)
	}
}

// requirePermission checks if user has specific permission
func (a *application) requirePermission(permission string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Check user permissions based on role
		// Administrator, Content Contributor, or System User
		next.ServeHTTP(w, r)
	}
}