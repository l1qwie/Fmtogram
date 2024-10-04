package formatter

import (
	"github.com/l1qwie/Fmtogram/formatter/chat"
	"github.com/l1qwie/Fmtogram/formatter/media"
	"github.com/l1qwie/Fmtogram/formatter/message"
	"github.com/l1qwie/Fmtogram/logs"
)

// Added a photo to the message you're building
func (msg *Message) AddPhoto(photo media.Photo) {
	ph := photo
	msg.fm.mh.storage[msg.fm.mh.i] = &ph
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.ObjectSaved("Photo")
}

// Added a video to the message you're building
func (m *Message) AddVideo(video media.Video) {
	vd := video
	m.fm.mh.storage[m.fm.mh.i] = &vd
	m.fm.mh.i++
	logs.ObjectSaved("Video")
}

// Added a document to the message you're building
func (m *Message) AddDocument(document media.Document) {
	dc := document
	m.fm.mh.storage[m.fm.mh.i] = &dc
	m.fm.mh.i++
	logs.ObjectSaved("Document")
}

// Added an audio to the message you're building
func (m *Message) AddAudio(audio media.Audio) {
	ad := audio
	m.fm.mh.storage[m.fm.mh.i] = &ad
	m.fm.mh.i++
	logs.ObjectSaved("Audio")
}

// Added a message to the message you're building
func (m *Message) AddMessage(message message.Message) {
	mes := message
	m.fm.m = &mes
	logs.ObjectSaved("Message")
}

// Added a chat to the message you're building
func (m *Message) AddChat(chat chat.Chat) {
	ch := chat
	m.fm.ch = &ch
	logs.ObjectSaved("Chat")
}
