package telegramserver

// Config ...
type Config struct {
	ApplicationToken string `toml:"application_token"`
	OwnerID          int    `toml:"owner_id"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		ApplicationToken: "undefined",
		OwnerID:          0,
	}
}
