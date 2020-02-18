package teststore_test

import (
	"testing"

	"github.com/skvoch/burst/internal/app/model"
	"github.com/skvoch/burst/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestBooksRepository_GetByType(t *testing.T) {

	s := teststore.New()
	s.Books().RemoveAll()

	typeFirst := &model.Type{
		ID:   0,
		Name: "Type first",
	}

	typeSecond := &model.Type{
		ID:   1,
		Name: "Type second",
	}

	s.Types().Create(typeFirst)
	s.Types().Create(typeSecond)

	for i := 0; i < 10; i++ {
		book1 := model.NewTestBook()

		book1.Type = typeFirst.ID
		err := s.Books().Create(book1)
		assert.NoError(t, err)

		book2 := model.NewTestBook()

		book2.Type = typeSecond.ID
		err = s.Books().Create(book2)
		assert.NoError(t, err)
	}

	books, err := s.Books().GetByType(typeFirst)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(books))

	books, err = s.Books().GetByType(typeSecond)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(books))
}
