package main

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/skvoch/burst/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func setEnviromentVaribles(c *apiserver.Config) {
	envDatabaseURL := os.Getenv("DATABASE_URL")

	if envDatabaseURL != "" {
		c.DataBaseURL = envDatabaseURL
	}
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	setEnviromentVaribles(config)

	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(config); err != nil {
		log.Fatal(err)
	}

	return
}
