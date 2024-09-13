package keyboard

import (
	"fmt"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/types"
)

const (
	inKB    string = "Inline Keyboard"
	replyKB string = "Reply Keyboard"
	button  string = "Button"
	inbtn   string = "Inline Button"
	rpbtn   string = "Reply Button"
)

func (in *Inline) Set(plan []int) {
	in.InlineKeyboard = make([][]*InlineKeyboardButton, len(plan))
	for i := range in.InlineKeyboard {
		in.InlineKeyboard[i] = make([]*InlineKeyboardButton, plan[i])
	}
}

func (in *Inline) NewButton(line, pos int) (*InlineKeyboardButton, error) {
	var err error

	if len(in.InlineKeyboard) > line && len(in.InlineKeyboard[line]) > pos {

		if in.InlineKeyboard[line][pos] != nil {
			logs.DataIsntEmply(inKB, fmt.Sprintf("%s line: %d, position: %d", button, line, pos), in.InlineKeyboard[line][pos])
		}
		in.InlineKeyboard[line][pos] = new(InlineKeyboardButton)

		return in.InlineKeyboard[line][pos], nil

	} else {

		if len(in.InlineKeyboard) > line {
			err = errors.ButtosDoesntFit("line", line)
		} else if len(in.InlineKeyboard[line]) > pos {
			err = errors.ButtosDoesntFit("pos", pos)
		}
	}
	return nil, err
}

func (inb *InlineKeyboardButton) WriteString(text string) {
	if inb.Text != "" {
		logs.DataIsntEmply(inbtn, "Text", inb.Text)
	}
	inb.Text = text
	logs.DataWrittenSuccessfully(inbtn, "Text")
}

func (inb *InlineKeyboardButton) WriteURL(url string) {
	if inb.Url != "" {
		logs.DataIsntEmply(inbtn, "URL", inb.Url)
	}
	inb.Url = url
	logs.DataWrittenSuccessfully(inbtn, "URL")
}

func (inb *InlineKeyboardButton) WriteCallbackData(text string) {
	if inb.CallbackData != "" {
		logs.DataIsntEmply(inbtn, "Callback Data", inb.CallbackData)
	}
	inb.CallbackData = text
	logs.DataWrittenSuccessfully(inbtn, "Callback Data")
}

func (inb *InlineKeyboardButton) WriteWebApp(wbapp *types.WebAppInfo) {
	if inb.WebApp != nil {
		logs.DataIsntEmply(inbtn, "Web App", *inb.WebApp)
	}
	inb.WebApp = wbapp
	logs.DataWrittenSuccessfully(inbtn, "Web App")
}

func (inb *InlineKeyboardButton) WriteLoginUrl(logurl *types.LoginUrl) {
	if inb.LoginUrl != nil {
		logs.DataIsntEmply(inbtn, "Login URL", *inb.LoginUrl)
	}
	inb.LoginUrl = logurl
	logs.DataWrittenSuccessfully(inbtn, "Login URL")
}

func (inb *InlineKeyboardButton) WriteSwitchInlineQuery(sw string) {
	if inb.SwitchInlineQuery != "" {
		logs.DataIsntEmply(inbtn, "Switch Inline Query", inb.SwitchInlineQuery)
	}
	inb.SwitchInlineQuery = sw
	logs.DataWrittenSuccessfully(inbtn, "Switch Inline Query")
}

func (inb *InlineKeyboardButton) WriteSwitchInlineQueryCurrentChat(swcch string) {
	if inb.SwitchInlineQueryCurrentChat != "" {
		logs.DataIsntEmply(inbtn, "Switch Inline Query Current Chat", inb.SwitchInlineQueryCurrentChat)
	}
	inb.SwitchInlineQueryCurrentChat = swcch
	logs.DataWrittenSuccessfully(inbtn, "Switch Inline Query Current Chat")
}

func (inb *InlineKeyboardButton) WriteSwitchInlineQueryChosenChat(sw *types.SwitchInlineQueryChosenChat) {
	if inb.SwitchInlineQueryChosenChat != nil {
		logs.DataIsntEmply(inbtn, "Switch Inline Query Chosen Chat", inb.SwitchInlineQueryChosenChat)
	}
	inb.SwitchInlineQueryChosenChat = sw
	logs.DataWrittenSuccessfully(inbtn, "Switch Inline Query Chosen Chat")
}

func (inb *InlineKeyboardButton) WriteCallbackGame(game *types.CallbackGame) {
	if inb.CallbackGame != nil {
		logs.DataIsntEmply(inbtn, "Callback Game", inb.CallbackGame)
	}
	inb.CallbackGame = game
	logs.DataWrittenSuccessfully(inbtn, "Callback Game")
}

func (inb *InlineKeyboardButton) WritePay() {
	if inb.Pay {
		logs.DataIsntEmply(inbtn, "Pay", inb.Pay)
	}
	inb.Pay = true
	logs.SettedParam("Pay", inbtn, true)
}

func (rp *Reply) Set(plan []int) {
	rp.Keyboard = make([][]*ReplyKeyboardButton, len(plan))
	for i := range rp.Keyboard {
		rp.Keyboard[i] = make([]*ReplyKeyboardButton, plan[i])
	}
}

func (rp *Reply) NewButton(line, pos int) (*ReplyKeyboardButton, error) {
	var err error

	if len(rp.Keyboard) > line && len(rp.Keyboard[line]) > pos {

		if rp.Keyboard[line][pos] != nil {
			logs.DataIsntEmply(replyKB, fmt.Sprintf("%s line: %d, position: %d", button, line, pos), rp.Keyboard[line][pos])
		}
		rp.Keyboard[line][pos] = new(ReplyKeyboardButton)

		return rp.Keyboard[line][pos], nil

	} else {
		if len(rp.Keyboard) > line {
			err = errors.ButtosDoesntFit("line", line)
		} else if len(rp.Keyboard[line]) > pos {
			err = errors.ButtosDoesntFit("pos", pos)
		}
	}
	return nil, err
}

func (rpb *ReplyKeyboardButton) WriteString(text string) {
	if rpb.Text != "" {
		logs.DataIsntEmply(rpbtn, "Text", rpb.Text)
	}
	rpb.Text = text
	logs.DataWrittenSuccessfully(rpbtn, "Text")
}

func (rpb *ReplyKeyboardButton) WriteRequestUsers(requs *types.KeyboardButtonRequestUsers) {
	if rpb.RequestUsers != nil {
		logs.DataIsntEmply(rpbtn, "Request Users", rpb.RequestUsers)
	}
	rpb.RequestUsers = requs
	logs.DataWrittenSuccessfully(rpbtn, "Request Users")
}

func (rpb *ReplyKeyboardButton) WriteRequestChat(reqch *types.KeyboardButtonRequestChat) {
	if rpb.RequestChat != nil {
		logs.DataIsntEmply(rpbtn, "Request Chat", rpb.RequestChat)
	}
	rpb.RequestChat = reqch
	logs.DataWrittenSuccessfully(rpbtn, "Request Chat")
}

func (rpb *ReplyKeyboardButton) WriteRequestContact() {
	if rpb.RequestContact {
		logs.DataIsntEmply(rpbtn, "Request Contact", rpb.RequestContact)
	}
	rpb.RequestContact = true
	logs.SettedParam("Request Contact", rpbtn, true)
}

func (rpb *ReplyKeyboardButton) WriteRequestLocation() {
	if rpb.RequestLocation {
		logs.DataIsntEmply(rpbtn, "Request Location", rpb.RequestLocation)
	}
	rpb.RequestLocation = true
	logs.SettedParam("Request Location", rpbtn, true)
}

func (rpb *ReplyKeyboardButton) WriteRequestPoll(poll *types.KeyboardButtonPollType) {
	if rpb.RequestPoll != nil {
		logs.DataIsntEmply(rpbtn, "Request Poll", rpb.RequestPoll)
	}
	rpb.RequestPoll = poll
	logs.DataWrittenSuccessfully(rpbtn, "Request Poll")
}

func (rpb *ReplyKeyboardButton) WriteWebApp(webapp *types.WebAppInfo) {
	if rpb.WebApp != nil {
		logs.DataIsntEmply(rpbtn, "Web App", rpb.WebApp)
	}
	rpb.WebApp = webapp
	logs.DataWrittenSuccessfully(rpbtn, "Web App")
}
