package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/skvoch/burst/internal/app/telegramserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "telegram.toml", "path to config file")
}

// Override config fields if env variables is set
func setEnviromentVaribles(c *telegramserver.Config) error {
	envApplicationToken := os.Getenv("TELEGRAM_APPLICATION_TOKEN")

	if envApplicationToken != "" {
		c.ApplicationToken = envApplicationToken
	}

	envOwnerID := os.Getenv("OWNER_ID")

	if envOwnerID != "" {
		id, err := strconv.Atoi(envOwnerID)

		if err != nil {
			return err
		}

		c.OwnerID = id
	}

	return nil
}

func main() {
	flag.Parse()

	config := telegramserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err := setEnviromentVaribles(config); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	s, err := telegramserver.New(config)

	if err != nil {
		log.Fatal(err)
	}
	s.SetupHandlers()
	s.Start()
}
