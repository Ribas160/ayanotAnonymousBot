package bot

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

func Run() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))

	if err != nil {
		errorLog(err.Error())
		return
	}

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			if update.Message.Text == "/start" {
				err := startMessage(bot, update)

				if err != nil {
					errorLog(err.Error())
				}

			} else {
				err := copyMessageToChannel(bot, update)
				if err != nil {
					errorLog(err.Error())
				}
			}
		}
	}
}

func startMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	var msg tgbotapi.MessageConfig

	msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Отправь мне любое сообщение и я перешлю его в анонимный канал.")
	msg.ReplyToMessageID = update.Message.MessageID

	_, err := bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func copyMessageToChannel(bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	var msg tgbotapi.CopyMessageConfig

	channelId, err := strconv.ParseInt(os.Getenv("CHANNEL_ID"), 10, 64)
	if err != nil {
		return err
	}

	if !filter(update) {
		return nil
	}

	msg = tgbotapi.NewCopyMessage(channelId, update.Message.Chat.ID, update.Message.MessageID)

	_, err = bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func errorLog(errorMsg string) {
	logger := log.New(os.Stderr, "[ERROR] ", log.Ltime|log.Lshortfile)

	newPath := filepath.Join(".", "logs")

	err := os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		logger.Println(err.Error())
	}

	currentTime := time.Now()
	currentDate := currentTime.Format("02-01-2006")

	fp, err := os.OpenFile(fmt.Sprintf("%s/%s", newPath, currentDate+".log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logger.Println(err.Error())
	}

	defer fp.Close()

	logger.SetOutput(fp)

	logger.Println(errorMsg)
}

func filter(update tgbotapi.Update) bool {

	if update.Message.Text != "" && !viper.GetBool("enableTextMessages") {
		return false
	}

	if update.Message.Voice != nil && !viper.GetBool("enableVoiceMessages") {
		return false
	}

	if update.Message.Photo != nil && !viper.GetBool("enableImages") {
		return false
	}

	if update.Message.Sticker != nil && !viper.GetBool("enableStickers") {
		return false
	}

	if update.Message.Animation != nil && !viper.GetBool("enableAnimation") {
		return false
	}

	if update.Message.Video != nil && !viper.GetBool("enableVideos") {
		return false
	}

	if update.Message.Audio != nil && !viper.GetBool("enableAudios") {
		return false
	}

	if update.Message.Document != nil && !viper.GetBool("enableFiles") {
		return false
	}

	if update.Message.Game != nil && !viper.GetBool("enableGames") {
		return false
	}

	return true
}
