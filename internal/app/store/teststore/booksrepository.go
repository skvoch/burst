package teststore

import (
	"github.com/skvoch/burst/internal/app/model"
	"golang.org/x/crypto/openpgp/errors"
)

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

func (b *BooksRepository) GetByID(ID int) (*model.Book, error) {

	for _, book := range b.books {
		if book.ID == ID {
			return book, nil
		}
	}

	return nil, nil
}

func (b *BooksRepository) UpdatedPreviewPath(ID int, path string) error {

	book := b.books[ID]

	if book == nil {
		return errors.ErrKeyIncorrect
	}

	book.PreviewPath = path

	return nil
}

func (b *BooksRepository) UpdatedFilePath(ID int, path string) error {
	book := b.books[ID]

	if book == nil {
		return errors.ErrKeyIncorrect
	}

	book.FilePath = path

	return nil
}
