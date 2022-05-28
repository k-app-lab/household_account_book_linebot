package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ko-app-lab/household_account_book_linebot/mypkg"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	router.Use(gin.Logger())

	// LINE Messaging API ルーティング
	router.POST("/callback", func(c *gin.Context) {
		// Lineでメッセージを取得
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				log.Print(err)
			}
			return
		}

		for _, event := range events {
			// イベントがメッセージの受信だった場合
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				// メッセージがテキスト形式の場合
				case *linebot.TextMessage:
					replyMessage := message.Text
					loginMessage, err := mypkg.FetchLoginMessage(replyMessage)
					if loginMessage != "" && err == nil {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(loginMessage)).Do()
					} else {
						// 上記意外は、おうむ返しで返信
						_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
						if err != nil {
							log.Print(err)
						}
					}
				}
			}
		}
	})
	router.Run(":" + port)
}
