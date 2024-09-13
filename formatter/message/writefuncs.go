package message

import (
	"fmt"

	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/types"
)

const (
	object string = "Message"
)

func (m *Message) WriteString(text string) {
	if m.Text != "" {
		logs.DataIsntEmply(object, "Text", m.Text)
	}
	m.Text = text
	logs.DataWrittenSuccessfully(object, "Text")
}

func (m *Message) WriteParseMode(parsemode string) {
	if m.ParseMode != "" {
		logs.DataIsntEmply(object, "Parse Mode", m.ParseMode)
	}
	m.ParseMode = parsemode
	logs.DataWrittenSuccessfully(object, "Parse Mode")
}

func (m *Message) WriteMessageThreadID(messageID int) {
	if m.MessageThreadID != 0 {
		logs.DataIsntEmply(object, "Message Thread ID", m.MessageThreadID)
	}
	m.MessageThreadID = messageID
	logs.DataWrittenSuccessfully(object, "Message Thread ID")
}

func (m *Message) WriteDisableNotification() {
	if m.DisableNotification {
		logs.DataIsntEmply(object, "Disable Notification", m.DisableNotification)
	}
	m.DisableNotification = true
	logs.SettedParam("Disable Notification", object, m.DisableNotification)
}

func (m *Message) WriteProtectContent() {
	if m.ProtectContent {
		logs.DataIsntEmply(object, "Protect Content", m.ProtectContent)
	}
	m.ProtectContent = true
	logs.SettedParam("Protect Content", object, m.ProtectContent)
}

func (m *Message) WriteMessageEffectID(messageID string) {
	if m.MessageEffectID != "" {
		logs.DataIsntEmply(object, "Message Effect ID", m.MessageEffectID)
	}
	m.MessageEffectID = messageID
	logs.DataWrittenSuccessfully(object, "Message Effect ID")
}

func (m *Message) WriteReplyParametrs(params *types.ReplyParameters) {
	if m.ReplyParameters != nil {
		logs.DataIsntEmply(object, "Reply Parametrs", m.ReplyParameters)
	}
	m.ReplyParameters = params
	logs.DataWrittenSuccessfully(object, "Reply Parametrs")
}

func (m *Message) WriteEntities(entities []*types.MessageEntity) {
	if len(m.Entities) != 0 {
		logs.DataIsntEmply(object, "Entities", m.Entities)
	}
	m.Entities = entities
	logs.DataWrittenSuccessfully(object, "Entities")
}

func (m *Message) WriteLinkPreviewOptions(lpo *types.LinkPreviewOptions) {
	if m.LinkPreviewOptions != nil {
		logs.DataIsntEmply(object, "Link Preview Options", m.LinkPreviewOptions)
	}
	m.LinkPreviewOptions = lpo
	logs.DataWrittenSuccessfully(object, "Link Preview Options")
}

func checkReplyTypes(markup interface{}) bool {
	var ok bool
	switch markup.(type) {
	case *types.InlineKeyboardMarkup:
		ok = true
	case *types.ReplyKeyboardMarkup:
		ok = true
	case *types.ReplyKeyboardRemove:
		ok = true
	case *types.ForceReply:
		ok = true
	}

	return ok
}

func (m *Message) WriteReplyMarkup(markup interface{}) error {
	var err error
	ok := checkReplyTypes(markup)
	if ok {
		if m.ReplyMarkup != nil {
			logs.DataIsntEmply(object, "Reply Markup", m.ReplyMarkup)
		}
		m.ReplyMarkup = markup
		logs.DataWrittenSuccessfully(object, "Reply Markup")
	} else {
		err = fmt.Errorf("WriteReplyMarkup is waiting an interface{}, but the exactly type of the interface{} must be *types.InlineKeyboardMarkup, *types.ReplyKeyboardMarkup, *types.ReplyKeyboardRemove or *types.ForceReply")
	}
	return err
}
