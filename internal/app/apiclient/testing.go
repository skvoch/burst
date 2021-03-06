package apiclient

// NewTestClient - helper function for testing
func NewTestClient(URL string) (*BurstClient, error) {
	config := NewConfig()
	config.ServerAddr = "http://localhost:8080"
	client, err := New(config.ServerAddr)

	if err != nil {
		return nil, err
	}

	return client, nil
}
