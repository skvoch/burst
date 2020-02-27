package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/skvoch/burst/internal/app/model"
)

// BurstClient ...
type BurstClient struct {
	serverAddr string
	client     *http.Client
}

// New ...
func New(config *Config) (*BurstClient, error) {
	client := &BurstClient{
		serverAddr: config.ServerAddr,
		client:     &http.Client{},
	}

	if err := client.ping(); err != nil {
		return nil, err
	}

	return client, nil
}

func (b *BurstClient) makeURL(endpoint string) string {
	return b.serverAddr + endpoint
}

func (b *BurstClient) ping() error {
	url := b.serverAddr + "/v1.0/ping/"
	r, err := b.client.Get(url)

	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return &UnableConnectToServer{}
	}

	return nil
}

// GetBookIDs ...
func (b *BurstClient) GetBookIDs(typeID int) ([]int, error) {

	id := strconv.Itoa(typeID)

	url := b.makeURL("/v1.0/types/" + id + "/books/")
	r, err := b.client.Get(url)

	if err != nil {
		return nil, err
	}

	respData := &GetBooksIDsResponse{}
	json.NewDecoder(r.Body).Decode(respData)

	return respData.BooksIDs, nil
}

// RemoveAllTypes ...
func (b *BurstClient) RemoveAllTypes() error {
	url := b.makeURL("/v1.0/types/")

	rec, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return err
	}

	r, err := b.client.Do(rec)

	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return &WrongResponseStatus{}
	}

	return nil
}

// RemoveAllBooks ...
func (b *BurstClient) RemoveAllBooks() error {
	url := b.makeURL("/v1.0/books/remove/")

	rec, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return err
	}

	r, err := b.client.Do(rec)

	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return &WrongResponseStatus{}
	}

	return nil
}

// GetAllTypes ...
func (b *BurstClient) GetAllTypes() ([]*model.Type, error) {

	url := b.makeURL("/v1.0/types/")
	r, err := b.client.Get(url)

	if err != nil {
		return nil, err
	}

	types := make([]*model.Type, 0)
	if err := json.NewDecoder(r.Body).Decode(&types); err != nil {
		return nil, err
	}

	return types, nil
}

// CreateType ...
func (b *BurstClient) CreateType(_type *model.Type) (int, error) {
	url := b.makeURL("/v1.0/types/create/")

	data := make([]byte, 0)
	data, err := json.Marshal(_type)

	if err != nil {
		return -1, err
	}

	reader := bytes.NewReader(data)

	res, err := b.client.Post(url, "application/json", reader)

	if err != nil {
		return -1, err
	}

	if res.StatusCode != http.StatusCreated {
		return -1, &WrongResponseStatus{}
	}

	response := &CreateTypeResponse{}
	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return -1, err
	}

	return response.ID, nil
}

// CreateBook ...
func (b *BurstClient) CreateBook(book *model.Book) (*BookUploadTokens, error) {
	url := b.makeURL("/v1.0/books/create/")

	data, err := json.Marshal(book)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(data)

	res, err := b.client.Post(url, "application/json", reader)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusCreated {
		return nil, &WrongResponseStatus{}
	}

	tokens := &BookUploadTokens{}
	if err := json.NewDecoder(res.Body).Decode(tokens); err != nil {
		return nil, err
	}

	return tokens, nil
}

// GetBookByID ...
func (b *BurstClient) GetBookByID(ID int) (*model.Book, error) {
	url := b.makeURL("/v1.0/books/" + strconv.Itoa(ID) + "/")

	book := &model.Book{}
	res, err := b.client.Get(url)

	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(book); err != nil {
		return nil, err
	}

	return book, nil
}

// GetBookPreview ...
func (b *BurstClient) GetBookPreview(ID int) ([]byte, error) {
	url := b.makeURL("/v1.0/books/" + strconv.Itoa(ID) + "/preview/")

	res, err := b.client.Get(url)

	if err != nil {
		return nil, err
	}
	fmt.Println(res.Body)

	return nil, nil
}

// SendBookFile ...
func (b *BurstClient) SendBookFile(data []byte, bookID int, UUID string) error {
	url := b.makeURL("/v1.0/books/" + strconv.Itoa(bookID) + "/preview/")

	reader := bytes.NewReader(data)
	_, err := http.Post(url, "binary/octet-stream", reader)

	if err != nil {
		return err
	}

	return nil
}

// GetBookFile ...
func (b *BurstClient) GetBookFile(ID int) ([]byte, error) {
	url := b.makeURL("/v1.0/books/" + strconv.Itoa(ID) + "/file/")

	res, err := b.client.Get(url)

	if err != nil {
		return nil, err
	}

	file, _, err := res.Request.FormFile("file")

	result := make([]byte, 0)
	_, err = file.Read(result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// SendPreviewFile ...
func (b *BurstClient) SendPreviewFile(filePath string, name string, bookID int, UUID string) error {
	url := b.makeURL("/v1.0/books/" + strconv.Itoa(bookID) + "/preview/")

	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		return err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("preview", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Token-UUID", UUID)

	client := &http.Client{}
	client.Do(req)

	return nil
}
