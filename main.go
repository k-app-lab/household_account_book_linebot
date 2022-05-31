package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	mydb "github.com/ko-app-lab/household_account_book_linebot/my-database"
	myMessage "github.com/ko-app-lab/household_account_book_linebot/my-message"
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

	// 登録されているユーザ名を取得する
	users, err := mydb.FetchUserName()
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		fmt.Println(user)
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
					// ログイン時の返信
					myMessage.ReplyLogin(bot, event, replyMessage)
					// 「ユーザ名,操作」の形かチェック
					if checkRegisterMessage(replyMessage) {
						name := splitMessages(replyMessage)[0]
						point, err := mypkg.UpdatePoint(name)
						if err == nil {
							pointMessage := name + "の家事ポイントは" + strconv.Itoa(point) + "ptだよ！"
							_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(pointMessage)).Do()
							if err != nil {
								log.Print(err)
							}
						}

					} else {
						// 上記以外は、不明なメッセージとして返信
						myMessage.ReplyUndefined(bot, event)
					}
				}
			}
		}
	})
	router.Run(":" + port)
}

var houseHoldKinds = []string{
	"洗濯",
	"掃除",
	"犬の散歩",
}

func checkRegisterMessage(message string) bool {
	splitMessages := splitMessages(message)
	// 二分割以外は不明なメッセージ
	if len(splitMessages) != 2 {
		return false
	}
	return false
}

func splitMessages(message string) []string {
	return strings.Split(message, ",")
}
