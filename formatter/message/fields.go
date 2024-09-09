package message

import (
	"encoding/json"
	"fmt"
	"mime/multipart"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/formatter/methods"
)

func field(w *multipart.Writer, fieldname, value string) error {
	err := w.WriteField(fieldname, value)
	if err != nil {
		errors.CantWriteField(err)
	}
	return err
}

func (m *Message) captionText(method string) {
	var found bool
	for i := 0; i < len(methods.Media) && !found; i++ {
		if methods.Media[i] == method {
			found = true
		}
	}
	if found {
		if method != methods.MediaGroup {
			m.Caption = m.Text
			m.Text = ""
		}
	}
}

func (m *Message) MultipartFields(writer *multipart.Writer, method string) error {
	var err error
	m.captionText(method)
	if m.BusinessConnectionID != "" {
		err = field(writer, "business_connection_id", m.BusinessConnectionID)
	}
	if err == nil && m.Text != "" {
		err = field(writer, "text", m.Text)
	}
	if err == nil && m.Caption != "" {
		err = field(writer, "caption", m.Caption)
	}
	if err == nil && m.ParseMode != "" {
		err = field(writer, "parse_mode", m.ParseMode)
	}
	if err == nil && m.ChatID != nil {
		err = field(writer, "chat_id", fmt.Sprint(m.ChatID))

	} else if m.ChatID == nil {
		err = errors.ChatIDIsMissed()
	}
	if err == nil && m.MessageThreadID != 0 {
		err = field(writer, "message_thread_id", string(m.MessageThreadID))
	}
	if err == nil && m.DisableNotification {
		err = field(writer, "disable_notification", "True")
	}
	if err == nil && m.ProtectContent {
		err = field(writer, "protect_content", "True")
	}
	if err == nil && m.MessageEffectID != "" {
		err = field(writer, "message_effect_id", m.MessageEffectID)
	}
	if err == nil && m.ReplyParameters != nil {
		body, err1 := json.Marshal(m.ReplyParameters)
		if err1 == nil {
			err = field(writer, "reply_parameters", string(body))
		} else {
			errors.CantMarshalJSON(err)
		}
	}
	return err
}

func CreateJSON(message *Message, method string) ([]byte, error) {
	message.captionText(method)
	body, err := json.Marshal(message)
	if err != nil {
		errors.CantMarshalJSON(err)
	}
	return body, err
}
