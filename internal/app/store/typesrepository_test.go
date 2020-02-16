package store_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/skvoch/burst/internal/app/model"
	"github.com/skvoch/burst/internal/app/store"
)

func TestTypeRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("types")

	repo := s.Types()

	_type := &model.Type{
		ID:   0,
		Name: "Nothing",
	}

	err := repo.Create(_type)
	assert.NoError(t, err)
}

func TestTypeRepository_GetAll(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("types")

	repo := s.Types()

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
