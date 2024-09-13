package keyboard

import (
	"encoding/json"
	"mime/multipart"
)

func (in *Inline) MultipartFields(writer *multipart.Writer) error {
	inbody, err := json.Marshal(in.InlineKeyboard)
	if err == nil {
		err = writer.WriteField("reply_markup", string(inbody))
	}
	return err
}

func (in *Inline) JsonFields() ([]byte, error) {
	return json.Marshal(in.InlineKeyboard)
}

func (rp *Reply) MultipartFields(writer *multipart.Writer) error {
	rpbody, err := json.Marshal(rp.Keyboard)
	if err == nil {
		err = writer.WriteField("reply_markup", string(rpbody))
	}
	return err
}

func (rp *Reply) JsonFields() ([]byte, error) {
	return json.Marshal(rp.Keyboard)
}
