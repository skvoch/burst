package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/skvoch/burst/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServerHandleTypesGet(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/types", nil)
	s := newServer(teststore.New())

	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
