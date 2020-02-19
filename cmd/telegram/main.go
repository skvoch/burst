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
	flag.StringVar(&configPath, "config-path", "configs/telegram.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := telegramserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	server := telegramserver.New(config)
	server.Start()
}
