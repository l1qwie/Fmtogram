package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/testbotdata"
	"github.com/l1qwie/Fmtogram/types"
)

func CreateMessage(tg *types.Telegram) *Message {
	m := new(Message)
	m.fm = new(formatter)
	m.fm.ch = &chat{ID: tg.Result[0].Message.Chat.ID, BusinessConnectionID: tg.Result[0].BusinessMessage.BusinessConnectionID}
	m.fm.inf = &information{}
	m.fm.mh = &mediaHolder{}
	return m
}

func CreateEmpltyMessage() *Message {
	m := new(Message)
	m.fm = new(formatter)
	m.fm.ch = &chat{}
	m.fm.inf = &information{}
	m.fm.mh = &mediaHolder{}
	return m
}

func requiredMedia(msg *Message, tgr *interface{}, object *string) (string, error) {
	var err error
	var APIMethod string

	if msg.fm.ch.ID == 0 {
		err = errors.MissedRequiredField("Chat", "ChatID", 0, false)
	}
	for i, j := 0, 0; (i < len(msg.fm.mh.storage)) && (err == nil); i++ {
		if msg.fm.mh.storage[i] != nil {
			switch m := msg.fm.mh.storage[i].(type) {
			case *photo:
				requiredPhotoData(m, i)
				if APIMethod == "" {
					APIMethod = methods.Photo
					*object = "Photo"
				}
			case *audio:
				requiredAudioData(m, i)
				if APIMethod == "" {
					APIMethod = methods.Audio
					*object = "Audio"
				}
			case *video:
				requiredVideoData(m, i)
				if APIMethod == "" {
					APIMethod = methods.Video
					*object = "Video"
				}
			case *document:
				requiredDocumentData(m, i)
				if APIMethod == "" {
					APIMethod = methods.Document
					*object = "Document"
				}
			}
			if j >= 1 {
				APIMethod = methods.MediaGroup
				*object = "for any media you have mentioned {Photo/Video/Document/Audio}"
				*tgr = new(types.Telegram)
			} else {
				*tgr = new(types.TelegramResponse)
			}
			j++
		}
	}
	return APIMethod, err
}

func requiredMessage(msg *Message, tgr *interface{}, APIMethod, object string) (string, error) {
	var err error

	if msg.fm.ch.ID == 0 {
		err = errors.MissedRequiredField("Chat", "ChatID", 0, false)
	}
	if APIMethod == "" && msg.fm.inf.Text == "" {
		err = errors.MissedRequiredField("Message", "Text", 0, false)
	}
	if APIMethod != "" && msg.fm.inf.Text != "" {
		logs.DataMightBeLost("Message", "Text", msg.fm.inf.Text, object, "Caption")
	}
	if APIMethod == "" {
		APIMethod = methods.Message
		msg.fm.contentType = "application/json"
		*tgr = new(types.TelegramResponse)
	}
	return APIMethod, err
}

func mediaPart(msg *Message) error {
	var err error

	if msg.fm.mh.amount == 1 {
		err = uniqueMedia(msg)
	} else {
		err = mediaGroup(msg)
	}
	return err
}

func messagePart(msg *Message) error {
	var err error
	var bytes []byte

	if msg.fm.mh.atLeastOnce {
		err = msg.fm.inf.multipartFields(msg.fm.writer)
	} else {

		bytes, err = msg.fm.inf.createJSON()

		if err == nil {
			msg.fm.buf.Write(bytes)
		}
	}
	return err
}

func keyboardPart(msg *Message) error {
	var err error
	var bytes []byte

	if msg.fm.kb != nil {

		if msg.fm.mh.atLeastOnce {
			err = msg.fm.kb.MultipartFields(msg.fm.writer)
		} else {

			bytes, err = msg.fm.kb.JsonFields()

			if err == nil {
				msg.fm.buf.Write(bytes)
			}
		}

	}
	return err
}

func chatPart(msg *Message) error {
	var err error
	var bytes []byte

	if msg.fm.mh.atLeastOnce {
		err = msg.fm.ch.multipartFields(msg.fm.writer)
	} else {

		bytes, err = msg.fm.ch.createJson()

		if err == nil {
			msg.fm.buf.Write(bytes)
		}
	}

	return err
}

// not a final result
func uniteEverything(msg *Message) error {
	decoder := json.NewDecoder(msg.fm.buf)

	result := make(map[string]interface{})

	for {
		var data map[string]interface{}
		if err := decoder.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("error decoding JSON: %s", err)
		}

		for k, v := range data {
			result[k] = v
		}
	}

	mergedJSON, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("error serializing merged JSON: %s", err)
	}

	msg.fm.buf = bytes.NewBuffer(nil)
	msg.fm.buf.Write(mergedJSON)
	return nil
}

func makeRequest(msg *Message, tgr *interface{}) (string, error) {
	var err error
	var APIMethod, object string
	var shouldSkip [4]bool
	// chat part must be skipped by default
	// it isn't skipped now, because of tests
	doPlan := [4]func(*Message) error{mediaPart, messagePart, keyboardPart, chatPart}

	msg.fm.buf = bytes.NewBuffer(nil)
	msg.fm.writer = multipart.NewWriter(msg.fm.buf)

	if msg.fm.mh.amount > 0 {
		APIMethod, err = requiredMedia(msg, tgr, &object)
	} else {
		shouldSkip[0] = true
	}

	if err == nil {
		APIMethod, err = requiredMessage(msg, tgr, APIMethod, object)
	}

	for i := 0; (i < len(doPlan)) && (err == nil); i++ {
		if !shouldSkip[i] {
			err = doPlan[i](msg)
		}
	}
	if err == nil && !msg.fm.mh.atLeastOnce {
		err = uniteEverything(msg)
	} else {
		err = msg.fm.writer.Close()
	}
	return APIMethod, err
}

func distributorTelegram(msg *Message, t *types.Telegram) {
	for i := 0; i < len(msg.fm.mh.storage); i++ {

		if msg.fm.mh.storage[i] != nil {
			switch m := msg.fm.mh.storage[i].(type) {
			case *photo:

				for j := 0; j < len(t.Result[0].Message.Photo); j++ {
					m.ResponseData[j] = &responseData{
						FileID:       t.Result[0].Message.Photo[j].FileID,
						FileSize:     t.Result[0].Message.Photo[j].FileSize,
						FileUniqueID: t.Result[0].Message.Photo[j].FileUniqueID,
						Width:        t.Result[0].Message.Photo[j].Width,
						Height:       t.Result[0].Message.Photo[j].Height,
					}
				}
				// case *media.Video:
				// case *media.Audio:
				// case *media.Document:
			}
		}
	}
}

func distributorTelegramResponse(msg *Message, t *types.TelegramResponse) {
	var j int
	for i := 0; i < len(msg.fm.mh.storage); i++ {

		if msg.fm.mh.storage[i] != nil {
			switch m := msg.fm.mh.storage[i].(type) {
			case *photo:
				for k := 0; k < len(t.Result.Photo); k++ {
					m.ResponseData[j] = &responseData{
						FileID:       t.Result.Photo[k].FileID,
						FileSize:     t.Result.Photo[k].FileSize,
						FileUniqueID: t.Result.Photo[k].FileUniqueID,
						Width:        t.Result.Photo[k].Width,
						Height:       t.Result.Photo[k].Height,
					}
				}
			case *video:
				m.ResponseData = t.Result.Video
			case *audio:
				m.ResponseData = t.Result.Audio
			case *document:
				m.ResponseData = t.Result.Document
			}
			j++
		}
	}
}

func sendRequest(msg *Message, tgr interface{}, APIMethod string) error {
	var resp *http.Response
	var body []byte

	log.Print(msg.fm.buf.String())
	log.Print(APIMethod)
	log.Print(msg.fm.contentType)

	url := fmt.Sprint(types.TelegramAPI, testbotdata.Token, "/", APIMethod)

	log.Print(url)

	req, err := http.NewRequest("POST", url, msg.fm.buf)
	req.Header.Set("Content-Type", msg.fm.contentType)

	if err != nil {
		err = errors.CantMakeRequest(err)
	} else {

		client := http.Client{}
		resp, err = client.Do(req)

		if err != nil {
			err = errors.CantGetResponse(err)
		} else {

			defer resp.Body.Close()
			body, err = io.ReadAll(resp.Body)
			if err != nil {
				err = errors.CantReadResponse(err)
			} else {

				fmt.Println(string(body))

				err = json.Unmarshal(body, tgr)
				if err != nil {
					err = errors.CantUnmarshal(err)
				} else {
					switch t := tgr.(type) {
					case *types.Telegram:
						if t.ErrorCode != 0 {
							err = errors.TelegramError(t.ErrorCode, t.Description)
						} else {
							distributorTelegram(msg, t)
						}
					case *types.TelegramResponse:
						if t.ErrorCode != 0 {
							err = errors.TelegramError(t.ErrorCode, t.Description)
						} else {
							distributorTelegramResponse(msg, t)
						}
					}
				}
			}
		}
	}
	return err
}

// Sends the message you created. May send you an error if you made a mistake during the build process.
// Also sends a new pointer to the structure. This is a quick response from Telegram to your message. It collects some data
// about the message sent. For example: if you want to send a photo, this structure will have data about it from Telegram (id, size, etc.)
func (msg *Message) Send() (interface{}, error) {
	var err error
	var APIMethod string
	var tgr interface{}

	if APIMethod, err = makeRequest(msg, &tgr); err == nil {
		if tgr == nil {
			panic("!~")
		}
		err = sendRequest(msg, tgr, APIMethod)
	}

	return tgr, err
}
