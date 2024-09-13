package chat

import (
	"fmt"

	"github.com/l1qwie/Fmtogram/logs"
)

const (
	object string = "Chat"
)

func (ch *Chat) WriteChatID(chatID int) {
	if ch.ID != nil {
		logs.DataIsntEmply(object, "Chat ID", ch.ID)
	}
	ch.ID = chatID
	logs.DataWrittenSuccessfully(object, "Chat ID")
}

func (ch *Chat) WriteChatName(chatname string) {
	if ch.ID != nil {
		logs.DataIsntEmply(object, "Chat Name", ch.ID)
	}
	ch.ID = fmt.Sprint("@", chatname)
	logs.DataWrittenSuccessfully(object, "Chat Name")
}

func (ch *Chat) WriteChatURL(chatURL string) {
	if ch.ID != nil {
		logs.DataIsntEmply(object, "Chat URL", ch.ID)
	}
	ch.ID = chatURL
	logs.DataWrittenSuccessfully(object, "Chat URL")
}

func (ch *Chat) WriteBusinessConnectionID(connectionID string) {
	if ch.BusinessConnectionID != "" {
		logs.DataIsntEmply(object, "Business Connection ID", ch.BusinessConnectionID)
	}
	ch.BusinessConnectionID = connectionID
	logs.DataWrittenSuccessfully(object, "Business Connection ID")
}
