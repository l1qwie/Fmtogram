package formatter

import (
	"github.com/l1qwie/Fmtogram/formatter/chat"
	"github.com/l1qwie/Fmtogram/formatter/keyboard"
	"github.com/l1qwie/Fmtogram/formatter/media"
	"github.com/l1qwie/Fmtogram/formatter/message"
	"github.com/l1qwie/Fmtogram/logs"
)

const (
	replyKB  string = "Reply Keyboard"
	inlineKB string = "Inline Keyboard"
)

func (fm *Formatter) NewPhoto() *media.Photo {
	photo := &media.Photo{Type: "photo"}
	fm.mediaStorage = append(fm.mediaStorage, photo)
	logs.NewObjectCreated("Photo")
	return photo
}

func (fm *Formatter) NewDocument() *media.Document {
	document := &media.Document{Type: "document"}
	fm.mediaStorage = append(fm.mediaStorage, document)
	logs.NewObjectCreated("Document")
	return document
}

func (fm *Formatter) NewAudio() *media.Audio {
	audio := &media.Audio{Type: "audio"}
	fm.mediaStorage = append(fm.mediaStorage, audio)
	logs.NewObjectCreated("Audio")
	return audio
}

func (fm *Formatter) NewVideo() *media.Video {
	video := &media.Video{Type: "video"}
	fm.mediaStorage = append(fm.mediaStorage, video)
	logs.NewObjectCreated("Video")
	return video
}

func (fm *Formatter) NewMessage() *message.Message {
	fm.m = new(message.Message)
	logs.NewObjectCreated("Message")
	return fm.m
}

func (fm *Formatter) NewChat() *chat.Chat {
	fm.ch = new(chat.Chat)
	logs.NewObjectCreated("Chat")
	return fm.ch
}

func (fm *Formatter) NewInlineKeyboard() *keyboard.Inline {
	if fm.kb != nil {
		logs.DataIsntEmply(inlineKB, "keyboard", fm.kb)
	}
	keyboard := new(keyboard.Inline)
	fm.kb = keyboard
	logs.NewObjectCreated("Inline Keyboard")
	return keyboard
}

func (fm *Formatter) NewReplyKeyboard() *keyboard.Reply {
	if fm.kb != nil {
		logs.DataIsntEmply(replyKB, "keyboard", fm.kb)
	}
	keyboard := new(keyboard.Reply)
	fm.kb = keyboard
	logs.NewObjectCreated("Reply Keyboard")
	return keyboard
}
