package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var (
	config  *Config
	baseUrl string
)

func init() {
	err := godotenv.Overload(".env", ".env.local")

	if err != nil {
		panic("Cant load environment variables")
	}

	config = &Config{
		ApiKey: os.Getenv("BOT_API_KEY"),
	}

	baseUrl = "https://api.telegram.org/bot" + config.ApiKey

}

func main() {

	offset := 0

	for {
		time.Sleep(time.Duration(1) * time.Second)
		updates, err := getUpdates(offset)
		if err != nil {
			log.Println(err.Error())
		}

		for _, update := range updates {
			offset = update.UpdateId
			fmt.Println(update.Message)
		}

	}
}
