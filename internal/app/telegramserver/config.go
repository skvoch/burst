package telegramserver

// Config ...
type Config struct {
	ApplicationToken string `toml:"application_token"`
	OwnerName        string `toml:"owner_name"`
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		ApplicationToken: "undefined",
		OwnerName:        "undefined",
	}
}
