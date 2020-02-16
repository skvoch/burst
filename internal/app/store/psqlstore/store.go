package psqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	db              *sql.DB
	booksRepository *BooksRepository
	typesRepository *TypesRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

/*
func (s *Store) getSourceName() string {

	dataSourceName := s.config.DatabaseUser
	dataSourceName += " "
	dataSourceName += s.config.DatabasePassword
	dataSourceName += " "
	dataSourceName += s.config.DatabaseURL

	return dataSourceName
}
*/

// Books ...
func (s *Store) Books() *BooksRepository {
	if s.booksRepository != nil {
		return s.booksRepository
	}

	s.booksRepository = &BooksRepository{
		store: s,
	}

	return s.booksRepository
}

// Types ...
func (s *Store) Types() *TypesRepository {
	if s.typesRepository != nil {
		return s.typesRepository
	}

	s.typesRepository = &TypesRepository{
		store: s,
	}

	return s.typesRepository
}
