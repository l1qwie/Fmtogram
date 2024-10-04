package chat

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

func MultipartFields(ch *Chat, writer *multipart.Writer) error {
	var err error
	if ch.ID != nil {
		err = field(writer, "chat_id", fmt.Sprint(ch.ID))

	} else if ch.ID == nil {
		err = errors.ChatIDIsMissed()
	}
	if ch.BusinessConnectionID != "" {
		err = field(writer, "business_connection_id", ch.BusinessConnectionID)
	}
	return err
}

func CreateJson(ch *Chat) ([]byte, error) {
	return json.Marshal(ch)
}
