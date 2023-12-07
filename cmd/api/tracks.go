package main

import (
	"fmt"
	"net/http"
	"time"

	"seu.tyfacey.dev/internal/data"
)

func (app *application) createTrackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showTrackHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	track := data.Track{
		ID:          id,
		CreatedAt:   time.Now(),
		Title:       "Cirandar",
		Artist:      []string{"Seu Jorge", "Almaz"},
		Year:        2010,
		Description: `The word "cirandar" refers to an old Portuguese dance. Some say this is Seu Jorge's serenade to the sea in which he asks an ocean goddess to "protect the fisherman" and "give the singers a good voice." It's unmistakably Brazilian, but it also has the undertow of that '60s surf guitar sound.`,
		Duration:    249000,
		Genres:      []string{"Funk", "Soul", "rock"},
		Links:       map[string]string{"track": "https://open.spotify.com/track/5LqMl8aFcuVoKX3ERagNjX?si=086043aec2034527", "video": "https://www.youtube.com/watch?v=KT8Zp22Sois"},
		Version:     1,
	}

	err = app.writeJSON(w, http.StatusOK, JSONEnvelope{"track": track}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
