package main

import (
	"github.com/joho/godotenv"
	"github.com/lil-shimon/national-tax/message"
	"github.com/lil-shimon/national-tax/question"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	qs := question.GetQuestion()
	msg := message.FmtMsg(qs)

	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	txt := linebot.NewTextMessage(msg)
	if _, err := bot.BroadcastMessage(txt).Do(); err != nil {
		log.Fatal(err)
	}

	txt = linebot.NewTextMessage(qs.Answer)
	if _, err := bot.BroadcastMessage(txt).Do(); err != nil {
		log.Fatal(err)
	}
}
