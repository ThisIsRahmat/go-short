package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/thisisrahmat/go-short/internal/models"
	_ "github.com/thisisrahmat/go-short/internal/models"

	_ "github.com/lib/pq"
)

//A sring containiing the application version number.

const version = "1.0.0"

// A config struct to hol all the configuration settings of our application.

type config struct {
	port int
	env  string

	// add a db struct to hold configuration setting for database connection  pool
	// now it holds the dsn - which is the address of the postgres db
	db struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
}

//application struct to hold dependencies for HTTP handlers, helpers and middleware

type application struct {
	config config
	logger *slog.Logger
	models models.ShortURLModel
}

func main() {

	// declare an instance of config struct

	var cfg config

	//read value of port and env command-line flags into config struct

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	//initialise a new structure logger which writes log entries to the standard out stream

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Call the openDB() helper function (see below) to create the connection pool,
	// passing in the config struct. If this returns an error, we log it and exit the
	// application immediately.
	db, err := models.OpenDB()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// // Assign the connection pool to app.models.DB.
	// app.models.DB = db

	// Defer a call to db.Close() so that the connection pool is closed before the
	// main() function exits.
	defer db.Close()

	// Also log a message to say that the connection pool has been successfully
	// established.
	logger.Info("database connection pool established")

	// declare instance of application struct, containing config struct and logger

	app := &application{
		config: cfg,
		logger: logger,
	}

	//declare a new servermux and add a /v1/healthckec rout to dispatch requests to healthckechHandler method

	router := mux.NewRouter()

	router.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	router.HandleFunc("/v1/", app.createShortURLHandler).Methods("POST")
	router.HandleFunc("/v1/{hash}", app.getShortURLHandler).Methods("GET")

	// Declare a HTTP server which listens on the port provided in the config struct,
	// uses the servemux we created above as the handler, has some sensible timeout
	// settings and writes any log messages to the structured logger at Error level.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	// Start the HTTP server.
	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)

}
