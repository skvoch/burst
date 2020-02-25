package apiclient

type Config struct {
	ServerAddr string `toml:"server_addr"`
}

func NewConfig() *Config {
	return &Config{}
}
