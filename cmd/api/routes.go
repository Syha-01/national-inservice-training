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

	// nit routes
	router.HandlerFunc(http.MethodPost, "/v1/nits", app.requirePermission("nits:write", app.createNitHandler))
	router.HandlerFunc(http.MethodGet, "/v1/nits/:id", app.requirePermission("nits:read", app.showNitHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/nits/:id", app.requirePermission("nits:write", app.updateNitHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/nits/:id", app.requirePermission("nits:write", app.deleteNitHandler))
	router.HandlerFunc(http.MethodGet, "/v1/nits", app.requirePermission("nits:read", app.listNitsHandler))

	router.HandlerFunc(http.MethodPost, "/v1/officers", app.requirePermission("officers:write", app.createOfficerHandler))
	router.HandlerFunc(http.MethodGet, "/v1/officers/:id", app.requirePermission("officers:read", app.displayOfficerHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/officers/:id", app.requirePermission("officers:write", app.updateOfficerHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/officers/:id", app.requirePermission("officers:write", app.deleteOfficerHandler))
	router.HandlerFunc(http.MethodGet, "/v1/officers", app.requirePermission("officers:read", app.listOfficersHandler))

	// facilitator routes
	router.HandlerFunc(http.MethodPost, "/v1/facilitators", app.requirePermission("facilitators:write", app.createFacilitatorHandler))
	router.HandlerFunc(http.MethodGet, "/v1/facilitators", app.requirePermission("facilitators:read", app.listFacilitatorsHandler))
	router.HandlerFunc(http.MethodGet, "/v1/facilitators/:id", app.requirePermission("facilitators:read", app.showFacilitatorHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/facilitators/:id", app.requirePermission("facilitators:write", app.updateFacilitatorHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/facilitators/:id", app.requirePermission("facilitators:write", app.deleteFacilitatorHandler))
	router.HandlerFunc(http.MethodPost, "/v1/sessions/:id/facilitators", app.requirePermission("facilitators:write", app.assignFacilitatorToSessionHandler))
	router.HandlerFunc(http.MethodGet, "/v1/sessions/:id/facilitators", app.requirePermission("facilitators:read", app.listFacilitatorsForSessionHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/sessions/:id/facilitators/:facilitator_id", app.requirePermission("facilitators:write", app.removeFacilitatorFromSessionHandler))

	// Facilitator feedback routes
	router.HandlerFunc(http.MethodPost, "/v1/facilitators/:id/feedback", app.requirePermission("feedback:write", app.createFacilitatorFeedbackHandler))
	router.HandlerFunc(http.MethodGet, "/v1/facilitators/:id/feedback", app.requirePermission("feedback:read", app.listFacilitatorFeedbackHandler))

	// COURSES ROUTES
	router.HandlerFunc(http.MethodGet, "/v1/courses", app.requirePermission("courses:read", app.listCoursesHandler))
	router.HandlerFunc(http.MethodPost, "/v1/courses", app.requirePermission("courses:write", app.createCourseHandler))
	router.HandlerFunc(http.MethodGet, "/v1/courses/:id", app.requirePermission("courses:read", app.showCourseHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/courses/:id", app.requirePermission("courses:write", app.updateCourseHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/courses/:id", app.requirePermission("courses:write", app.deleteCourseHandler))
	// Course feedback routes
	router.HandlerFunc(http.MethodPost, "/v1/courses/:id/feedback", app.requirePermission("feedback:write", app.createCourseFeedbackHandler))
	router.HandlerFunc(http.MethodGet, "/v1/courses/:id/feedback", app.requirePermission("feedback:read", app.listCourseFeedbackHandler))

	// User routes
	// router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/tokens", app.createAuthenticationTokenHandler)

	// Ratings and feedback routes
	router.HandlerFunc(http.MethodPost, "/v1/enrollments/:id/courserating", app.requireActivatedUser(app.createCourseRatingHandler))
	router.HandlerFunc(http.MethodPost, "/v1/enrollments/:id/facilitatorrating", app.requireActivatedUser(app.createFacilitatorRatingHandler))

	// TODO: Add more routes as you build your API
	// router.HandlerFunc(http.MethodGet, "/v1/officers", app.listOfficersHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/officers", app.createOfficerHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/officers/:id", app.showOfficerHandler)
	// router.HandlerFunc(http.MethodPatch, "/v1/officers/:id", app.updateOfficerHandler)
	// router.HandlerFunc(http.MethodDelete, "/v1/officers/:id", app.deleteOfficerHandler)

	// router.HandlerFunc(http.MethodGet, "/v1/trainings", app.listTrainingsHandler)
	// router.HandlerFunc(http.MethodPost, "/v1/trainings", app.createTrainingHandler)

	// User authentication routes (public)
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	// Permissions routes
	router.HandlerFunc(http.MethodPost, "/v1/users/:id/permissions", app.requirePermission("admin:all", app.addUserPermissionHandler))

	//Middleware chain
	// 1. recoverPanic - catches any panics and returns 500
	// 2. enableCORS - adds CORS headers and handles preflight
	// 3. rateLimit - limits request rate (add when implemented)
	// 4. authenticate - checks authentication (add when implemented)
	return app.recoverPanic(app.rateLimit(app.authenticate(app.enableCORS(router))))
}
