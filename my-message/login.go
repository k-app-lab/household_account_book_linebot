package myMessage

import (
	"github.com/ko-app-lab/household_account_book_linebot/mypkg"
	"github.com/line/line-bot-sdk-go/linebot"
)

// ユーザ名でログインしたときに終わった家事を選択させる
func ReplyLogin(bot *linebot.Client, event *linebot.Event, replyMessage string) {
	loginMessage, err := mypkg.FetchLoginMessage(replyMessage)
	if loginMessage == "" || err != nil {
		return
	}
	// ログインできたら家事の選択
	askTitle := "家事選択"
	// ログインメッセージと家事選択を促す
	askDoneHousehold := loginMessage + "\n終わった家事を選択してね！"
	// 「ユーザ名,家事名」の形で送信させる
	var householdActions = []linebot.TemplateAction{
		linebot.NewMessageAction("洗濯", replyMessage+",洗濯"),
		linebot.NewMessageAction("掃除", replyMessage+",掃除"),
		linebot.NewMessageAction("犬の散歩", replyMessage+",犬の散歩"),
	}
	template := linebot.NewButtonsTemplate("", askDoneHousehold, askTitle, householdActions...)
	bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage(askTitle, template)).Do()
}
