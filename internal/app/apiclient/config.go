package apiclient

// Config ...
type Config struct {
	ServerAddr string `toml:"server_addr"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
