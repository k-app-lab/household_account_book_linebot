package myMessage

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

// ユーザ名でログインしたときに終わった家事を選択させる
func ReplyHouseholdRegister(bot *linebot.Client, event *linebot.Event, name string) {
	// ログインできたら家事の選択
	askTitle := "家事選択"
	// ログインメッセージと家事選択を促す
	askDoneHousehold := name + "\n終わった家事を選択してね！"
	// 「ユーザ名,家事名」の形で送信させる
	var householdActions = []linebot.TemplateAction{
		linebot.NewMessageAction("洗濯", name+",洗濯"),
		linebot.NewMessageAction("掃除", name+",掃除"),
		linebot.NewMessageAction("犬の散歩", name+",犬の散歩"),
	}
	template := linebot.NewButtonsTemplate("", askDoneHousehold, askTitle, householdActions...)
	bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage(askTitle, template)).Do()
}
