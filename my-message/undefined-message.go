package myMessage

import (
	"log"
	"math/rand"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

// 不明なメッセージを取得したときの返信
var undefinedMessage = []string{
	"ちょっと何言っているかわかんない...",
	"日本語話してもろてもいいですか？",
	"すみません、よくわかりません",
	"管理者に使い方教えてもらえ",
}

func ReplyUndefined(bot *linebot.Client, event *linebot.Event) {
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
