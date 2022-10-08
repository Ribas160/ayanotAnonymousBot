package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/Ribas160/ayanotAnonymousBot/pkg/bot"
	"github.com/spf13/viper"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env variables: %s", err.Error())
	}

	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config %s", err.Error())
	}

	viper.WatchConfig()

	bot.Run()
}
