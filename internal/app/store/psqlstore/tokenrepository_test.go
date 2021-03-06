package psqlstore_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/skvoch/burst/internal/app/model"
	"github.com/skvoch/burst/internal/app/store/psqlstore"
	"github.com/stretchr/testify/assert"
)

func TestPDFTokenRepository_CreateAndGetByUID(t *testing.T) {

	db, teardown := psqlstore.TestDB(t, databaseURL)
	defer teardown("pdf_tokens", "books", "types")
	s := psqlstore.New(db)
	tokensRepo := s.TokensPDF()
	booksRepo := s.Books()
	typesRepo := s.Types()

	if err := tokensRepo.RemoveAll(); err != nil {
		assert.NoError(t, err)
	}

	_type := model.NewTestType()
	typesRepo.Create(_type)
	book := model.NewTestBook()
	book.Type = _type.ID
	booksRepo.Create(book)

	_token := &model.PDFToken{
		UID:    uuid.New().String(),
		BookID: book.ID,
	}

	err := tokensRepo.Create(_token)
	assert.NoError(t, err)

	{
		token, err := tokensRepo.GetByUID(_token.UID)
		assert.NoError(t, err)
		assert.Equal(t, _token.UID, token.UID)
		assert.Equal(t, _token.BookID, token.BookID)
	}
}

func TestPDFTokenRepository_Remove(t *testing.T) {
	db, teardown := psqlstore.TestDB(t, databaseURL)
	defer teardown("pdf_tokens", "books", "types")
	s := psqlstore.New(db)
	tokensRepo := s.TokensPDF()
	booksRepo := s.Books()
	typesRepo := s.Types()

	if err := tokensRepo.RemoveAll(); err != nil {
		assert.NoError(t, err)
	}

	_type := model.NewTestType()
	typesRepo.Create(_type)
	book := model.NewTestBook()
	book.Type = _type.ID
	booksRepo.Create(book)

	_token := &model.PDFToken{
		UID:    uuid.New().String(),
		BookID: book.ID,
	}

	err := tokensRepo.Create(_token)
	assert.NoError(t, err)

	err = tokensRepo.Remove(_token)
	assert.NoError(t, err)
}

func TestPreviewTokenRepository_CreateAndGetByUID(t *testing.T) {

	db, teardown := psqlstore.TestDB(t, databaseURL)
	defer teardown("pdf_tokens", "books", "types")
	s := psqlstore.New(db)
	tokensRepo := s.TokensPreview()
	booksRepo := s.Books()
	typesRepo := s.Types()

	if err := tokensRepo.RemoveAll(); err != nil {
		assert.NoError(t, err)
	}

	_type := model.NewTestType()
	typesRepo.Create(_type)
	book := model.NewTestBook()
	book.Type = _type.ID
	booksRepo.Create(book)

	_token := &model.PreviewToken{
		UID:    uuid.New().String(),
		BookID: book.ID,
	}

	err := tokensRepo.Create(_token)
	assert.NoError(t, err)

	{
		token, err := tokensRepo.GetByUID(_token.UID)
		assert.NoError(t, err)
		assert.Equal(t, _token.UID, token.UID)
		assert.Equal(t, _token.BookID, token.BookID)
	}
}

func TestPreviewTokenRepository_Remove(t *testing.T) {
	db, teardown := psqlstore.TestDB(t, databaseURL)
	defer teardown("pdf_tokens", "books", "types")
	s := psqlstore.New(db)
	tokensRepo := s.TokensPreview()
	booksRepo := s.Books()
	typesRepo := s.Types()

	if err := tokensRepo.RemoveAll(); err != nil {
		assert.NoError(t, err)
	}

	_type := model.NewTestType()
	typesRepo.Create(_type)
	book := model.NewTestBook()
	book.Type = _type.ID
	booksRepo.Create(book)

	_token := &model.PreviewToken{
		UID:    uuid.New().String(),
		BookID: book.ID,
	}

	err := tokensRepo.Create(_token)
	assert.NoError(t, err)

	err = tokensRepo.Remove(_token)
	assert.NoError(t, err)
}
