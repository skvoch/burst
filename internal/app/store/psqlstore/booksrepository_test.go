package psqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/skvoch/burst/internal/app/model"
	"github.com/skvoch/burst/internal/app/store/psqlstore"
)

func TestBooksRepository_GetByType(t *testing.T) {
	db, teardown := psqlstore.TestDB(t, databaseURL)
	defer teardown("books", "types")

	s := psqlstore.New(db)
	s.Books().RemoveAll()

	typeFirst := &model.Type{
		ID:   0,
		Name: "Type first",
	}

	typeSecond := &model.Type{
		ID:   0,
		Name: "Type second",
	}

	s.Types().Create(typeFirst)
	s.Types().Create(typeSecond)

	book := &model.Book{
		ID:          0,
		Name:        "Golang book",
		Description: "Super cool book",
		Review:      "I want to recomend it for you!",
		Rating:      5,
		Type:        0,
	}

	for i := 0; i < 10; i++ {
		book.Type = typeFirst.ID
		err := s.Books().Create(book)
		assert.NoError(t, err)

		book.Type = typeSecond.ID
		err = s.Books().Create(book)
		assert.NoError(t, err)
	}

	books, err := s.Books().GetByType(typeFirst)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(books))

	books, err = s.Books().GetByType(typeSecond)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(books))

}
