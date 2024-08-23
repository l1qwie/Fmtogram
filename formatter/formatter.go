package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/l1qwie/Fmtogram/errors"
	"github.com/l1qwie/Fmtogram/executer"
	"github.com/l1qwie/Fmtogram/logs"
	"github.com/l1qwie/Fmtogram/types"
)

// Save a text of the message for sending
func (fm *Formatter) WriteString(text string) {
	logs.DataWrittenSuccessfully("Text Of The Message")
	fm.Message.Text = text
}

// Save a Chat name of a chat for sending
func (fm *Formatter) WriteChatName(chatname string) {
	logs.DataWrittenSuccessfully("Chatname")
	fm.Message.ChatID = fmt.Sprint("@", chatname)
}

// Save a Chat ID of a chat for sending
func (fm *Formatter) WriteChatId(chatID int) {
	logs.DataWrittenSuccessfully("ChatID")
	fm.Message.ChatID = chatID
}

// Save a Parse Mode of the message for sending
func (fm *Formatter) WriteParseMode(mode string) {
	logs.DataWrittenSuccessfully("Parse Mode")
	fm.Message.ParseMode = mode
}

// Save a Deleted Message ID for future deleting
func (fm *Formatter) WriteDeleteMesId(mesId int) {
	logs.DataWrittenSuccessfully("Deleted Message ID")
	fm.DeletedMessage.MessageId = mesId
}

// Save an Edited Message ID for future editing
func (fm *Formatter) WriteEditMesId(mesId int) {
	logs.DataWrittenSuccessfully("Edited Message ID")
	fm.Message.MessageId = mesId
}

// Save an Error to handle it soon
func (fm *Formatter) Error(err error) {
	logs.DataWrittenSuccessfully("Error")
	fm.Err = err
}

func (fm *Formatter) Complete() error {
	if fm.Err != nil {
		errors.MadeMisstake(fm.Err)
	}
	return fm.Err
}

func (fm *Formatter) CheckDelete() (err error) {
	var (
		function    string
		jsonMessage []byte
		finalBuffer *bytes.Buffer
	)
	if fm.DeletedMessage.MessageId != 0 {
		function = "deleteMessage"
		if chatid, ok := fm.Message.ChatID.(int); ok {
			fm.DeletedMessage.ChatId = chatid
		}
		jsonMessage, err = json.Marshal(fm.DeletedMessage)
		if err == nil {
			fm.contenttype = "application/json"
			finalBuffer = bytes.NewBuffer(jsonMessage)
		}
		if err == nil {
			_ = executer.Send(finalBuffer, function, fm.contenttype, false)
		}
	}
	return err
}

func (fm *Formatter) ReadyKB() {
	if fm.Keyboard.Keyboard != nil {
		jsonKeyboard, err := json.Marshal(fm.Keyboard)
		if err == nil {
			fm.Message.ReplyMarkup = string(jsonKeyboard)
		}
	}
}

func (fm *Formatter) makebuf() (*bytes.Buffer, error) {
	var buf *bytes.Buffer
	jsonMessage, err := json.Marshal(fm.Message)
	if err == nil {
		fm.contenttype = "application/json"
		buf = bytes.NewBuffer(jsonMessage)
	}
	return buf, err
}

func (fm *Formatter) sendMessage() (*bytes.Buffer, string, error) {
	log.Print("In sendMessage()")
	buf, err := fm.makebuf()
	return buf, "sendMessage", err
}

func (fm *Formatter) editMessage() (*bytes.Buffer, string, error) {
	buf, err := fm.makebuf()
	return buf, "editMessageText", err
}

func (fm *Formatter) mediaGroup() error {
	buf, err := fm.makebuf()
	_ = executer.Send(buf, "sendMediaGroup", "application/json", true)
	return err
}

func (fm *Formatter) media() (*bytes.Buffer, string, error) {
	var (
		buf *bytes.Buffer
		err error
		f   string
	)
	if fm.kindofmedia[0] == fromStorage {
		buf = bytes.NewBuffer(nil)
		fm.contenttype, err = fm.prepareMedia(buf)
	} else {
		if fm.mediatype[0] == "photo" {
			f = "sendPhoto"
		} else {
			f = "sendVideo"
		}
		buf, err = fm.makebuf()
	}
	return buf, f, err
}

func (fm *Formatter) Make() (*types.Returned, error) {
	var (
		buf     *bytes.Buffer
		f       string
		res     *types.Returned
		mshstat bool
	)
	err := fm.CheckDelete()
	if err == nil {
		fm.ReadyKB()
	}
	if err == nil {
		if len(fm.Message.InputMedia) == 0 && fm.Message.Photo == "" && fm.Message.Video == "" {
			mshstat = true
			if fm.Message.MessageId == 0 {
				buf, f, err = fm.sendMessage()
			} else {
				buf, f, err = fm.editMessage()
			}
		} else if len(fm.Message.InputMedia) != 0 || fm.Message.Photo != "" || fm.Message.Video != "" {
			if len(fm.Message.InputMedia) > 1 {
				err = fm.mediaGroup()
				if err == nil {
					mshstat = true
					buf, f, err = fm.sendMessage()
				}
			} else {
				buf, f, err = fm.media()
			}
		}
	}
	if err == nil {
		log.Print("THE MESSAGE ID DELETE OR EDIT AND MARSHAL STATUS: ", fm.DeletedMessage.MessageId, fm.Message.MessageId, mshstat)
		res = executer.Send(buf, f, fm.contenttype, mshstat)
	}

	return res, err
}

func (fm *Formatter) SendToChannel() error {
	log.Print("In SendToChannel()")
	fm.ReadyKB()
	buf, f, err := fm.sendMessage()
	if err == nil {
		executer.Send(buf, f, fm.contenttype, true)
	} else {
		log.Printf("ERROR after fm.sendMessage(): %v", err)
	}
	return err
}

func (fm *Formatter) DeleteMessage() {
	finalBuffer := new(bytes.Buffer)
	function := "deleteMessage"
	if fm.DeletedMessage.MessageId != nil {
		if fm.DeletedMessage.ChatId != nil {
			jsonMessage, err := json.Marshal(fm.DeletedMessage)
			if err == nil {
				fm.contenttype = "application/json"
				finalBuffer = bytes.NewBuffer(jsonMessage)
			}
			if err == nil {
				_ = executer.Send(finalBuffer, function, fm.contenttype, false)
			}
		} else {
			errors.ChatIDIsNil()
		}
	} else {
		errors.MessageIDIsNil()
	}
}

// func (fm *Formatter) EditMessage() {
// 	finalBuffer := new(bytes.Buffer)
// 	function := "editMessageText"
// }

func (fm *Formatter) Send() (mes *types.Returned, err error) {
	var (
		jsonMessage   []byte
		finalBuffer   *bytes.Buffer
		function      string
		marshalstatus bool
	)

	err = fm.CheckDelete()
	if err == nil {
		fm.ReadyKB()
	}
	if err == nil {
		if len(fm.Message.InputMedia) == 0 && fm.Message.Photo == "" && fm.Message.Video == "" {
			if fm.Message.MessageId == 0 {
				marshalstatus = true
				function = "sendMessage"
			} else {
				function = "editMessageText"
			}
			jsonMessage, err = json.Marshal(fm.Message)
			if err == nil {
				fm.contenttype = "application/json"
				finalBuffer = bytes.NewBuffer(jsonMessage)
			}
		} else if len(fm.Message.InputMedia) != 0 || fm.Message.Photo != "" || fm.Message.Video != "" {
			marshalstatus = true
			if len(fm.Message.InputMedia) > 1 {
				function = "sendMediaGroup"
				m := fm.Message.Text
				fm.Message.Text = ""
				//finalBuffer = bytes.NewBuffer(nil)
				jsonMessage, err = json.Marshal(fm.Message)
				if err == nil {
					fm.contenttype = "application/json"
					finalBuffer = bytes.NewBuffer(jsonMessage)
				}
				//fm.contenttype, err = fm.createMediaGroup(finalBuffer)
				_ = executer.Send(finalBuffer, function, fm.contenttype, marshalstatus)
				fm.Message.InputMedia = []types.Media{{}}
				fm.Message.Text = m
				function = "sendMessage"
			} else {
				if fm.mediatype[0] == "photo" {
					function = "sendPhoto"
				} else if fm.mediatype[0] == "video" {
					function = "sendVideo"
				}
				if fm.kindofmedia[0] == fromStorage {
					finalBuffer = bytes.NewBuffer(nil)
					fm.contenttype, err = fm.prepareMedia(finalBuffer)
				} else {
					jsonMessage, err = json.Marshal(fm.Message)
					if err == nil {
						fm.contenttype = "application/json"
						finalBuffer = bytes.NewBuffer(jsonMessage)
					}
				}
			}
		}

	}
	if err == nil {
		mes = executer.Send(finalBuffer, function, fm.contenttype, marshalstatus)
	}

	return mes, err
}
