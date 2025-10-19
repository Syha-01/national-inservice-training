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
	router.HandlerFunc(http.MethodPost, "/v1/nits", app.createNitHandler)
	router.HandlerFunc(http.MethodPost, "/v1/officers", app.createOfficerHandler)
	router.HandlerFunc(http.MethodGet, "/v1/officers/:id", app.displayOfficerHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/officers/:id", app.updateOfficerHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/officers/:id", app.deleteOfficerHandler)
	router.HandlerFunc(http.MethodGet, "/v1/officers", app.listOfficersHandler)

	// facilitator routes
	router.HandlerFunc(http.MethodPost, "/v1/facilitators", app.createFacilitatorHandler)
	router.HandlerFunc(http.MethodGet, "/v1/facilitators", app.listFacilitatorsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/facilitators/:id", app.showFacilitatorHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/facilitators/:id", app.updateFacilitatorHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/facilitators/:id", app.deleteFacilitatorHandler)
	router.HandlerFunc(http.MethodPost, "/v1/sessions/:id/facilitators", app.assignFacilitatorToSessionHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/sessions/:id/facilitators/:facilitator_id", app.removeFacilitatorFromSessionHandler)

	// Facilitator feedback routes
	router.HandlerFunc(http.MethodPost, "/v1/facilitators/:id/feedback", app.createFacilitatorFeedbackHandler)
	router.HandlerFunc(http.MethodGet, "/v1/facilitators/:id/feedback", app.listFacilitatorFeedbackHandler)

	// COURSES ROUTES
	router.HandlerFunc(http.MethodGet, "/v1/courses", app.listCoursesHandler)
	router.HandlerFunc(http.MethodPost, "/v1/courses", app.createCourseHandler)
	router.Handle(http.MethodGet, "/v1/courses/:id", httprouter.Handle(app.showCourseHandler))
	router.Handle(http.MethodPatch, "/v1/courses/:id", httprouter.Handle(app.updateCourseHandler))
	router.Handle(http.MethodDelete, "/v1/courses/:id", httprouter.Handle(app.deleteCourseHandler))

	// Course feedback routes
	router.HandlerFunc(http.MethodPost, "/v1/courses/:id/feedback", app.createCourseFeedbackHandler)
	router.HandlerFunc(http.MethodGet, "/v1/courses/:id/feedback", app.listCourseFeedbackHandler)

	// User routes
	// router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/tokens", app.createAuthenticationTokenHandler)

	// Ratings and feedback routes
	router.HandlerFunc(http.MethodPost, "/v1/enrollments/:id/courserating", app.createCourseRatingHandler)
	router.HandlerFunc(http.MethodPost, "/v1/enrollments/:id/facilitatorrating", app.createFacilitatorRatingHandler)

	// TODO: Add more routes as you build your API
	// router.HandlerFunc(http.MethodGet, "/v1/officers", app.listOfficersHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/officers", app.createOfficerHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/officers/:id", app.showOfficerHandler)
	// router.HandlerFunc(http.MethodPatch, "/v1/officers/:id", app.updateOfficerHandler)
	// router.HandlerFunc(http.MethodDelete, "/v1/officers/:id", app.deleteOfficerHandler)

	// router.HandlerFunc(http.MethodGet, "/v1/trainings", app.listTrainingsHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/trainings", app.createTrainingHandler)

	//Middleware chain
	// 1. recoverPanic - catches any panics and returns 500
	// 2. enableCORS - adds CORS headers and handles preflight
	// 3. rateLimit - limits request rate (add when implemented)
	// 4. authenticate - checks authentication (add when implemented)
	return app.recoverPanic(app.rateLimit(app.enableCORS(router)))
}
