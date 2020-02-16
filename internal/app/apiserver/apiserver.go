package apiserver

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/skvoch/burst/internal/app/store/psqlstore"
)

// Start ...
func Start(config *Config) error {

	db, err := newDB(config.DataBaseURL)

	if err != nil {
		return err
	}

	defer db.Close()
	store := psqlstore.New(db)
	server := newServer(store)

	log.Println("Starting HTTP server on", config.BindAddr, "...")
	return http.ListenAndServe(config.BindAddr, server)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
