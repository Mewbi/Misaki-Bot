package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type BotTelegram struct {
	Token string
	Path  string
}

type Config struct {
	Bot BotTelegram `toml:"bot-telegram"`
}

var Conf Config

func Load() {
	// Read Config File
	data, err := os.ReadFile("config/config.toml")
	if err != nil {
		log.Fatal(err)
	}

	// Parse Config File
	if _, err := toml.Decode(string(data), &Conf); err != nil {
		log.Fatal(err)
	}
}
