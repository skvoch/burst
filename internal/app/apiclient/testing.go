package apiclient

// NewTestClient ...
func NewTestClient(URL string) (*BurstClient, error) {
	config := NewConfig()
	config.ServerAddr = "http://localhost:8080"
	client, err := New(config)

	if err != nil {
		return nil, err
	}

	return client, nil
}
