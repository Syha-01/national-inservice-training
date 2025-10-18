package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Syha-01/national-inservice-training/internal/data"
	"github.com/Syha-01/national-inservice-training/internal/validator"
	"github.com/julienschmidt/httprouter"
)

func (a *application) createNitHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		CourseID  int64     `json:"course_id"`
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		Location  string    `json:"location"`
	}

	err := a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	nit := &data.Nit{
		CourseID:  input.CourseID,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
		Location:  input.Location,
	}

	v := validator.New()

	data.ValidateNit(v, nit)

	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", nit)
}

func (a *application) createOfficerHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		RegulationNumber string `json:"regulation_number"`
		FirstName        string `json:"first_name"`
		LastName         string `json:"last_name"`
		Sex              string `json:"sex"`
		RankID           int64  `json:"rank_id"`
		FormationID      int64  `json:"formation_id"`
		PostingID        int64  `json:"posting_id"`
		IsActive         bool   `json:"is_active"`
	}

	err := a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	officer := &data.Officer{
		RegulationNumber: input.RegulationNumber,
		FirstName:        input.FirstName,
		LastName:         input.LastName,
		Sex:              input.Sex,
		RankID:           input.RankID,
		FormationID:      input.FormationID,
		PostingID:        input.PostingID,
		IsActive:         input.IsActive,
	}

	v := validator.New()

	if data.ValidateOfficer(v, officer); !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = a.models.Officers.CreateOfficer(officer)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/officers/%d", officer.ID))

	err = a.writeJSON(w, http.StatusCreated, envelope{"officer": officer}, headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) displayOfficerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	officer, err := a.models.Officers.GetOfficer(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	data := envelope{
		"officer": officer,
	}

	err = a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}

func (a *application) updateOfficerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	officer, err := a.models.Officers.GetOfficer(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		RegulationNumber *string `json:"regulation_number"`
		FirstName        *string `json:"first_name"`
		LastName         *string `json:"last_name"`
		Sex              *string `json:"sex"`
		RankID           *int64  `json:"rank_id"`
		FormationID      *int64  `json:"formation_id"`
		PostingID        *int64  `json:"posting_id"`
		IsActive         *bool   `json:"is_active"`
	}

	err = a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	if input.RegulationNumber != nil {
		officer.RegulationNumber = *input.RegulationNumber
	}
	if input.FirstName != nil {
		officer.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		officer.LastName = *input.LastName
	}
	if input.Sex != nil {
		officer.Sex = *input.Sex
	}
	if input.RankID != nil {
		officer.RankID = *input.RankID
	}
	if input.FormationID != nil {
		officer.FormationID = *input.FormationID
	}
	if input.PostingID != nil {
		officer.PostingID = *input.PostingID
	}
	if input.IsActive != nil {
		officer.IsActive = *input.IsActive
	}

	v := validator.New()

	if data.ValidateOfficer(v, officer); !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = a.models.Officers.UpdateOfficer(officer)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			a.editConflictResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"officer": officer}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) deleteOfficerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	err = a.models.Officers.DeleteOfficer(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"message": "officer successfully deleted"}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) listOfficersHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		data.Filters
	}

	v := validator.New()
	qs := r.URL.Query()

	input.Filters.Page = a.readInt(qs, "page", 1, v)
	input.Filters.PageSize = a.readInt(qs, "page_size", 20, v)

	data.ValidateFilters(v, input.Filters)

	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	officers, metadata, err := a.models.Officers.GetAllOfficers(input.Filters)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"officers": officers, "metadata": metadata}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

// listCoursesHandler returns all courses
func (a *application) listCoursesHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		data.Filters
	}

	v := validator.New()
	qs := r.URL.Query()

	input.Filters.Page = a.readInt(qs, "page", 1, v)
	input.Filters.PageSize = a.readInt(qs, "page_size", 20, v)

	data.ValidateFilters(v, input.Filters)

	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	courses, metadata, err := a.models.Courses.GetAllCourses(input.Filters)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"courses": courses, "metadata": metadata}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

// showCourseHandler returns a single course by ID
func (app *application) showCourseHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		app.badRequestResponse(w, r, errors.New("invalid course ID"))
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

	err = app.writeJSON(w, http.StatusOK, envelope{"course": course}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// createCourseHandler creates a new course
func (app *application) createCourseHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Category    string  `json:"category"`
		CreditHours float64 `json:"credit_hours"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	course := &data.Course{
		Title:       input.Title,
		Description: input.Description,
		Category:    input.Category,
		CreditHours: input.CreditHours,
	}

	v := validator.New()
	if data.ValidateCourse(v, course); !v.IsEmpty() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Courses.CreateCourse(course)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", "/v1/courses/"+strconv.FormatInt(course.ID, 10))

	err = app.writeJSON(w, http.StatusCreated, envelope{"course": course}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// updateCourseHandler updates an existing course
func (app *application) updateCourseHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		app.badRequestResponse(w, r, errors.New("invalid course ID"))
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

	var input struct {
		Title       *string  `json:"title"`
		Description *string  `json:"description"`
		Category    *string  `json:"category"`
		CreditHours *float64 `json:"credit_hours"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Update only provided fields
	if input.Title != nil {
		course.Title = *input.Title
	}
	if input.Description != nil {
		course.Description = *input.Description
	}
	if input.Category != nil {
		course.Category = *input.Category
	}
	if input.CreditHours != nil {
		course.CreditHours = *input.CreditHours
	}

	v := validator.New()
	if data.ValidateCourse(v, course); !v.IsEmpty() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Courses.UpdateCourse(course)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"course": course}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// deleteCourseHandler deletes a course
func (app *application) deleteCourseHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		app.badRequestResponse(w, r, errors.New("invalid course ID"))
		return
	}

	err = app.models.Courses.DeleteCourse(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "course successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
