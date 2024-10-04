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
	"github.com/l1qwie/Fmtogram/formatter/chat"
	"github.com/l1qwie/Fmtogram/formatter/media"
	"github.com/l1qwie/Fmtogram/formatter/message"
	"github.com/l1qwie/Fmtogram/formatter/methods"
	"github.com/l1qwie/Fmtogram/testbotdata"
	"github.com/l1qwie/Fmtogram/types"
)

func CreateMessage(tg *types.Telegram) *Message {
	m := new(Message)
	m.fm = new(formatter)
	m.fm.ch = &chat.Chat{ID: tg.Result[0].Message.Chat.ID, BusinessConnectionID: tg.Result[0].BusinessMessage.BusinessConnectionID}
	m.fm.m = &message.Message{}
	m.fm.mh = &mediaHolder{}
	return m
}

func CreateEmpltyMessage() *Message {
	m := new(Message)
	m.fm = new(formatter)
	m.fm.ch = &chat.Chat{}
	m.fm.m = &message.Message{}
	m.fm.mh = &mediaHolder{}
	return m
}

func requiredMedia(msg *Message) (string, error) {
	var err error
	var APIMethod string
	if msg.fm.ch.ID == 0 {
		err = errors.MissedRequiredField("Chat", "ChatID", 0, false)
	}
	for i := 0; (i < len(msg.fm.mh.storage)) && (err == nil); i++ {
		if msg.fm.mh.storage[i] != nil {
			switch m := msg.fm.mh.storage[i].(type) {
			case *media.Photo:
				media.RequiredPhotoData(m, i)
				if APIMethod == "" {
					APIMethod = methods.Photo
				}
			case *media.Audio:
				media.RequiredAudioData(m, i)
				if APIMethod == "" {
					APIMethod = methods.Audio
				}
			case *media.Video:
				media.RequiredVideoData(m, i)
				if APIMethod == "" {
					APIMethod = methods.Video
				}
			case *media.Document:
				media.RequiredDocumentData(m, i)
				if APIMethod == "" {
					APIMethod = methods.Document
				}
			}
			if APIMethod != "" {
				APIMethod = methods.MediaGroup
			}
		}
	}
	return APIMethod, err
}

func requiredMessage(msg *Message, APIMethod string) (string, error) {
	var err error
	if msg.fm.ch.ID == 0 {
		err = errors.MissedRequiredField("Chat", "ChatID", 0, false)
	}
	if msg.fm.m.Text == "" {
		err = errors.MissedRequiredField("Message", "Text", 0, false)
	}
	if APIMethod == "" {
		APIMethod = methods.Message
		msg.fm.contenttype = "application/json"
	}
	return APIMethod, err
}

func mediaPart(msg *Message) error {
	var err error

	if msg.fm.mh.amount > 0 {
		err = uniqueMedia(msg)
	} else {
		err = mediaGroup(msg)
	}
	return err
}

func messagePart(msg *Message) error {
	var err error
	var bytes []byte

	if msg.fm.mh.evenone {
		err = message.MultipartFields(msg.fm.m, msg.fm.writer)
	} else {

		bytes, err = message.CreateJSON(msg.fm.m)

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

		if msg.fm.mh.evenone {
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

	if msg.fm.mh.evenone {
		err = chat.MultipartFields(msg.fm.ch, msg.fm.writer)
	} else {

		bytes, err = chat.CreateJson(msg.fm.ch)

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

func makeRequest(msg *Message) (string, error) {
	var err error
	var APIMethod string
	var shouldSkip [4]bool
	// chat part must be skipped by default
	// it isn't skipped now, because of tests
	doPlan := [4]func(*Message) error{mediaPart, messagePart, keyboardPart, chatPart}

	msg.fm.buf = bytes.NewBuffer(nil)
	msg.fm.writer = multipart.NewWriter(msg.fm.buf)

	if msg.fm.mh.amount > 0 {
		APIMethod, err = requiredMedia(msg)
	} else {
		shouldSkip[0] = true
	}

	if err == nil {
		APIMethod, err = requiredMessage(msg, APIMethod)
	}

	for i := 0; (i < len(doPlan)) && (err == nil); i++ {
		if !shouldSkip[i] {
			err = doPlan[i](msg)
		}
	}
	if err == nil && !msg.fm.mh.evenone {
		err = uniteEverything(msg)
	}
	return APIMethod, err
}

func sendRequest(msg *Message, tg *types.Telegram, APIMethod string) error {
	var resp *http.Response
	var body []byte
	log.Print(msg.fm.buf.String())
	log.Print(APIMethod)
	log.Print(msg.fm.contenttype)
	url := fmt.Sprint(types.TelegramAPI, testbotdata.Token, "/", APIMethod)
	log.Print(url)
	req, err := http.NewRequest("POST", url, msg.fm.buf)
	req.Header.Set("Content-Type", msg.fm.contenttype)

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
				err = json.Unmarshal(body, tg)
				if err != nil {
					err = errors.CantUnmarshal(err)
				}
			}
		}
	}
	return err
}

// Sends the message you created. May send you an error if you made a mistake during the build process.
// Also sends a new pointer to the structure. This is a quick response from Telegram to your message. It collects some data
// about the message sent. For example: if you want to send a photo, this structure will have data about it from Telegram (id, size, etc.)
func (msg *Message) Send() (*types.Telegram, error) {
	var err error
	var tg *types.Telegram
	var APIMethod string

	if APIMethod, err = makeRequest(msg); err == nil {
		err = sendRequest(msg, tg, APIMethod)
	}

	return tg, err
}
