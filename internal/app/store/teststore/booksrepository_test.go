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
		bookFirst := model.NewTestBook()

		bookFirst.Type = typeFirst.ID
		err := s.Books().Create(bookFirst)
		assert.NoError(t, err)

		bookSecond := model.NewTestBook()

		bookSecond.Type = typeSecond.ID
		err = s.Books().Create(bookSecond)
		assert.NoError(t, err)
	}

	books, err := s.Books().GetByType(typeFirst)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(books))

	books, err = s.Books().GetByType(typeSecond)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(books))
}

func TestBooksRepository_UpdateFilePath(t *testing.T) {
	s := teststore.New()
	s.Books().RemoveAll()

	typeFirst := &model.Type{
		ID:   0,
		Name: "Type first",
	}

	dstFilePath := "example/of/file/path"

	book := model.NewTestBook()

	s.Types().Create(typeFirst)
	s.Books().Create(book)

	err := s.Books().UpdatedFilePath(book.ID, dstFilePath)
	assert.NoError(t, err)

	foundBook, err := s.Books().GetByID(book.ID)
	assert.NoError(t, err)

	assert.Equal(t, book.FilePath, foundBook.FilePath)
}

func TestBooksRepository_UpdatePreviewPath(t *testing.T) {
	s := teststore.New()
	s.Books().RemoveAll()

	typeFirst := &model.Type{
		ID:   0,
		Name: "Type first",
	}

	dstPreviewPath := "example/of/preview/path"

	book := model.NewTestBook()

	s.Types().Create(typeFirst)
	s.Books().Create(book)

	err := s.Books().UpdatedPreviewPath(book.ID, dstPreviewPath)
	assert.NoError(t, err)

	foundBook, err := s.Books().GetByID(book.ID)
	assert.NoError(t, err)

	assert.Equal(t, book.PreviewPath, foundBook.PreviewPath)
}
