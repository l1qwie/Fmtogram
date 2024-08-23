package fmtogram

import (
	"time"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/executer"
	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/helper"
	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/types"
)

var StartFunc func(*types.Telegram, *types.Returned) *formatter.Formatter

func firstStep(offset *int) {
	logs.GetOffset()

	err := executer.RequestOffset(types.BotID, offset)
	for err != nil {
		err = executer.RequestOffset(types.BotID, offset)
		time.Sleep(time.Second / 10)
	}

	logs.GottenOffset()
}

func queue(reg *executer.RegTable, tg *types.Telegram, chatID, index int) {
	if index != executer.None {
		logs.AlreadyInQueue()

		reg.Reg[index].Chu <- tg
	} else {
		logs.FirstTimeUser()

		index = reg.NewIndex()
		reg.Reg[index].UserId = chatID
		reg.Reg[index].Chu = make(chan *types.Telegram, 1)
		reg.Reg[index].Chu <- tg
	}
}

func pullResponse(output chan *formatter.Formatter, reg *executer.RegTable) {
	var offset int
	firstStep(&offset)
	for {
		tg := new(types.Telegram)
		err := executer.Updates(&offset, tg)
		if len(tg.Result) != 0 && err == nil {
			logs.FoundSomeIformation()

			chatID := helper.ReturnChatId(tg)
			index := reg.Seeker(chatID)

			queue(reg, tg, chatID, index)
			go worker(reg.Reg[index].Chu, reg.Reg[index].Chb, output)

			offset = offset + 1
		} else if err != nil {
			errors.ErrorInUserData(err)
		}
		time.Sleep(time.Second / 10)
	}
}

func worker(input chan *types.Telegram, returned chan *types.Returned, output chan *formatter.Formatter) {
	userstruct := new(types.Telegram)
	mes := new(types.Returned)

	for len(input) > 0 {
		userstruct = <-input
		if len(returned) > 0 {
			logs.ReturnedIsntEmply()
			mes = <-returned
		} else {
			logs.ReturnedIsEmply()
			mes = &types.Returned{
				Ok: false,
				Result: types.ReturnedMessage{
					MessageId: 0,
					Chat: types.Chat{
						Id: 0,
					},
					Photo: []types.Photo{{
						FileId: "",
					}},
				},
			}
		}
		logs.CallDeveloperFunc()
		fm := StartFunc(userstruct, mes)
		if err := fm.Complete(); err == nil {
			output <- fm
		}
	}
}

func pushRequest(requests <-chan *formatter.Formatter, reg *executer.RegTable) {
	for r := range requests {
		logs.GottenRequest()

		mes, err := r.Make()
		if err != nil {
			logs.ErrorDuringSending()
		}
		if mes.Ok {
			logs.ReturnedIsntEmply()
			index := reg.Seeker(mes.Result.Chat.Id)
			reg.Reg[index].Chb = make(chan *types.Returned, 1)
			reg.Reg[index].Chb <- mes
		}
	}
}

func Start() {
	defer logs.TurnOff()
	logs.DebugOrInfo()
	logs.TurnOn()

	requests := make(chan *formatter.Formatter)
	reg := new(executer.RegTable)
	go pullResponse(requests, reg)
	go pushRequest(requests, reg)

	for {
		time.Sleep(time.Second * 10)
	}
}
