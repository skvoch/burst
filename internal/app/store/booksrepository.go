package store

import "github.com/skvoch/burst/internal/app/model"

// BooksRepository - provide access to manipulating with books
type BooksRepository interface {
	Create(b *model.Book) error
	GetByType(t *model.Type) ([]*model.Book, error)
	GetByID(ID int) (*model.Book, error)
	UpdatePreviewPath(ID int, path string) error
	UpdateFilePath(ID int, path string) error
	RemoveAll() error
}
