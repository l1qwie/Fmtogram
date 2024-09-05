package formatter

import (
	"github.com/l1qwie/Fmtogram/formatter/message"
)

func (fm *Formatter) NewChat() *message.Message {
	message := new(message.Message)
	fm.m = message
	return message
}
