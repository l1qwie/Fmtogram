package fmtogram

import (
	"log"
	"time"

	"github.com/l1qwie/Fmtogram/executer"
	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/helper"
	"github.com/l1qwie/Fmtogram/types"
)

var StartFunc func(*types.Telegram, *types.Returned) *formatter.Formatter

func pollResponse(output chan *formatter.Formatter, reg *executer.RegTable) {
	var (
		offset int
		err    error
		tg     *types.Telegram
		index  int
		chatID int
	)

	err = executer.RequestOffset(types.BotID, &offset)
	for err != nil {
		err = executer.RequestOffset(types.BotID, &offset)
		time.Sleep(time.Second / 10)
	}
	log.Print("The bot has found the first request from Telegram")
	for {
		tg = new(types.Telegram)
		err = executer.Updates(&offset, tg)
		if len(tg.Result) != 0 && err == nil {
			log.Print("The bots has found some information about user without any misstakes")
			chatID = helper.ReturnChatId(tg)
			index = reg.Seeker(chatID)
			if index != executer.None {
				reg.Reg[index].Chu <- tg
			} else {
				index = reg.NewIndex()
				reg.Reg[index].UserId = chatID
				reg.Reg[index].Chu = make(chan *types.Telegram, 1)
				reg.Reg[index].Chu <- tg
			}
			go worker(reg.Reg[index].Chu, reg.Reg[index].Chb, output)
			offset = offset + 1
		} else if err != nil {
			log.Print("ERR FROM Updates():", err)
		}
		time.Sleep(time.Second / 10)
	}
}

func worker(input chan *types.Telegram, mesoutput chan *types.Returned, output chan *formatter.Formatter) {
	var (
		fm         *formatter.Formatter
		userstruct *types.Telegram
		mes        *types.Returned
	)
	for len(input) > 0 {
		userstruct = <-input
		if len(mesoutput) > 0 {
			mes = <-mesoutput
		} else {
			mes = &types.Returned{
				Ok: false,
				Result: types.Message{
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
		fm = StartFunc(userstruct, mes)
		if err := fm.Complete(); err == nil {
			output <- fm
		}
	}
}

func pushRequest(requests <-chan *formatter.Formatter, reg *executer.RegTable) {
	for r := range requests {
		mes, err := r.Make()
		if err != nil {
			log.Print(err)
		}
		if mes.Ok {
			index := reg.Seeker(mes.Result.Chat.Id)
			reg.Reg[index].Chb = make(chan *types.Returned, 1)
			reg.Reg[index].Chb <- mes
		}
	}
}

func Start() {
	log.Print("The bot has been turned on")
	requests := make(chan *formatter.Formatter)
	reg := new(executer.RegTable)
	go pollResponse(requests, reg)
	go pushRequest(requests, reg)
	for {
		time.Sleep(time.Second)
	}
}
