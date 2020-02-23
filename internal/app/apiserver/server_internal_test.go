package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/skvoch/burst/internal/app/model"
	"github.com/skvoch/burst/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestServerHandleTypesCreateAndGet(t *testing.T) {
	log := logrus.New()
	s := newServer(teststore.New(), log)

	types := make([]*model.Type, 0)

	types = append(types, &model.Type{ID: 0, Name: "C++ books"})
	types = append(types, &model.Type{ID: 1, Name: "C# books"})
	types = append(types, &model.Type{ID: 2, Name: "Go books"})
	types = append(types, &model.Type{ID: 3, Name: "Math books"})

	for _, _type := range types {
		json, err := json.Marshal(_type)
		assert.NoError(t, err)

		reader := bytes.NewReader(json)

		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/v1.0/types/create/", reader)

		s.ServeHTTP(rec, req)

		assert.Equal(t, rec.Code, http.StatusCreated)
	}

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1.0/types/", nil)

	s.ServeHTTP(rec, req)
	foundTypes := make([]*model.Type, 0)
	json.Unmarshal(rec.Body.Bytes(), &foundTypes)

	for index := 0; index < len(types); index++ {
		assert.Equal(t, types[index].ID, foundTypes[index].ID)
		assert.Equal(t, types[index].Name, foundTypes[index].Name)
	}
}
