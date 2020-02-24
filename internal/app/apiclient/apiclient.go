package apiclient

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/skvoch/burst/internal/app/model"
)

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

func (b *BurstClient) ping() error {
	url := b.serverAddr + "/v1.0/ping/"
	r, err := b.client.Get(url)

	if r.StatusCode != http.StatusOK {
		return &UnableConnectToServer{}
	}

	if err != nil {
		return err
	}

	return nil
}

func (b *BurstClient) GetBookIDs(typeID int) ([]int, error) {

	type Response struct {
		BooksIDs []int `json:"books_ids"`
	}

	id := strconv.Itoa(typeID)

	url := b.serverAddr + "/v1.0/types/" + id + "/books/"
	r, err := b.client.Get(url)

	if err != nil {
		return nil, err
	}
	respData := &Response{}
	json.NewDecoder(r.Body).Decode(respData)

	return respData.BooksIDs, nil
}

func (b *BurstClient) GetAllTypes() ([]*model.Type, error) {

	url := b.serverAddr + "/v1.0/types/"
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

func (b *BurstClient) CreateType(_type *model.Type) error {
	url := b.serverAddr + "/v1.0/types/create/"

	data := make([]byte, 0)
	if err := json.Unmarshal(data, _type); err != nil {
		return err
	}
	reader := bytes.NewReader(data)

	res, err := b.client.Post(url, "application/json", reader)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusCreated {
		return &WrongResponseStatus{}
	}

	return nil
}
