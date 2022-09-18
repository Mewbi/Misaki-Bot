package bot

import (
	"log"
	"regexp"

	"github.com/Mewbi/Misaki-Bot/conf"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MisakiBot struct {
    Bot *tg.BotAPI
    Msg *tg.Message
    Commands []Command
}

type Command struct {
    Pattern *regexp.Regexp
    Action func(MisakiBot) error
}

func (m MisakiBot) Process() {
    if !m.Msg.IsCommand()  {
        return
    }
    msg := m.Msg.Command()
    log.Println(msg) // Just to debug
    for _, command := range m.Commands {
        if !command.Pattern.MatchString(msg) {
            continue
        }
        if err := command.Action(m); err != nil {
            log.Panic(err)
        }
    }

}

func Start() {
    var misaki MisakiBot
    config := conf.Get()
	bot, err := tg.NewBotAPI(config.Bot.Token)
	if err != nil {
		log.Panic(err)
	}
    misaki.Bot = bot
    misaki.Bot.Debug = config.Debug

    misaki.Commands = CommandList()

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tg.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := misaki.Bot.GetUpdatesChan(updateConfig)

	for update := range updates {
        if update.Message == nil  {
            continue
        }
        misaki.Msg = update.Message
        misaki.Process()
	}
}

