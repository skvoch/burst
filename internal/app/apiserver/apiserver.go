package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/skvoch/burst/internal/app/store/psqlstore"
)

// Start ...
func Start(config *Config) error {

	db, err := newDB(config.DataBaseURL)
	log := logrus.New()
	if err != nil {
		return err
	}

	defer db.Close()
	store := psqlstore.New(db)
	server := newServer(store, log)

	log.Info("Starting HTTP server on", config.BindAddr, "...")
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
