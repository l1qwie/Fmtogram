package message

import (
	"fmt"

	"github.com/l1qwie/Fmtogram/logs"
)

func (m *Message) WriteChatID(chatID int) {
	m.ChatID = chatID
	logs.DataWrittenSuccessfully("Chat ID")
}

func (m *Message) WriteChatName(chatname string) {
	m.ChatID = fmt.Sprint("@", chatname)
	logs.DataWrittenSuccessfully("Chat Name")
}

func (m *Message) WriteChatURL(chatURL string) {
	m.ChatID = chatURL
	logs.DataWrittenSuccessfully("Chat URL")
}

func (m *Message) WriteBusinessConnectionID(connectionID string) {
	m.BusinessConnectionID = connectionID
	logs.DataWrittenSuccessfully("Business Connection ID")
}

// func (m *Message) WriteThreadID(threadID string) {
// 	m.ThreadID = threadID
// 	logs.DataWrittenSuccessfully("Message Thread ID")
// }

// func (m *Message) WriteReplyMessageID(messageID int) {
// 	fm.Chat.ReplyParameters.MessageID = messageID
// 	logs.DataWrittenSuccessfully("Replyed Message ID")
// }

// func (m *Message) WriteReplyChatID(chatID int) {
// 	fm.Chat.ReplyParameters.ChatID = chatID
// 	logs.DataWrittenSuccessfully("Replyed Chat ID")
// }

// func (m *Message) WriteReplyChatName(chatname string) {
// 	fm.Chat.ReplyParameters.ChatID = chatname
// 	logs.DataWrittenSuccessfully("Replyed Chat Name")
// }

// func (m *Message) WriteReplyChatURL(chatURL string) {
// 	fm.Chat.ReplyParameters.ChatID = chatURL
// 	logs.DataWrittenSuccessfully("Replyed Chat URL")
// }

// func (m *Message) MakeAllowedSendingWithoutReply() {
// 	fm.Chat.ReplyParameters.AllowSendingWithoutReply = true
// 	logs.DataWrittenSuccessfully("Allow Sending Without Reply")
// }

// func (m *Message) WriteReplyQuote(quote string) {
// 	fm.Chat.ReplyParameters.Quote = quote
// 	logs.DataWrittenSuccessfully("Replyed Quote")
// }

// func (m *Message) WriteReplyQuoteParseMode(parsemode string) {
// 	fm.Chat.ReplyParameters.QuoteParseMode = parsemode
// 	logs.DataWrittenSuccessfully("Replyed Quote Parse Mode")
// }

// func (m *Message) AddReplyQuoteEntities(entities []*types.MessageEntity) {
// 	fm.Chat.ReplyParameters.QuoteEntities = entities
// 	logs.DataWrittenSuccessfully("Replyed Quote Entities")
// }

// func (m *Message) AddReplyQuotePosition(position int) {
// 	fm.Chat.ReplyParameters.QuotePosition = position
// 	logs.DataWrittenSuccessfully("Replyed Quote Position")
// }
