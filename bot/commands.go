package bot

import (
	"log"
	"regexp"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func compileRegex(s string) *regexp.Regexp {
    r, err := regexp.Compile(s)
    if err != nil {
        log.Fatal(err)
    }
    return r
}

func CommandList() []Command {
    return []Command{
        Command{Pattern: compileRegex(`^help$`), Action: HelpMenu},
        Command{Pattern: compileRegex(`^help2$`), Action: HelpMenu2},
    }
}

func HelpMenu(m MisakiBot) error {
    msg := tg.NewMessage(m.Msg.Chat.ID, "É o menu de ajuda porra!")
    msg.ReplyToMessageID = m.Msg.MessageID

    if _, err := m.Bot.Send(msg); err != nil {
        log.Printf("Error sending message: %v", err)
    }
    return nil
}

func HelpMenu2(m MisakiBot) error {
    msg := tg.NewMessage(m.Msg.Chat.ID, "É o menu mais foda de teste!")
    msg.ReplyToMessageID = m.Msg.MessageID

    if _, err := m.Bot.Send(msg); err != nil {
        log.Printf("Error sending message: %v", err)
    }
    return nil
}
