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

	qs := question.GetQuestion()

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	print(qs.Question)

	message := linebot.NewTextMessage("国税試験問題\n" + qs.Question)
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}

	message = linebot.NewTextMessage(qs.Answer)
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
