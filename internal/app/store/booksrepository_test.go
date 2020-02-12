package store_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/skvoch/burst/internal/app/model"
	"github.com/skvoch/burst/internal/app/store"
)

func TestBooksRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("books")

	book := &model.Book{
		ID:          0,
		Name:        "Golang book",
		Description: "Super cool book",
		Review:      "I want to recomend it for you!",
		Rating:      5,
		Type:        0,
	}

	err := s.Books().Create(book)

	assert.NoError(t, err)
}
