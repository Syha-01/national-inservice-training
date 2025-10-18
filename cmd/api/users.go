package main

/*
func (a *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		PersonnelID int64  `json:"personnel_id"`
	}
	err := a.readJSON(w, r, &input)
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}
	user := &data.User{
		Email:     input.Email,
		Activated: false,
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
	token, err := a.models.Tokens.New(user.ID, 3*24*time.Hour, data.ScopeActivation)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
	a.background(func() {
		data := map[string]any{
			"activationToken": token.Plaintext,
			"userID":          user.ID,
		}
		err = a.mailer.Send(user.Email, "user_welcome.tmpl", data)
		if err != nil {
			a.logger.PrintError(err, nil)
		}
	})
	err = a.writeJSON(w, http.StatusAccepted, envelope{"user": user}, nil)
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
*/
