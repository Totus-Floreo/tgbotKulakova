package main

import (
	"fmt"

	"github.com/Totus-Floreo/tgbotKulakova/config"
	logging "github.com/Totus-Floreo/tgbotKulakova/utils/logging"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var AllowedUpdates = []string{
	"message",
	"edited_message",
	"my_chat_member",
	"callback_query",
	"chat_member",
	"channel_post",
	"edited_channel_post",
	// ! unused updates
	// "inline_query",
	// "chosen_inline_result",
	// "shipping_query",
	// "pre_checkout_query",
	// "poll",
	// "poll_answer",
	// "chat_join_request",
}

func main() {
	fmt.Println("Starting...")
	logger := logging.Initialize()
	defer logger.Sync()
	logger.Info("Starting the application...")
	bot, err := tgbotapi.NewBotAPI(config.TELEGRAM_APITOKEN)
	if err != nil {
		logger.DPanic(err.Error())
		panic(err)
	}
	bot.Debug = true

	defer bot.Send(tgbotapi.NewMessage(config.Owner_ID, "Im dead"))
	if _, err := bot.Send(tgbotapi.NewMessage(config.Owner_ID, "Im live")); err != nil {
		logger.DPanic(err.Error())
		panic(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updateConfig.AllowedUpdates = AllowedUpdates
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		go func(u tgbotapi.Update) {
			var answer tgbotapi.MessageConfig
			switch u.FromChat().Type {
			case "private":
				if u.Message == nil || !u.Message.IsCommand() {
					return
				}
				answer = msgInChat(u)
			case "channel":
				if u.ChannelPost == nil || !u.ChannelPost.IsCommand() {
					return
				}
				answer = pstInChan(u)
			case "supergroup":
				if u.Message == nil || !u.Message.IsCommand() {
					return
				}
				answer = msgInGroup(u)
			}
			if _, err := bot.Send(answer); err != nil {
				logger.DPanic(err.Error())
				panic(err)
			}
		}(update)
	}
}

func msgInChat(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "done")
	msg.ReplyToMessageID = update.Message.MessageID
	return msg
}

func pstInChan(update tgbotapi.Update) tgbotapi.MessageConfig {
	channelName := fmt.Sprintf("@%s", update.ChannelPost.SenderChat.UserName)
	pst := tgbotapi.NewMessageToChannel(channelName, "done")
	return pst
}

func msgInGroup(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "done")
	msg.ReplyToMessageID = update.Message.MessageID
	return msg
}
