package message

import (
	"encoding/json"
	"fmt"
	"mime/multipart"

	"github.com/l1qwie/Fmtogram/errors"
)

func field(w *multipart.Writer, fieldname, value string) error {
	err := w.WriteField(fieldname, value)
	if err != nil {
		errors.CantWriteField(err)
	}
	return err
}

func MultipartFields(m *Message, writer *multipart.Writer) error {
	var err error
	if m.Text != "" {
		err = field(writer, "text", m.Text)
	}
	if err == nil && m.Caption != "" {
		err = field(writer, "caption", m.Caption)
	}
	if err == nil && m.ParseMode != "" {
		err = field(writer, "parse_mode", m.ParseMode)
	}
	if err == nil && m.MessageThreadID != 0 {
		err = field(writer, "message_thread_id", fmt.Sprint(m.MessageThreadID))
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

func CreateJSON(message *Message) ([]byte, error) {
	body, err := json.Marshal(message)
	if err != nil {
		errors.CantMarshalJSON(err)
	}
	return body, err
}
