package bot

import (
	"fmt"
	"log"

	"github.com/Mewbi/Misaki-Bot/telegram"
)

func Start() {
	for {
		messages, err := telegram.Listen()
		if err != nil {
			log.Print(err)
			continue
		}
		for _, message := range messages {
			fmt.Println(message.Text)
		}
	}
}
