package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	bot := NewBot()
	err := bot.Run()

	if err != nil {
		bot.ErrorLog.Fatalln(err.Error())
	}
}
