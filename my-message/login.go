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
		return false
	}
	userName := replyMessage
	// ログインメッセージとやりたいことの選択を促す
	title := "やりたいこと"
	wantToDoMessage := loginMessage + "\nやりたいことを選択してね！"
	var toDoActions = []linebot.TemplateAction{
		linebot.NewMessageAction("終わった家事を登録", userName+",終了した家事登録"),
		linebot.NewMessageAction("みんなの家事ポイント確認", "家事ポイント確認"),
	}
	template := linebot.NewButtonsTemplate("", wantToDoMessage, title, toDoActions...)
	bot.ReplyMessage(event.ReplyToken, linebot.NewTemplateMessage(title, template)).Do()
	return true
}
