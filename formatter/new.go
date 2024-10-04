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

// Creates a new "photo" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "media" package
func (m *Message) NewPhoto() *media.Photo {
	logs.NewObjectCreated("Photo")
	return &media.Photo{Type: "photo"}
}

// Creates a new "document" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "media" package
func (m *Message) NewDocument() *media.Document {
	logs.NewObjectCreated("Document")
	return &media.Document{Type: "document"}
}

// Creates a new "audio" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "media" package
func (m *Message) NewAudio() *media.Audio {
	logs.NewObjectCreated("Audio")
	return &media.Audio{Type: "audio"}
}

// Creates a new "video" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "media" package
func (m *Message) NewVideo() *media.Video {
	logs.NewObjectCreated("Video")
	return &media.Video{Type: "video"}
}

// Creates a new "message" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "message" package
func (m *Message) NewMessage() *message.Message {
	logs.NewObjectCreated("Message")
	return new(message.Message)
}

// Creates a new "chat" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "chat" package
func (m *Message) NewChat() *chat.Chat {
	logs.NewObjectCreated("Chat")
	return new(chat.Chat)
}

// Creates a new "inline keyboard" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "keyboard" package
func (m *Message) NewInlineKeyboard() *keyboard.Inline {
	logs.NewObjectCreated("Inline Keyboard")
	return new(keyboard.Inline)
}

// Creates a new "reply keyboard" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "keyboard" package
func (m *Message) NewReplyKeyboard() *keyboard.Reply {
	logs.NewObjectCreated("Reply Keyboard")
	return new(keyboard.Reply)
}
