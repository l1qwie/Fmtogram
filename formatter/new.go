package formatter

import (
	"github.com/l1qwie/Fmtogram/logs"
)

// Creates a new photo stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (m *Message) NewPhoto() IPhoto {
	logs.NewInterface("IPhoto")
	return &photo{Type: "photo"}
}

// Creates a new document stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (m *Message) NewDocument() IDocument {
	logs.NewObjectCreated("Document")
	return &document{Type: "document"}
}

// Creates a new audio stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (m *Message) NewAudio() IAudio {
	logs.NewObjectCreated("Audio")
	return &audio{Type: "audio"}
}

// Creates a new video stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (m *Message) NewVideo() IVideo {
	logs.NewObjectCreated("Video")
	return &video{Type: "video"}
}

// Creates a new message(inf) stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (m *Message) NewMessage() MSGInformation {
	logs.NewObjectCreated("Message")
	return &information{}
}

// Creates a new chat stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (m *Message) NewChat() IChat {
	logs.NewObjectCreated("Chat")
	return &chat{}
}

// Creates a new inline-keyboard stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (m *Message) NewInlineKeyboard() IInline {
	logs.NewObjectCreated("Inline Keyboard")
	return &inline{}
}

// Creates a new reply-keyboard stucture and returns an interface with connected functions. This is the only one way to add your data to an object
func (m *Message) NewReplyKeyboard() IReply {
	logs.NewObjectCreated("Reply Keyboard")
	return &reply{}
}
