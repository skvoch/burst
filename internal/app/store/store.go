package store

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	config *Config
	db     *sql.DB
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {

	dataSourceName := s.getSourceName()
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) getSourceName() string {

	dataSourceName := s.config.DatabaseUser
	dataSourceName += " "
	dataSourceName += s.config.DatabasePassword
	dataSourceName += " "
	dataSourceName += s.config.DatabaseURL

	return dataSourceName
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}
