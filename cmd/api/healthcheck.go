package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	if err := app.writeJSON(w, http.StatusOK, data, nil); err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Server could not process your request", http.StatusInternalServerError)
	}
}
