package telegramserver

type Config struct {
	ApplicationToken string `toml:"application_token"`
}

func NewConfig() *Config {
	return &Config{
		ApplicationToken: "undefined",
	}
}
