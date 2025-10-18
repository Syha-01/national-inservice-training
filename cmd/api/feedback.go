package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Syha-01/national-inservice-training/internal/data"
)

func (a *application) createFacilitatorFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	var input struct {
		SessionEnrollmentID int64  `json:"session_enrollment_id"`
		Score               int    `json:"score"`
		Comment             string `json:"comment"`
	}

	err = a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	feedback := &data.FacilitatorFeedback{
		FacilitatorID:       id,
		SessionEnrollmentID: input.SessionEnrollmentID,
		Score:               input.Score,
		Comment:             input.Comment,
	}

	err = a.models.Feedback.InsertFacilitatorFeedback(feedback)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/facilitators/%d/feedback/%d", feedback.FacilitatorID, feedback.ID))

	err = a.writeJSON(w, http.StatusCreated, envelope{"feedback": feedback}, headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) listFacilitatorFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	feedback, err := a.models.Feedback.GetAllForFacilitator(id)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"feedback": feedback}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) createCourseRatingHandler(w http.ResponseWriter, r *http.Request) {
	enrollmentID, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	var input struct {
		Score   int    `json:"score"`
		Comment string `json:"comment"`
	}

	err = a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	feedback := &data.CourseFeedback{
		SessionEnrollmentID: enrollmentID,
		Score:               input.Score,
		Comment:             input.Comment,
	}

	err = a.models.Feedback.InsertCourseFeedback(feedback)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/courses/%d/feedback/%d", feedback.CourseID, feedback.ID))

	err = a.writeJSON(w, http.StatusCreated, envelope{"feedback": feedback}, headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) createFacilitatorRatingHandler(w http.ResponseWriter, r *http.Request) {
	enrollmentID, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	var input struct {
		FacilitatorID int64  `json:"facilitator_id"`
		Score         int    `json:"score"`
		Comment       string `json:"comment"`
	}

	err = a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	feedback := &data.FacilitatorFeedback{
		FacilitatorID:       input.FacilitatorID,
		SessionEnrollmentID: enrollmentID,
		Score:               input.Score,
		Comment:             input.Comment,
	}

	err = a.models.Feedback.InsertFacilitatorFeedback(feedback)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/facilitators/%d/feedback/%d", feedback.FacilitatorID, feedback.ID))

	err = a.writeJSON(w, http.StatusCreated, envelope{"feedback": feedback}, headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) createCourseFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	_, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	var input struct {
		SessionEnrollmentID int64  `json:"session_enrollment_id"`
		Score               int    `json:"score"`
		Comment             string `json:"comment"`
	}

	err = a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	feedback := &data.CourseFeedback{
		SessionEnrollmentID: input.SessionEnrollmentID,
		Score:               input.Score,
		Comment:             input.Comment,
	}

	err = a.models.Feedback.InsertCourseFeedback(feedback)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/courses/%d/feedback/%d", feedback.CourseID, feedback.ID))

	err = a.writeJSON(w, http.StatusCreated, envelope{"feedback": feedback}, headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) listCourseFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	course, err := a.models.Courses.GetCourse(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}
	if course == nil {
		a.notFoundResponse(w, r)
		return
	}

	feedback, err := a.models.Feedback.GetAllForCourse(id)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"feedback": feedback}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}
