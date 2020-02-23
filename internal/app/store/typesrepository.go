package store

import "github.com/skvoch/burst/internal/app/model"

// TypesRepository ...
type TypesRepository interface {
	Create(t *model.Type) error
	GetAll() ([]*model.Type, error)
	GetByID(ID int) (*model.Type, error)
	RemoveAll() error
}
