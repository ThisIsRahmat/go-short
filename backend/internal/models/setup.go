package models

import (
	"database/sql"
	"fmt"
)

func OpenDB() (*sql.DB, error) {

	// Use sql.Open() to create an empty connection pool, using the DSN that is stored in psqlInfo

	psqlInfo := fmt.Sprintf("postgres", "user=go_short password=secret dbname=go_short sslmode=disable")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic("Failed to connect to database")
	}

	// Return the sql.DB connection pool.
	return db, nil

}
