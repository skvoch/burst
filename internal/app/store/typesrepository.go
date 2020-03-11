package store

import "github.com/skvoch/burst/internal/app/model"

// PreviewTokenRepository - provide access to manipulating with tokens for book preview
type TypesRepository interface {
	Create(t *model.Type) error
	GetAll() ([]*model.Type, error)
	GetByID(ID int) (*model.Type, error)
	RemoveAll() error
}
