package chat

import (
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

func (ch *Chat) MultipartFields(writer *multipart.Writer) error {
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
