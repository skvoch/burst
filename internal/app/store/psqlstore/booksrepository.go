package psqlstore

import "github.com/skvoch/burst/internal/app/model"

// BooksRepository ...
type BooksRepository struct {
	store *Store
}

// RemoveAll ...
func (b *BooksRepository) RemoveAll() error {
	if _, err := b.store.db.Query("RUNCATE TABLE books CASCADE;"); err != nil {
		return err
	}

	return nil
}

// Create ...
func (b *BooksRepository) Create(book *model.Book) error {
	if err := b.store.db.QueryRow(
		"INSERT INTO books (name, description, review, rating, type) VALUES ($1, $2, $3, $4, $5) RETURNING id",
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
func (b *BooksRepository) GetByType(_type *model.Type) ([]*model.Book, error) {

	rows, err := b.store.db.Query(
		"SELECT id, name, description, review, rating, type FROM books WHERE type = $1",
		_type.ID,
	)

	if err != nil {
		return nil, err
	}

	var books []*model.Book

	for rows.Next() {
		book := &model.Book{}

		err := rows.Scan(&book.ID, &book.Name, &book.Description, &book.Review, &book.Rating, &book.Type)

		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
