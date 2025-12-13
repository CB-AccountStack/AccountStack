package middleware

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

// contextKey is a custom type for context keys to avoid collisions
type contextKey string

const userIDKey contextKey = "userID"

// AuthMiddleware extracts the user ID from the X-User-ID header
// In a production system, this would validate JWT tokens or session cookies
func AuthMiddleware(logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// For demo purposes, we'll use a simple X-User-ID header
			// In production, this should validate JWT tokens or session cookies
			userID := r.Header.Get("X-User-ID")

			// If no user ID provided, default to demo user for easy testing
			if userID == "" {
				userID = "user-001" // Default demo user
			}

			// Add user ID to request context
			ctx := context.WithValue(r.Context(), userIDKey, userID)

			logger.WithField("userId", userID).Debug("User authenticated")

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserID extracts the user ID from the request context
func GetUserID(r *http.Request) string {
	userID, ok := r.Context().Value(userIDKey).(string)
	if !ok {
		return "user-001" // Default fallback
	}
	return userID
}
