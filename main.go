package main

import (
	"log"
	"os"

	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ko-app-lab/household_account_book_linebot/mypkg"

	"github.com/line/line-bot-sdk-go/linebot"
)

var householdActions = []linebot.TemplateAction{
	linebot.NewMessageAction("Say message", "洗濯"),
	linebot.NewMessageAction("Say message", "掃除"),
	linebot.NewMessageAction("Say message", "犬の散歩"),
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
						template := linebot.NewButtonsTemplate("", askDoneHousehold, askTitle, householdActions...)
						bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage(askTitle, template)).Do()
					} else {
						// 上記以外は、不明なメッセージとして返信
						replyUndefineMessage(bot, event)
					}
				}
			}
		}
	})
	router.Run(":" + port)
}

// 不明なメッセージを取得したときの返信
var undefinedMessage = []string{
	"ちょっと何言っているかわかんない...",
	"日本語話してもろてもいいですか？",
	"すみません、よくわかりません",
	"管理者に使い方教えてもらえ",
}

func replyUndefineMessage(bot *linebot.Client, event *linebot.Event) {
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(generateRandamUndefineMessage())).Do()
	if err != nil {
		log.Print(err)
	}
}

func generateRandamUndefineMessage() string {
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn((len(undefinedMessage) - 1))
	return undefinedMessage[randIndex]
}
