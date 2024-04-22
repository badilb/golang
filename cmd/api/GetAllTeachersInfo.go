package main

import (
	"net/http"
)

func (app *application) GetAllTeachersInfo(w http.ResponseWriter, r *http.Request) {
	teachers, err := app.teachers.TeacherInfo.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"teachers": teachers}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
