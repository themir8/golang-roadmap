package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	token = "1616952945:AAFcXHr7oYqhyXC1-eVkw2ZjzHn-mt83vAo"
)

type username struct {
	x, y string
}

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if !update.Message.IsCommand() {
			MessageHandler(update, bot)
		} else if update.Message.IsCommand() {
			CommandHandler(update, bot)
		}

		log.Println(username{update.Message.From.UserName, update.Message.Text})
	}
}

func CommandHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	// Extract the command from the Message.
	switch update.Message.Command() {
	case "SocialMedia":
		msg.Text = os.Getenv("SocialMedia")
	case "sayhi":
		msg.Text = "Hi :)"
	case "status":
		msg.Text = "I'm ok."
	default:
		msg.Text = "I don't know that command"
	}

	bot.Send(msg)
}

func MessageHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch {
	case update.Message.Text == "I'm Mirsaid":
		msg.Text = "Hey Boss"
	default:
		msg.Text = "Hey User loser ;)"
	}

	bot.Send(msg)
}
