package formatter

import (
	"fmt"

	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/types"
)

func (fm *Formatter) WriteChatID(chatID int) {
	fm.Chat.ChatID = chatID
	logs.DataWrittenSuccessfully("Chat ID")
}

func (fm *Formatter) WriteChatName(chatname string) {
	fm.Chat.ChatID = fmt.Sprint("@", chatname)
	logs.DataWrittenSuccessfully("Chat Name")
}

func (fm *Formatter) WriteChatURL(chatURL string) {
	fm.Chat.ChatID = chatURL
	logs.DataWrittenSuccessfully("Chat URL")
}

func (fm *Formatter) WriteBusinessConnectionID(connectionID string) {
	fm.Business.BusinessConnectionID = connectionID
	logs.DataWrittenSuccessfully("Business Connection ID")
}

func (fm *Formatter) WriteThreadID(threadID string) {
	fm.Message.ThreadID = threadID
	logs.DataWrittenSuccessfully("Message Thread ID")
}

func (fm *Formatter) WriteReplyMessageID(messageID int) {
	fm.Chat.ReplyParameters.MessageID = messageID
	logs.DataWrittenSuccessfully("Replyed Message ID")
}

func (fm *Formatter) WriteReplyChatID(chatID int) {
	fm.Chat.ReplyParameters.ChatID = chatID
	logs.DataWrittenSuccessfully("Replyed Chat ID")
}

func (fm *Formatter) WriteReplyChatName(chatname string) {
	fm.Chat.ReplyParameters.ChatID = chatname
	logs.DataWrittenSuccessfully("Replyed Chat Name")
}

func (fm *Formatter) WriteReplyChatURL(chatURL string) {
	fm.Chat.ReplyParameters.ChatID = chatURL
	logs.DataWrittenSuccessfully("Replyed Chat URL")
}

func (fm *Formatter) MakeAllowedSendingWithoutReply() {
	fm.Chat.ReplyParameters.AllowSendingWithoutReply = true
	logs.DataWrittenSuccessfully("Allow Sending Without Reply")
}

func (fm *Formatter) WriteReplyQuote(quote string) {
	fm.Chat.ReplyParameters.Quote = quote
	logs.DataWrittenSuccessfully("Replyed Quote")
}

func (fm *Formatter) WriteReplyQuoteParseMode(parsemode string) {
	fm.Chat.ReplyParameters.QuoteParseMode = parsemode
	logs.DataWrittenSuccessfully("Replyed Quote Parse Mode")
}

func (fm *Formatter) AddReplyQuoteEntities(entities []*types.MessageEntity) {
	fm.Chat.ReplyParameters.QuoteEntities = entities
	logs.DataWrittenSuccessfully("Replyed Quote Entities")
}

func (fm *Formatter) AddReplyQuotePosition(position int) {
	fm.Chat.ReplyParameters.QuotePosition = position
	logs.DataWrittenSuccessfully("Replyed Quote Position")
}
