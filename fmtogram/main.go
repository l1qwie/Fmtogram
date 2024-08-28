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

var StartFunc func(*types.Telegram, *types.Telegram) *formatter.Formatter

func firstStep(offset *int) {
	logs.GetOffset()

	err := executer.GetOffset(offset)
	for err != nil {
		err = executer.GetOffset(offset)
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
		err := executer.GetUpdates(tg, &offset)
		if len(tg.Result) != 0 && err == nil {
			logs.FoundSomeIformation()

			chatID := helper.GetUserID(tg)
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

func worker(input chan *types.Telegram, returned chan *types.Telegram, output chan *formatter.Formatter) {
	// // tg := new(types.Telegram)
	// // tgReturned := new(types.Telegram)

	// for len(input) > 0 {
	// 	tg = <-input
	// 	if len(returned) > 0 {
	// 		logs.ReturnedIsntEmply()
	// 		tgReturned = <-returned
	// 	} else {
	// 		logs.ReturnedIsEmply()
	// 	}
	// 	logs.CallDeveloperFunc()
	// 	// fm := StartFunc(tg, tgReturned)
	// 	// if err := fm.Complete(); err == nil {
	// 	// 	output <- fm
	// 	// }
	// }
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
			index := reg.Seeker(mes.Result[0].Message.From.ID)
			reg.Reg[index].Chb = make(chan *types.Telegram, 1)
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
