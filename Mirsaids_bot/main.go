package main

import (
	"log"

	"github.com/joho/godotenv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	token = "1616952945:AAFcXHr7oYqhyXC1-eVkw2ZjzHn-mt83vAo"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Instagram", "https://instagram.com/mirsaid.m8/"),
		tgbotapi.NewInlineKeyboardButtonURL("Telegram", "https://t.me/Mirzakhidov_M/"),
		tgbotapi.NewInlineKeyboardButtonURL("Github", "https://github.com/mirsaid-mirzohidov/"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // ignore any non-Message Updates
			if !update.Message.IsCommand() {
				MessageHandler(update, bot)
			} else if update.Message.IsCommand() {
				CommandHandler(update, bot)
			}
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			// And finally, send a message containing the data received.
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
}

func CommandHandler(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	// If the message was open, add a copy of our numeric keyboard.
	// Extract the command from the Message.
	switch update.Message.Command() {
	case "SocialMedia":
		msg.ParseMode = "Instagram: <a href='https://www.instagram.com/mirsaid.m8/'>@mirsaid.m8</a> <br>Telegram: @Mirzakhidov_M"
	case "sayhi":
		msg.Text = "Hi :)"
		msg.ReplyMarkup = numericKeyboard
	case "status":
		msg.Text = "I'm ok."
	default:
		msg.Text = "I don't know that command"
	}

	// Send the message.
	if _, err = bot.Send(msg); err != nil {
		panic(err)
	}
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
