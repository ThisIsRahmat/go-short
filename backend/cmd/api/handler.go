package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

// struct to form data

type ShortURLFormData struct {
	LongURL string `json:"longURL"`
}

//handler that does healthchecks for our api

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *application) createShortURLHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "creating a new short url")

	// Parse the JSON-formatted request body into ShortURLFormData struct
	var formData ShortURLFormData

	if err := json.NewDecoder(r.Body).Decode(&formData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Read the longURL from the request body, query parameters, or wherever it's supposed to come from.
	longURL := formData.LongURL

	shortURL, err := app.models.Create(longURL)
	if err != nil {
		//to-do better error response
		http.Error(w, "Error creating short URL: "+err.Error(), http.StatusInternalServerError)
		return

	}

	//Respond to the client

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Created Short URL: %s", shortURL)

	// Write a JSON response with a 201 Created status code, including the short URL in the response body,
	// and the Location header.
	err = app.writeJSON(w, http.StatusCreated, envelope{"shortURL": shortURL}, nil)
	if err != nil {
		fmt.Fprintf(w, "")
	}

}

func (app *application) getShortURLHandler(w http.ResponseWriter, r *http.Request) {

	//reads the longURL from the form parameter
	longURL, err := app.readFormParam(r)
	if err != nil {
		// app.notFoundResponse(w, r)
		panic(err)
		return
	}

	fmt.Fprintln(w, "retrieve a short url from the database")

	// Call the Get() method to fetch the data of longURL

	shortURL, err := app.models.Get(longURL)
	if err != nil {
		switch {
		case errors.Is(err, ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"shortURL": shortURL}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
