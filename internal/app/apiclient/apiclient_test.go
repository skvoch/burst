package apiclient_test

import (
	"testing"

	"github.com/skvoch/burst/internal/app/apiclient"
	"github.com/skvoch/burst/internal/app/model"
	"github.com/stretchr/testify/assert"
)

const (
	URL string = "localhost:8080"
)

func TestConnection(t *testing.T) {
	client, err := apiclient.NewTestClient(URL)

	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestTypes(t *testing.T) {
	client, err := apiclient.NewTestClient(URL)

	assert.NoError(t, err)
	assert.NotNil(t, client)

	if err := client.RemoveAllTypes(); err != nil {
		assert.Fail(t, err.Error())
	}

	types := make([]*model.Type, 0)
	types = append(types, &model.Type{Name: "C++ Books"})
	types = append(types, &model.Type{Name: "C# Books"})
	types = append(types, &model.Type{Name: "Go Books"})

	for _, _type := range types {
		err := client.CreateType(_type)
		assert.NoError(t, err)
	}

	foundTypes, err := client.GetAllTypes()

	assert.NoError(t, err)

	for i := 0; i < len(types); i++ {
		assert.Equal(t, types[i].Name, foundTypes[i].Name)
	}
}
