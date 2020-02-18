package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/skvoch/burst/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServerHandleTypesGet(t *testing.T) {
	log := logrus.New()
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/types", nil)
	s := newServer(teststore.New(), log)

	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
