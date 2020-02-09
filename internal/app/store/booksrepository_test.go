package store_test

import (
	"testing"

	"github.com/skvoch/burst/internal/app/store"
)

func TestBooksRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("books")

	b, err := s.Books().Create( &model.Book{
		Name: "Golang book",
		Description: "Super cool book",
		Review: "I want to recomend it for you!",
		Rating: 5,
		Type: 1,
	}); err != nil {
		t.Fatal()
	}
}
