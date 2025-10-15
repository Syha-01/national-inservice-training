package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Syha-01/national-inservice-training/internal/data"
)

func (app *application) createFacilitatorFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		SessionEnrollmentID int64  `json:"session_enrollment_id"`
		Score               int    `json:"score"`
		Comment             string `json:"comment"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	feedback := &data.FacilitatorFeedback{
		FacilitatorID:       id,
		SessionEnrollmentID: input.SessionEnrollmentID,
		Score:               input.Score,
		Comment:             input.Comment,
	}

	err = app.models.Feedback.InsertFacilitatorFeedback(feedback)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/facilitators/%d/feedback/%d", feedback.FacilitatorID, feedback.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"feedback": feedback}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listFacilitatorFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	facilitator, err := app.models.Facilitators.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	if facilitator == nil {
		app.notFoundResponse(w, r)
		return
	}

	feedback, err := app.models.Feedback.GetAllForFacilitator(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"feedback": feedback}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createCourseFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	_, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		SessionEnrollmentID int64  `json:"session_enrollment_id"`
		Score               int    `json:"score"`
		Comment             string `json:"comment"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	feedback := &data.CourseFeedback{
		SessionEnrollmentID: input.SessionEnrollmentID,
		Score:               input.Score,
		Comment:             input.Comment,
	}

	err = app.models.Feedback.InsertCourseFeedback(feedback)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/courses/%d/feedback/%d", feedback.CourseID, feedback.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"feedback": feedback}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listCourseFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	course, err := app.models.Courses.GetCourse(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	if course == nil {
		app.notFoundResponse(w, r)
		return
	}

	feedback, err := app.models.Feedback.GetAllForCourse(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"feedback": feedback}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
