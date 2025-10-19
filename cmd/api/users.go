package main

import (
	"errors"
	"time"
		"database/sql"
	"github.com/Syha-01/national-inservice-training/internal/data"
	"github.com/Syha-01/national-inservice-training/internal/validator"
	"net/http"
)

func (a *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email       string        `json:"email"`
		Password    string        `json:"password"`
		PersonnelID sql.NullInt64 `json:"personnel_id"`
		RoleID      int           `json:"role_id"` // Add this - allow client to specify role
	}

	err := a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	// Set default role if not provided (e.g., 3 = "system_user")
	if input.RoleID == 0 {
		input.RoleID = 3 // Default to system_user role
	}

	user := &data.User{
		Email:       input.Email,
		Activated:   false,
		PersonnelID: input.PersonnelID,
		RoleID:      input.RoleID, // Add this
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	if data.ValidateEmail(v, user.Email); !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	if data.ValidatePasswordPlaintext(v, input.Password); !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	if data.ValidateUser(v, user); !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = a.models.Users.Insert(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			a.failedValidationResponse(w, r, v.Errors)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	// Generate an activation token
	token, err := a.models.Tokens.New(user.ID, 3*24*time.Hour, data.ScopeActivation)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	// Send welcome email in background (if you have mailer set up)
	// a.background(func() {
	// 	data := map[string]any{
	// 		"activationToken": token.Plaintext,
	// 		"userID":          user.ID,
	// 	}
	// 	err = a.mailer.Send(user.Email, "user_welcome.tmpl", data)
	// 	if err != nil {
	// 		a.logger.Error(err.Error())
	// 	}
	// })

	// For now, return the activation token in the response (in production, send via email)
	err = a.writeJSON(w, http.StatusCreated, envelope{
		"user":             user,
		"activation_token": token.Plaintext,
	}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) activateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		TokenPlaintext string `json:"token"`
	}

	err := a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	if data.ValidateTokenPlaintext(v, input.TokenPlaintext); !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := a.models.Users.GetForToken(data.ScopeActivation, input.TokenPlaintext)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			v.AddError("token", "invalid or expired activation token")
			a.failedValidationResponse(w, r, v.Errors)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	user.Activated = true

	err = a.models.Users.Update(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			a.editConflictResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	// Delete all activation tokens for the user
	err = a.models.Tokens.DeleteAllForUser(data.ScopeActivation, user.ID)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	err = a.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *application) createAuthenticationTokenHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	data.ValidateEmail(v, input.Email)
	data.ValidatePasswordPlaintext(v, input.Password)

	if !v.IsEmpty() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := a.models.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.invalidCredentialsResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if !match {
		a.invalidCredentialsResponse(w, r)
		return
	}

	// Generate authentication token valid for 24 hours
	token, err := a.models.Tokens.New(user.ID, 24*time.Hour, data.ScopeAuthentication)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	err = a.writeJSON(w, http.StatusCreated, envelope{"authentication_token": token}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

