package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	myMessage "github.com/ko-app-lab/household_account_book_linebot/my-message"
	"github.com/ko-app-lab/household_account_book_linebot/mypkg"

	"github.com/line/line-bot-sdk-go/linebot"
)

var householdActions = []linebot.TemplateAction{
	linebot.NewMessageAction("洗濯", ",洗濯"),
	linebot.NewMessageAction("掃除", ",掃除"),
	linebot.NewMessageAction("犬の散歩", ",犬の散歩"),
}

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
						// ログインできたら家事の選択
						askTitle := "家事選択"
						// ログインメッセージと家事選択を促す
						askDoneHousehold := loginMessage + "\n終わった家事を選択してね！"
						// 「ユーザ名,家事名」の形で送信させる
						template := linebot.NewButtonsTemplate("", askDoneHousehold, askTitle, householdActions...)
						bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage(askTitle, template)).Do()
					} else {
						// 上記以外は、不明なメッセージとして返信
						myMessage.ReplyUndefined(bot, event)
					}
					// else if checkRegisterMessage(replyMessage) {
					// 	name := splitMessages(replyMessage)[0]
					// 	point, err := mypkg.UpdatePoint(name)
					// 	if err == nil {
					// 		pointMessage := name + "の家事ポイントは" + strconv.Itoa(point) + "だよ！"
					// 		_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(pointMessage)).Do()
					// 		if err != nil {
					// 			log.Print(err)
					// 		}
					// 	}

					// }

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
	// 家事名が一つでも一致すればtrue
	for _, kind := range houseHoldKinds {
		if kind == splitMessages[1] {
			return true
		}
	}
	return false
}

func splitMessages(message string) []string {
	return strings.Split(message, ",")
}
