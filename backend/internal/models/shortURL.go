package models

import (
	"database/sql"
	"errors"
	"hash/crc32"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

// Define a ShortURL type to hold the data for an individual snippet. Notice how
// the fields of the struct correspond to the fields in our MySQL snippets
// table?
type URL_db struct {
	hashkey     []byte    `json:"hashkey"`
	shortURL    string    `json:"shortURL"`
	longURL     string    `json:"longURL"`
	CreatedOn   time.Time `json:""`
	ExpiresOn   time.Time `json:"-"`
	LastVisited time.Time `json:"-"`
}

// Define a ShortURLModel type which wraps a sql.DB connection pool.
type ShortURLModel struct {
	DB *sql.DB
}

// Define a custom ErrRecordNotFound error. We'll return this from our Get() method when
// looking up a URL that doesn't exist in the database
var (
	ErrRecordNotFound = errors.New("record not found")
)

// creates short_url and stores it in database
// This will insert a new short url into the database.
func (u *ShortURLModel) Create(longURL string) (string, error) {

	const baseURL string = "www.go-short.vercel.app/"

	// get the hashed key from the helper function
	hashkey := hashURL(longURL)

	shortenedURL := baseURL + strconv.FormatUint(uint64(hashkey), 10)

	// Define the SQL query for inserting a new record in the URL table
	query := `
	 INSERT INTO
	  URL_db(hashkey, shortURL, longURL)
	   VALUES($1,$2,$3)`

	//updates the database with shorturl ,
	_, err := u.DB.Exec(query, hashkey, shortenedURL, longURL)

	if err != nil {
		return "", err
	}

	return shortenedURL, nil

}

// retrieves long url
// This will return a specific long url based on the hashkey.
func (u *ShortURLModel) Get(shortURL string) (string, error) {

	var longURL string

	// Define the SQL query for retrieving the long url data.
	query := `
  SELECT longURL
  FROM URL_db
  WHERE shortURL = $1`

	//query the database and fetch the longURL
	err := u.DB.QueryRow(query, shortURL).Scan(&longURL)

	// error checking
	if err != nil {
		// If row is not found return a specific error.
		if err == sql.ErrNoRows {
			return "", errors.New("URL not found")
		}
		return " ", err
	}

	return longURL, nil
}

// hashURL is a helper function
// the hashing function uses the golang hash package, takes in the input of the longURL
// and returns the hashed URL

func hashURL(longURL string) uint32 {

	crc32Table := crc32.MakeTable(crc32.IEEE)
	url_hash := crc32.Checksum([]byte(longURL), crc32Table)

	return url_hash
}
