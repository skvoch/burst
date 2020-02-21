package teststore_test

import (
	"testing"

	"github.com/skvoch/burst/internal/app/model"
	"github.com/skvoch/burst/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestPDFTokenRepository_CreateAndGetByUID(t *testing.T) {

	s := teststore.New()
	repo := s.TokensPDF()

	_token := &model.PDFToken{
		UID:    "zxcvbnm",
		BookID: 0,
	}

	err := repo.Create(_token)
	assert.NoError(t, err)

	{
		token, err := repo.GetByUID(_token.UID)
		assert.NoError(t, err)
		assert.Equal(t, _token.UID, token.UID)
		assert.Equal(t, _token.BookID, token.BookID)
	}
}

func TestPDFTokenRepository_Remove(t *testing.T) {

	s := teststore.New()
	repo := s.TokensPDF()

	_token := &model.PDFToken{
		UID:    "zxcvbnm",
		BookID: 0,
	}

	err := repo.Create(_token)
	assert.NoError(t, err)

	err = repo.Remove(_token)
	assert.NoError(t, err)
}
