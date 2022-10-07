package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

type App struct {
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	bot := &Bot{}

	bot.Run()
}

func (a *App) LogError(errorMsg string) {
	logger := log.New(os.Stderr, "[ERROR] ", log.Ltime|log.Lshortfile)

	newPath := filepath.Join(".", "logs")

	err := os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		logger.Println(err.Error())
	}

	currentTime := time.Now()
	currentDate := currentTime.Format("01-02-2006")

	fp, err := os.OpenFile(fmt.Sprintf("%s/%s", newPath, currentDate+".log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logger.Println(err.Error())
	}

	defer fp.Close()

	logger.SetOutput(fp)

	logger.Println(errorMsg)
}
