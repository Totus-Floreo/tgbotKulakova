package main

import (
	"fmt"

	config "github.com/Totus-Floreo/tgbotKulakova/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	fmt.Println("Starting...")
	bot, err := tgbotapi.NewBotAPI(config.TELEGRAM_APITOKEN)
	if err != nil {
		panic(err)
	}
	fmt.Println("Working...")
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.ChosenInlineResult == nil {
			continue
		}
		switch update.FromChat().Type {
		case "private":
			if update.Message == nil {
				continue
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		case "channel":
			if  {
				continue
			}
			channelName := fmt.Sprintf("@%s", update.ChannelPost.SenderChat.UserName)
			pst := tgbotapi.NewMessageToChannel(channelName, update.ChannelPost.Text)
			if _, err := bot.Send(pst); err != nil {
				panic(err)
			}
		case "supergroup":
			if update.Message == nil {
				continue
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			if _, err := bot.Send(msg); err != nil {
				panic(err)
			}
		}
	}
	fmt.Println("Stoping...")
}

//pm

// Endpoint: getUpdates, response: {"ok":true,"result":[{"update_id":115121081,
// "message":{"message_id":54,"from":{"id":214659520,"is_bot":false,"first_name":"Timur","last_name":"Kulakov","username":"hietotusfloreo","language_code":"en"},
// "chat":{"id":214659520,"first_name":"Timur","last_name":"Kulakov","username":"hietotusfloreo","type":"private"},"date":1680548947,"text":"12"}}]}

//ch

// Endpoint: getUpdates, response: {"ok":true,"result":[{"update_id":115121083,
// "channel_post":{"message_id":29,"sender_chat":{"id":-1001961815317,"title":"testChannel","type":"channel"},
// "chat":{"id":-1001961815317,"title":"testChannel","type":"channel"},"date":1680549050,"text":"12"}}]}

//gr

// Endpoint: getUpdates, response: {"ok":true,"result":[{"update_id":115121132,
// "message":{"message_id":6,"from":{"id":214659520,"is_bot":false,"first_name":"Timur","last_name":"Kulakov","username":"hietotusfloreo","language_code":"en"},
// "chat":{"id":-1001906889117,"title":"testGroup","type":"supergroup"},"date":1680553456,"text":"1"}}]}
