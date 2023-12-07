package main

import (
	"fmt"
	"net/http"
)

// logError is a generic helper for logging an error mesage and
// request method.
func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)
	app.logger.Error(err.Error(), "method", method, "uri", uri)
}

// errorResponse is a generic helper for sending JSON-formatted error
// messages to the client with a given status code.
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	jsonEnvelope := JSONEnvelope{"error": message}

	err := app.writeJSON(w, status, jsonEnvelope, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	msg := "the request resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, msg)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("The %s method i not allowed for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, msg)
}
