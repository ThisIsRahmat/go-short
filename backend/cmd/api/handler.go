package main

import (
	"fmt"
	_ "fmt"
	"hash/crc32"
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

//this is a helper function
// the hashing function uses the golang hash package, takes in the input of the longURL
// and returns the hashed URL

func hashURL(longURL string) uint32 {

	crc32Table := crc32.MakeTable(crc32.IEEE)
	url_hash := crc32.Checksum([]byte(longURL), crc32Table)

	return url_hash
}
