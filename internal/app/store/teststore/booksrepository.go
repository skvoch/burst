package teststore

import "github.com/skvoch/burst/internal/app/model"

type BooksRepository struct {
	books map[int]*model.Book
}

func (b *BooksRepository) RemoveAll() error {
	b.books = make(map[int]*model.Book)

	return nil
}

// Create ...
func (b *BooksRepository) Create(book *model.Book) error {

	index := len(b.books)

	book.ID = index
	b.books[index] = book

	return nil
}

// GetByType ...
func (b *BooksRepository) GetByType(_type *model.Type) ([]*model.Book, error) {

	result := make([]*model.Book, 0)

	for _, book := range b.books {
		if book.Type == _type.ID {
			result = append(result, book)
		}
	}

	return result, nil
}
