package main

import (
	"os"

	config "github.com/Totus-Floreo/tgbotKulakova/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv(config.TELEGRAM_APITOKEN))
	if err != nil {
		panic(err)
	}

	bot.Debug = true
}
