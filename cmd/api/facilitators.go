package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Syha-01/national-inservice-training/internal/data"
	"github.com/Syha-01/national-inservice-training/internal/validator"
)

func (a *application) createFacilitatorHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		PersonnelID *int64 `json:"personnel_id"`
	}

	err := a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	facilitator := &data.Facilitator{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}

	if input.PersonnelID != nil {
		v := validator.New()
		if _, err := a.models.Officers.GetOfficer(*input.PersonnelID); err != nil {
			v.AddError("personnel_id", "personnel_id does not exist")
			a.failedValidationResponse(w, r, v.Errors)
			return
		}

		if _, err := a.models.Facilitators.GetByPersonnelID(*input.PersonnelID); err == nil {
			v.AddError("personnel_id", "a facilitator with this personnel_id already exists")
			a.failedValidationResponse(w, r, v.Errors)
			return
		}

		facilitator.PersonnelID.Int64 = *input.PersonnelID
		facilitator.PersonnelID.Valid = true
	}

	err = a.models.Facilitators.Create(facilitator)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/facilitators/%d", facilitator.ID))

	err = a.writeJSON(w, http.StatusCreated, envelope{"facilitator": facilitator}, headers)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) listFacilitatorsHandler(w http.ResponseWriter, r *http.Request) {
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

	facilitators, metadata, err := a.models.Facilitators.GetAll(input.Filters)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"facilitators": facilitators, "metadata": metadata}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) showFacilitatorHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	facilitator, err := a.models.Facilitators.Get(id)
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
		"facilitator": facilitator,
	}

	err = a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}

func (a *application) updateFacilitatorHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	facilitator, err := a.models.Facilitators.Get(id)
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
		FirstName   *string `json:"first_name"`
		LastName    *string `json:"last_name"`
		Email       *string `json:"email"`
		PersonnelID *int64  `json:"personnel_id"`
	}

	err = a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	if input.FirstName != nil {
		facilitator.FirstName = *input.FirstName
	}
	if input.LastName != nil {
		facilitator.LastName = *input.LastName
	}
	if input.Email != nil {
		facilitator.Email = *input.Email
	}
	if input.PersonnelID != nil {
		v := validator.New()
		if _, err := a.models.Officers.GetOfficer(*input.PersonnelID); err != nil {
			v.AddError("personnel_id", "personnel_id does not exist")
			a.failedValidationResponse(w, r, v.Errors)
			return
		}
		facilitator.PersonnelID.Int64 = *input.PersonnelID
		facilitator.PersonnelID.Valid = true
	}

	err = a.models.Facilitators.Update(facilitator)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			a.editConflictResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"facilitator": facilitator}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) deleteFacilitatorHandler(w http.ResponseWriter, r *http.Request) {
	id, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	err = a.models.Facilitators.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"message": "facilitator successfully deleted"}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) assignFacilitatorToSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	var input struct {
		FacilitatorID int64 `json:"facilitator_id"`
	}

	err = a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	_, err = a.models.Nits.Get(sessionID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	_, err = a.models.Facilitators.Get(input.FacilitatorID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	err = a.models.Facilitators.AssignToSession(sessionID, input.FacilitatorID)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	err = a.writeJSON(w, http.StatusCreated, envelope{"message": "facilitator assigned to session successfully"}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) listFacilitatorsForSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

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

	facilitators, metadata, err := a.models.Facilitators.GetAllForSession(sessionID, input.Filters)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"facilitators": facilitators, "metadata": metadata}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) removeFacilitatorFromSessionHandler(w http.ResponseWriter, r *http.Request) {
	sessionID, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	facilitatorID, err := a.readIDParam(r)
	if err != nil {
		a.notFoundResponse(w, r)
		return
	}

	_, err = a.models.Nits.Get(sessionID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	_, err = a.models.Facilitators.Get(facilitatorID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	err = a.models.Facilitators.RemoveFromSession(sessionID, facilitatorID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.notFoundResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"message": "facilitator removed from session successfully"}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}
