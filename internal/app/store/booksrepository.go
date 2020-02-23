package store

import "github.com/skvoch/burst/internal/app/model"

// BooksRepository ...
type BooksRepository interface {
	Create(b *model.Book) error
	GetByType(t *model.Type) ([]*model.Book, error)
	GetByID(ID int) (*model.Book, error)
	UpdatedPreviewPath(ID int, path string) error
	UpdatedFilePath(ID int, path string) error
	RemoveAll() error
}
