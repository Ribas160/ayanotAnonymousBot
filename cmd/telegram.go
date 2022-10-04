package main

import (
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
}

func (b *Bot) Run() error {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		return err
	}

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			var msg tgbotapi.MessageConfig

			if update.Message.Text == "/start" {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Отправь мне любое сообщение и я перешлю его в анонимный канал.")
				msg.ReplyToMessageID = update.Message.MessageID

			} else {
				channelId, err := strconv.ParseInt(os.Getenv("CHANNEL_ID"), 10, 64)
				if err != nil {
					return err
				}

				msg = tgbotapi.NewMessage(channelId, update.Message.Text)
			}

			bot.Send(msg)
		}
	}

	return nil
}
