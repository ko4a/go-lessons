package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var (
	config  *Config
	baseUrl string
	db      *Store
)

func init() {
	err := godotenv.Overload(".env", ".env.local")

	if err != nil {
		panic("Cant load environment variables")
	}

	config = &Config{
		ApiKey:   os.Getenv("BOT_API_KEY"),
		dbConfig: NewDbConfig(os.Getenv("DB_URL")),
		LogUrl:   os.Getenv("LOG_URL"),
	}

	db := NewStore(config.dbConfig)

	if err = db.Open(); err != nil {
		panic(err.Error())
	}

	baseUrl = "https://api.telegram.org/bot" + config.ApiKey

	outFile, _ = os.OpenFile(config.LogUrl, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	LogFile = log.New(outFile, "", 0)
	LogFile.SetPrefix(time.Now().String())

}

func main() {

	offset := 0

	for {
		time.Sleep(time.Duration(500) * time.Microsecond)

		updates, err := getUpdates(offset)
		WriteErrorLog(err)

		for _, update := range updates {
			WriteUpdateLog(update)
			go update.process()
			offset = update.UpdateId + 1
		}
	}
}
