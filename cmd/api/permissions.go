package main

import (
	"net/http"

	"github.com/Syha-01/national-inservice-training/internal/data"
	"github.com/Syha-01/national-inservice-training/internal/validator"
)

func (app *application) addUserPermissionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		Code string `json:"code"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	if data.ValidatePermissionCode(v, input.Code); !v.IsEmpty() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Permissions.AddForUser(id, input.Code)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "permission added successfully"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
