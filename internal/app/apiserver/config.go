package apiserver

// Config - has fields for Burst server
type Config struct {
	BindAddr          string `toml:"bind_addr"`
	LogLevel          string `toml:"log_level"`
	DataBaseURL       string `toml:"database_url"`
	FilesDirectory    string `toml:"files_directory"`
	PreviewsDirectory string `toml:"previews_directory"`
}

// NewConfig - helper function
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
