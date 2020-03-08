package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/skvoch/burst/internal/app/telegramserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "telegram.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := telegramserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

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
