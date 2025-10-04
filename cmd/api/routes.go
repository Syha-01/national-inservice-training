package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// setup a new router
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// setup routes
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	
	// TODO: Add more routes as you build your API
	// router.HandlerFunc(http.MethodGet, "/v1/officers", app.listOfficersHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/officers", app.createOfficerHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/officers/:id", app.showOfficerHandler)
	// router.HandlerFunc(http.MethodPatch, "/v1/officers/:id", app.updateOfficerHandler)
	// router.HandlerFunc(http.MethodDelete, "/v1/officers/:id", app.deleteOfficerHandler)
	
	// router.HandlerFunc(http.MethodGet, "/v1/trainings", app.listTrainingsHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/trainings", app.createTrainingHandler)
	
	// Middleware chain (order matters!):
	// 1. recoverPanic - catches any panics and returns 500
	// 2. enableCORS - adds CORS headers and handles preflight
	// 3. rateLimit - limits request rate (add when implemented)
	// 4. authenticate - checks authentication (add when implemented)
	return app.recoverPanic(app.enableCORS(router))
}