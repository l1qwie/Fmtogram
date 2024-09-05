package message

import (
	"encoding/json"
	"fmt"
	"mime/multipart"

	"github.com/l1qwie/Fmtogram/errors"
)

func (m *Message) MultipartFiellds(writer *multipart.Writer) error {
	var err error
	if m.BusinessConnectionID != "" {
		err = writer.WriteField("business_connection_id", m.BusinessConnectionID)

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && m.ChatID != nil {
		err = writer.WriteField("chat_id", fmt.Sprint(m.ChatID))

		if err != nil {
			errors.CantWriteField(err)
		}
	} else if m.ChatID == nil {
		err = errors.ChatIDIsMissed()
	}
	if err == nil && m.MessageThreadID != 0 {
		err = writer.WriteField("message_thread_id", string(m.MessageThreadID))

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && m.DisableNotification {
		err = writer.WriteField("disable_notification", "True")

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && m.ProtectContent {
		err = writer.WriteField("protect_content", "True")

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && m.MessageEffectID != "" {
		err = writer.WriteField("message_effect_id", m.MessageEffectID)

		if err != nil {
			errors.CantWriteField(err)
		}
	}
	if err == nil && m.ReplyParameters != nil {
		body, err1 := json.Marshal(m.ReplyParameters)
		if err1 == nil {
			err = writer.WriteField("reply_parameters", string(body))

			if err != nil {
				errors.CantWriteField(err)
			}
		} else {
			errors.CantMarshalJSON(err)
		}
	}
	return err
}
