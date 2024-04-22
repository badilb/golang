package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/modules", app.createModuleInfoHandler)
	router.HandlerFunc(http.MethodGet, "/v1/modules/:id", app.getModuleInfoHandler)
	router.HandlerFunc(http.MethodPut, "/v1/modules/:id", app.editModuleInfoHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/modules/:id", app.deleteModuleInfoHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.createUserInfoHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/:id", app.getUserInfoHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/:id", app.editUserInfoHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/users/:id", app.deleteUserInfoHandler)

	router.HandlerFunc(http.MethodGet, "/v1/teachers/:id", app.GetAllTeachersInfo)
	return router
}
