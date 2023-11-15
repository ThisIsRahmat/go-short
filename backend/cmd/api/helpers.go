package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// Encode the data to JSON, returning the error if there was one.
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Append a newline to make it easier to view in terminal applications.
	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	// Add the "Content-Type: application/json" header, then write the status code and
	// JSON response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil

}

func (app *application) readFormParam(r *http.Request) (string, error) {
	// Retrieve the value of the 'longURL' form parameter
	longURL := r.FormValue("longURL")

	// Check if the 'longURL' parameter is empty
	if longURL == "" {
		return "", errors.New("longURL parameter is required")
	}

	return longURL, nil

}
