package main

import (
	"github.com/lil-shimon/national-tax/question"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
)

func main() {
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatal(err)
	//}

	question.GetQuestion()

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage("Hello, world!")
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}

}
