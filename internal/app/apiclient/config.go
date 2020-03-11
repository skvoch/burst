package apiclient

// Config - has only Burst server addr
type Config struct {
	ServerAddr string `toml:"server_addr"`
}

// NewConfig - helper function
func NewConfig() *Config {
	return &Config{}
}
