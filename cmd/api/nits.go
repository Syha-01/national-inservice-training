package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Syha-01/national-inservice-training/internal/data"
	"github.com/Syha-01/national-inservice-training/internal/validator"
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

func (a *application) displayOfficerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	officer, err := a.models.Officers.Get(id)
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

	officer, err := a.models.Officers.Get(id)
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

	err = a.models.Officers.Update(officer)
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
