package main

import (
	"github.com/Mewbi/Misaki-Bot/bot"
	"github.com/Mewbi/Misaki-Bot/config"
)

func main() {
	config.Load()
	bot.Start()
}
