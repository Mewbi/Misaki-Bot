package telegram

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Mewbi/Misaki-Bot/config"
)

type TelegramMessage struct {
	Text string
}

func Listen() ([]TelegramMessage, error) {
	var messages []TelegramMessage
	offset := "-1"
	url := config.Conf.Bot.Path + config.Conf.Bot.Token + "/getUpdates"
	resp, err := http.Get(url + "?offset=" + offset)
	if err != nil {
		return messages, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return messages, err
	}

	sb := string(body)
	fmt.Println(sb)

	return messages, nil
}
