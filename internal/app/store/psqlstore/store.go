package psqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
	"github.com/skvoch/burst/internal/app/store"
)

// Store ...
type Store struct {
	db                     *sql.DB
	booksRepository        *BooksRepository
	typesRepository        *TypesRepository
	pdfTokenRepository     *PDFTokenRepository
	previewTokenRepository *PreviewTokenRepository
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
func (s *Store) Books() store.BooksRepository {
	if s.booksRepository != nil {
		return s.booksRepository
	}

	s.booksRepository = &BooksRepository{
		store: s,
	}

	return s.booksRepository
}

// Types ...
func (s *Store) Types() store.TypesRepository {
	if s.typesRepository != nil {
		return s.typesRepository
	}

	s.typesRepository = &TypesRepository{
		store: s,
	}

	return s.typesRepository
}

func (s *Store) TokensPDF() store.PDFTokenRepository {
	if s.pdfTokenRepository != nil {
		return s.pdfTokenRepository
	}

	s.pdfTokenRepository = &PDFTokenRepository{
		store: s,
	}

	return s.pdfTokenRepository
}

func (s *Store) TokensPreview() store.PreviewTokenRepository {
	if s.previewTokenRepository != nil {
		return s.previewTokenRepository
	}

	s.previewTokenRepository = &PreviewTokenRepository{
		store: s,
	}

	return s.previewTokenRepository
}
