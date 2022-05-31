package myMessage

import (
	"github.com/ko-app-lab/household_account_book_linebot/mypkg"
	"github.com/line/line-bot-sdk-go/linebot"
)

// ログイン時に行いたいことを選択させる
func ReplyLogin(bot *linebot.Client, event *linebot.Event, replyMessage string) bool {
	loginMessage, err := mypkg.FetchLoginMessage(replyMessage)
	if loginMessage == "" || err != nil {
		// 登録されていないユーザ名が来た
		bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("登録されていないユーザだよ")).Do()
		return false
	}
	// ログインメッセージとやりたいことの選択を促す
	// askDoneHousehold := loginMessage + "\nやりたいことを選択してね！"
	// var operationActions = []linebot.TemplateAction{
	// 	linebot.NewMessageAction("終わった家事を登録", replyMessage+",終了した家事登録"),
	// 	linebot.NewMessageAction("みんなの家事ポイント確認", "家事ポイント確認"),
	// }
	// template := linebot.NewButtonsTemplate("", askDoneHousehold, "", operationActions...)
	// bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage("", template)).Do()
	bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(loginMessage)).Do()
	return true
}
