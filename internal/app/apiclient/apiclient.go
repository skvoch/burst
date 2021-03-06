package apiclient

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime"
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

// New - helper function
func New(serverAddr string) (*BurstClient, error) {
	client := &BurstClient{
		serverAddr: serverAddr,
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

// GetBookIDs - get books id by type id
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

// RemoveAllTypes - remove all types with books
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

// RemoveAllBooks - just remove all books
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

// GetAllTypes - getting all types
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

// CreateType - just creating type
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

// CreateBook - just creating book
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

// GetBookByID - geting book by id
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

// GetBookPreview - geting book preview (some image like a JPG, PNG)
func (b *BurstClient) GetBookPreview(ID int) *FileResponse {
	url := b.makeURL("/v1.0/books/" + strconv.Itoa(ID) + "/preview/")

	res, err := b.client.Get(url)

	if err != nil {
		return &FileResponse{Err: err}
	}

	if res.StatusCode == http.StatusInternalServerError {
		return &FileResponse{Err: err}
	}

	_, params, err := mime.ParseMediaType(res.Header.Get("Content-Disposition"))
	fileName := params["filename"]
	bodyBytes, err := ioutil.ReadAll(res.Body)

	return &FileResponse{FileName: fileName, Data: bodyBytes, Err: nil}
}

// SendBookFile - uploading book file (like a PDF), requries upload token (UUID)
func (b *BurstClient) SendBookFile(filePath string, bookID int, UUID string) error {
	url := b.makeURL("/v1.0/books/" + strconv.Itoa(bookID) + "/file/")

	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		return err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Token-UUID", UUID)

	client := &http.Client{}
	client.Do(req)

	return nil
}

// GetBookFile - just getting bool file 
func (b *BurstClient) GetBookFile(ID int) *FileResponse {
	url := b.makeURL("/v1.0/books/" + strconv.Itoa(ID) + "/file/")

	res, err := b.client.Get(url)

	if err != nil {
		return &FileResponse{Err: err}
	}

	if res.StatusCode == http.StatusInternalServerError {
		return &FileResponse{Err: err}
	}

	_, params, err := mime.ParseMediaType(res.Header.Get("Content-Disposition"))
	fileName := params["filename"]
	bodyBytes, err := ioutil.ReadAll(res.Body)

	return &FileResponse{FileName: fileName, Data: bodyBytes, Err: nil}
}

// SendBookFile - uploading book preview (like a JPG, PNG), requries upload token (UUID)
func (b *BurstClient) SendPreview(filePath string, bookID int, UUID string) error {
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
