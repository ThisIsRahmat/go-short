package main

import (
	"fmt"
	"net/http"
)

//handler that does healthchecks for our api

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *application) createShortURLHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "creating a new short url")
}
