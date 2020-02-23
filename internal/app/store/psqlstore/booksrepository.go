package psqlstore

import "github.com/skvoch/burst/internal/app/model"

// BooksRepository ...
type BooksRepository struct {
	store *Store
}

// RemoveAll ...
func (b *BooksRepository) RemoveAll() error {
	if _, err := b.store.db.Query("TRUNCATE TABLE books CASCADE;"); err != nil {
		return err
	}

	return nil
}

// Create ...
func (b *BooksRepository) Create(book *model.Book) error {
	if err := b.store.db.QueryRow(
		"INSERT INTO books (name, description, review, rating, file_path, preview_path, type) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		book.Name,
		book.Description,
		book.Review,
		book.Rating,
		book.FilePath,
		book.PreviewPath,
		book.Type,
	).Scan(&book.ID); err != nil {
		return err
	}

	return nil
}

// GetByType ...
func (b *BooksRepository) GetByType(_type *model.Type) ([]*model.Book, error) {

	rows, err := b.store.db.Query(
		"SELECT id, name, description, review, rating,file_path, preview_path, type FROM books WHERE type = $1",
		_type.ID,
	)

	if err != nil {
		return nil, err
	}

	var books []*model.Book

	for rows.Next() {
		book := &model.Book{}

		err := rows.Scan(&book.ID,
			&book.Name,
			&book.Description,
			&book.Review,
			&book.Rating,
			&book.FilePath,
			&book.PreviewPath,
			&book.Type)

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

// GetByID ...
func (b *BooksRepository) GetByID(ID int) (*model.Book, error) {

	book := &model.Book{}
	if err := b.store.db.QueryRow(
		"SELECT id, name, description, review, rating,file_path, preview_path, type FROM books WHERE id = $1",
		ID,
	).Scan(&book.ID,
		&book.Name,
		&book.Description,
		&book.Review,
		&book.Rating,
		&book.FilePath,
		&book.PreviewPath,
		&book.Type); err != nil {
		return nil, err
	}

	return book, nil

}

func (b *BooksRepository) UpdatedPreviewPath(ID int, path string) error {

	_, err := b.store.db.Exec("UPDATE books SET preview_path = $1 WHERE id = $2;", path, ID)

	return err
}

func (b *BooksRepository) UpdatedFilePath(ID int, path string) error {

	_, err := b.store.db.Exec("UPDATE books SET file_path = $1 WHERE id = $2;", path, ID)

	return err
}
