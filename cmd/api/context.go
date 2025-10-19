package main

import (
	"context"
	"net/http"

	"github.com/Syha-01/national-inservice-training/internal/data"
)

// Custom type for context keys to avoid collisions
type contextKey string

const userContextKey = contextKey("user")

// contextSetUser adds the user to the request context
func (a *application) contextSetUser(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

// contextGetUser retrieves the user from the request context
func (a *application) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}