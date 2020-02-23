package apiclient

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	url := b.serverAddr + "/ping/"
	r, err := b.client.Get(url)

	if r.StatusCode != http.StatusOK {
		return &UnableConnectToServer{}
	}

	if err != nil {
		return err
	}

	return nil
}

func (b *BurstClient) getBookIDs(typeID int) ([]int, error) {

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
