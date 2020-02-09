package store

import (
	"github.com/skvoch/burst/internal/model"
)

// BooksRepository ...
type BooksRepository struct {
	store *Store
}

// Create ...
func (b *BooksRepository) Create(book *model.Book) error {
	if err := b.store.db.QueryRow(
		"INSERT INTRO books (name, description, review, rating, type) VALUES ($1, $2, $3, $4, $5) RETURNING id;",
		book.Name,
		book.Description,
		book.Review,
		book.Rating,
		book.Type,
	).Scan(&book.ID); err != nil {
		return err
	}

	return nil
}

// GetByType ...
func (b *BooksRepository) GetByType(_type *model.Type) (*model.Book, error) {
	return nil, nil
}
