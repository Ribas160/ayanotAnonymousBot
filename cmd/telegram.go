package main

import (
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
}

func (b *Bot) Run() {
	app := &App{}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))

	if err != nil {
		app.LogError(err.Error())
		return
	}

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			if update.Message.Text == "/start" {
				err := b.startMessage(bot, update)

				if err != nil {
					app.LogError(err.Error())
				}

			} else {
				err := b.copyMessageToChannel(bot, update)
				if err != nil {
					app.LogError(err.Error())
				}
			}
		}
	}
}

func (b *Bot) startMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	var msg tgbotapi.MessageConfig

	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Отправь мне любое сообщение и я перешлю его в анонимный канал.")
	msg.ReplyToMessageID = update.Message.MessageID

	_, err := bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) copyMessageToChannel(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	var msg tgbotapi.CopyMessageConfig

	channelId, err := strconv.ParseInt(os.Getenv("CHANNEL_ID"), 10, 64)
	if err != nil {
		return err
	}

	msg = tgbotapi.NewCopyMessage(channelId, update.Message.Chat.ID, update.Message.MessageID)

	_, err = bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
