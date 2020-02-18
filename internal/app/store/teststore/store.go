package teststore

import (
	"github.com/skvoch/burst/internal/app/model"
	"github.com/skvoch/burst/internal/app/store"
)

// Store ...
type Store struct {
	booksRepository *BooksRepository
	typesRepository *TypesRepository
}

// New ...
func New() *Store {
	return &Store{}
}

// Books ...
func (s *Store) Books() store.BooksRepository {
	if s.booksRepository != nil {
		return s.booksRepository
	}

	s.booksRepository = &BooksRepository{
		books: make(map[int]*model.Book),
	}

	return s.booksRepository
}

// Types ...
func (s *Store) Types() store.TypesRepository {
	if s.typesRepository != nil {
		return s.typesRepository
	}

	s.typesRepository = &TypesRepository{
		types: make(map[int]*model.Type),
	}

	return s.typesRepository
}
