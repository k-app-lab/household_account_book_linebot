package main

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// LINE bot クライアント生成
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	// エラー時はログ出力
	if err != nil {
		log.Fatal(err)
	}

	// テキストメッセージを生成
	message := linebot.NewTextMessage("うんちまん参上！")
	// テキストメッセージを友達登録しているユーザ全員に配信する
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
