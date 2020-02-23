package teststore

import "github.com/skvoch/burst/internal/app/model"

type TypesRepository struct {
	types map[int]*model.Type
}

// RemoveAll ...
func (t *TypesRepository) RemoveAll() error {
	t.types = make(map[int]*model.Type)

	return nil
}

func (t *TypesRepository) GetByID(id int) (*model.Type, error) {

	return t.types[id], nil
}

// Create ...
func (t *TypesRepository) Create(_type *model.Type) error {

	index := len(t.types)

	_type.ID = index
	t.types[index] = _type

	return nil
}

// GetAll ...
func (t *TypesRepository) GetAll() ([]*model.Type, error) {

	result := make([]*model.Type, 0)

	for _, _type := range t.types {
		result = append(result, _type)
	}

	return result, nil
}
