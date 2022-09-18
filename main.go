package main

import (
	"github.com/Mewbi/Misaki-Bot/bot"
	"github.com/Mewbi/Misaki-Bot/conf"
)

func main() {
	conf.Load()
	bot.Start()
}
