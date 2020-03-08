package apiclient_test

import (
	"crypto/md5"
	"crypto/rand"
	"io/ioutil"
	"log"
	"os"
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
		id, err := client.CreateType(_type)
		_type.ID = id
		assert.NoError(t, err)
	}

	foundTypes, err := client.GetAllTypes()

	assert.NoError(t, err)

	for i := 0; i < len(types); i++ {
		assert.Equal(t, types[i].Name, foundTypes[i].Name)
	}
}

func TestGetBooksIDsByType(t *testing.T) {
	client, err := apiclient.NewTestClient(URL)

	assert.NoError(t, err)
	assert.NotNil(t, client)

	_type := &model.Type{Name: "Go books"}

	if err := client.RemoveAllBooks(); err != nil {
		assert.NoError(t, err)
	}

	id, err := client.CreateType(_type)
	if err != nil {
		assert.NoError(t, err)
	}
	_type.ID = id

	books := make([]*model.Book, 0)

	books = append(books, model.NewTestBookWithType(_type.ID))
	books = append(books, model.NewTestBookWithType(_type.ID))
	books = append(books, model.NewTestBookWithType(_type.ID))
	books = append(books, model.NewTestBookWithType(_type.ID))
	books = append(books, model.NewTestBookWithType(_type.ID))

	for _, book := range books {
		if _, err := client.CreateBook(book); err != nil {
			assert.NoError(t, err)
		}
	}

	ids, err := client.GetBookIDs(_type.ID)

	assert.NotNil(t, ids)
	assert.NoError(t, err)
	assert.True(t, len(ids) == 5)
}

func TestGetBookPreview(t *testing.T) {
	client, err := apiclient.NewTestClient(URL)

	assert.NoError(t, err)
	assert.NotNil(t, client)

	_type := model.NewTestType()

	id, err := client.CreateType(_type)
	assert.NoError(t, err)
	_type.ID = id

	book := model.NewTestBookWithType(_type.ID)
	tokens, err := client.CreateBook(book)

	assert.NoError(t, err)
	assert.NotNil(t, tokens)
	book.ID = tokens.BookID

	previewData := make([]byte, 256)

	rand.Read(previewData)

	file, err := ioutil.TempFile("/tmp/", "preview")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	file.Write(previewData)

	err = client.SendPreview(file.Name(), book.ID, tokens.PreviewUUID)
	assert.NoError(t, err)

	fileRespond := client.GetBookPreview(tokens.BookID)

	assert.NotNil(t, fileRespond.Data)
	assert.NoError(t, err)

	hash := md5.Sum(previewData)
	foundHash := md5.Sum(fileRespond.Data)
	assert.Equal(t, hash, foundHash)
}

func TestGetBookFile(t *testing.T) {
	client, err := apiclient.NewTestClient(URL)

	assert.NoError(t, err)
	assert.NotNil(t, client)

	_type := model.NewTestType()

	id, err := client.CreateType(_type)
	assert.NoError(t, err)
	_type.ID = id

	book := model.NewTestBookWithType(_type.ID)
	tokens, err := client.CreateBook(book)

	assert.NoError(t, err)
	assert.NotNil(t, tokens)
	book.ID = tokens.BookID

	fileData := make([]byte, 256)

	rand.Read(fileData)

	file, err := ioutil.TempFile("/tmp/", "file")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	file.Write(fileData)

	err = client.SendBookFile(file.Name(), book.ID, tokens.FileUUID)
	assert.NoError(t, err)

	fileRespond := client.GetBookFile(tokens.BookID)

	assert.NotNil(t, fileRespond.Data)
	assert.NoError(t, err)

	hash := md5.Sum(fileData)
	foundHash := md5.Sum(fileRespond.Data)
	assert.Equal(t, hash, foundHash)
}
