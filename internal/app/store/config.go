package store

// Config ...
type Config struct {
	DatabaseURL      string `toml:"database_url"`
	DatabaseUser     string `toml:"database_user"`
	DatabasePassword string `toml:"database_password"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{}
}
