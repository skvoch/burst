package psqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/skvoch/burst/internal/app/model"
	"github.com/skvoch/burst/internal/app/store/psqlstore"
)

func TestTypeRepository_Create(t *testing.T) {
	db, teardown := psqlstore.TestDB(t, databaseURL)
	defer teardown("types")

	s := psqlstore.New(db)
	repo := s.Types()

	_type := &model.Type{
		ID:   0,
		Name: "Nothing",
	}

	err := repo.Create(_type)
	assert.NoError(t, err)
}

func TestTypeRepository_GetAll(t *testing.T) {
	db, teardown := psqlstore.TestDB(t, databaseURL)
	defer teardown("types")

	s := psqlstore.New(db)
	s.Books().RemoveAll()

	repo := s.Types()
	repo.RemoveAll()

	_type := &model.Type{
		ID:   0,
		Name: "Nothing",
	}

	for i := 0; i < 10; i++ {
		err := repo.Create(_type)
		assert.NoError(t, err)
	}

	types, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 10, len(types))
}

func TestTypeRepository_GetByID(t *testing.T) {
	db, teardown := psqlstore.TestDB(t, databaseURL)
	defer teardown("types")

	s := psqlstore.New(db)
	s.Books().RemoveAll()

	repo := s.Types()
	repo.RemoveAll()

	_type := &model.Type{
		ID:   0,
		Name: "Nothing",
	}

	err := repo.Create(_type)
	assert.NoError(t, err)
	_typeFound, err := repo.GetByID(_type.ID)
	assert.NoError(t, err)
	assert.NotNil(t, _typeFound)

}
