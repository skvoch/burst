package psqlstore

import (
	"github.com/skvoch/burst/internal/app/model"
)

// TypesRepository ...
type TypesRepository struct {
	store *Store
}

// RemoveAll ...
func (b *TypesRepository) RemoveAll() error {
	if _, err := b.store.db.Query("TRUNCATE TABLE types CASCADE;"); err != nil {
		return err
	}

	return nil
}

// Create ...
func (t *TypesRepository) Create(_type *model.Type) error {
	if err := t.store.db.QueryRow(
		"INSERT INTO types (name) VALUES ($1) RETURNING id",
		_type.Name,
	).Scan(&_type.ID); err != nil {
		return err
	}

	return nil
}

// GetAll ...
func (t *TypesRepository) GetAll() ([]*model.Type, error) {
	rows, err := t.store.db.Query("SELECT * FROM types")

	if err != nil {
		return nil, err
	}
	var types []*model.Type

	for rows.Next() {
		_type := &model.Type{}

		if err := rows.Scan(&_type.ID, &_type.Name); err != nil {
			return nil, err
		}

		types = append(types, _type)
	}

	return types, nil
}
