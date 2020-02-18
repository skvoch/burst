package teststore_test

import (
	"testing"

	"github.com/skvoch/burst/internal/app/model"
	"github.com/skvoch/burst/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestTypeRepository_Create(t *testing.T) {

	s := teststore.New()
	repo := s.Types()

	_type := &model.Type{
		ID:   0,
		Name: "Nothing",
	}

	err := repo.Create(_type)
	assert.NoError(t, err)
}

func TestTypeRepository_GetAll(t *testing.T) {

	s := teststore.New()
	s.Books().RemoveAll()

	repo := s.Types()
	repo.RemoveAll()

	for i := 0; i < 10; i++ {
		_type := model.NewTestType()
		err := repo.Create(_type)
		assert.NoError(t, err)
	}

	types, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, 10, len(types))
}
