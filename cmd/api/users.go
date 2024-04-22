package main

import (
	"Greenlight/internal/data"
	"Greenlight/internal/validator"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Add a createModuleInfoHandler for the "POST /v1/modules" endpoint
func (app *application) createUserInfoHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		UserName    string   `json:"user_name"`
		UserSurname string   `json:"user_surname"`
		Email       string   `json:"email"`
		Role        []string `json:"role"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	module := &data.User{
		UserName:    input.UserName,
		UserSurname: input.UserSurname,
		Email:       input.Email,
		Role:        input.Role,
	}

	v := validator.New()
	// Call the ValidateModule() function
	if data.ValidateUser(v, module); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	//Insert() method on our movies model
	err = app.models.UserInfo.Insert(module)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/users/%d", module.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"module": module}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

// getModuleInfoHandler for the "GET /v1/modules/:id" endpoint
func (app *application) getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {

		app.notFoundResponse(w, r)
		return
	}

	module, err := app.models.UserInfo.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"module": module}, nil)
	if err != nil {

		app.serverErrorResponse(w, r, err)
	}

}

// editModuleInfoHandler for the "PUT /v1/modules/:id" endpoint
func (app *application) editUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	module, err := app.models.UserInfo.Get(id)
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
		UserName    string   `json:"user_name"`
		UserSurname string   `json:"user_surname"`
		Email       string   `json:"email"`
		Role        []string `json:"role"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	module.UpdatedAt = time.Now()
	module.UserName = input.UserName
	module.UserSurname = input.UserSurname
	module.Email = input.Email
	module.Role = input.Role

	v := validator.New()

	if data.ValidateUser(v, module); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.UserInfo.Update(module)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"module": module}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

// deleteModuleInfoHandler for the "DELETE /v1/modules/:id" endpoint
func (app *application) deleteUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.UserInfo.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "user successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
