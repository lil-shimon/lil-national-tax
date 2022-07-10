package main

import (
	"github.com/lil-shimon/national-tax/message"
	"github.com/lil-shimon/national-tax/question"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
)

func main() {
	//err := godotenv.Load(".env")
	//if err != nil {
	//	log.Fatal(err)
	//}

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

	//txt = linebot.NewTextMessage(qs.Answer)
	//if _, err := bot.BroadcastMessage(txt).Do(); err != nil {
	//	log.Fatal(err)
	//}

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(qs.Answer)).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	})
}
