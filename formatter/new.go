package formatter

import (
	"github.com/l1qwie/Fmtogram/logs"
)

// Creates a new "photo" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "media" package
func (m *Message) NewPhoto() IPhoto {
	logs.NewInterface("IPhoto")
	return &photo{Type: "photo"}
}

// Creates a new "document" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "media" package
func (m *Message) NewDocument() IDocument {
	logs.NewObjectCreated("Document")
	return &document{Type: "document"}
}

// Creates a new "audio" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "media" package
func (m *Message) NewAudio() IAudio {
	logs.NewObjectCreated("Audio")
	return &audio{Type: "audio"}
}

// Creates a new "video" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "media" package
func (m *Message) NewVideo() IVideo {
	logs.NewObjectCreated("Video")
	return &video{Type: "video"}
}

// Creates a new "message" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "message" package
func (m *Message) NewMessage() MSGInformation {
	logs.NewObjectCreated("Message")
	return &information{}
}

// Creates a new "chat" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "chat" package
func (m *Message) NewChat() IChat {
	logs.NewObjectCreated("Chat")
	return &chat{}
}

// Creates a new "inline keyboard" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "keyboard" package
func (m *Message) NewInlineKeyboard() IInline {
	logs.NewObjectCreated("Inline Keyboard")
	return new(inline)
}

// Creates a new "reply keyboard" object. The preferred way to create a project if you need to see the logs.
// If you don't need that, you can create the object directly from "keyboard" package
func (m *Message) NewReplyKeyboard() IReply {
	logs.NewObjectCreated("Reply Keyboard")
	return new(reply)
}
