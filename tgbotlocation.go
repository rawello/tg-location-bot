package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	var TELEGRAM_API_TOKEN string = "your token"

	var resp string
	var chatID int64

	bot, err := tgbotapi.NewBotAPI(TELEGRAM_API_TOKEN)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s (https://t.me/%s)", bot.Self.UserName, bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID = update.Message.Chat.ID
		t := update.Message.Text
		log.Printf("[%s] %s (command: %v, location: %v)\n", update.Message.From.UserName, t, update.Message.IsCommand(), update.Message.Location)
		switch {
		case update.Message.IsCommand():
			switch update.Message.Command() {
			case "start", "help":
				bot.Send(tgbotapi.NewMessage(chatID, "🤖 я живое\n\nкароч я стартую."))
			case "alert":
				bot.Send(tgbotapi.NewMessage(chatID, "🚨 ктоя 🤔"))
			case "sosi":
				bot.Send(tgbotapi.NewMessage(chatID, "🤔 сам саси"))
			default:
				bot.Send(tgbotapi.NewMessage(chatID, "🤔 непон"))
			}
		case update.Message.Location != nil:
			resp = "я выезжаю, жди меня 😸"
			msg := tgbotapi.NewMessage(chatID, resp)
			if _, e := bot.Send(msg); e != nil {
				log.Printf("Error sending message to telegram.\nMessage: %v\nError: %v", msg, e)
			}
		default:
			resp = fmt.Sprintf("\"%v\" ща как жахну 💥", t)
			msg := tgbotapi.NewMessage(chatID, resp)

			if _, e := bot.Send(msg); e != nil {
				log.Printf("Error sending message to telegram.\nMessage: %v\nError: %v", msg, e)
			}
		}

	}
}