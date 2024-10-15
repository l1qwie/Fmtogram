package formatter

import (
	"github.com/l1qwie/Fmtogram/logs"
)

// Added a photo to the message you're building
func (msg *Message) AddPhoto(ph IPhoto) {
	msg.fm.mh.storage[msg.fm.mh.i] = ph.(*photo)
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.InterfaceSaved("Photo")
}

// Added a video to the message you're building
func (msg *Message) AddVideo(vd IVideo) {
	msg.fm.mh.storage[msg.fm.mh.i] = vd.(*video)
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.InterfaceSaved("Video")
}

// Added a document to the message you're building
func (msg *Message) AddDocument(dc IDocument) {
	msg.fm.mh.storage[msg.fm.mh.i] = dc.(*document)
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.InterfaceSaved("Document")
}

// Added an audio to the message you're building
func (msg *Message) AddAudio(ad IAudio) {
	msg.fm.mh.storage[msg.fm.mh.i] = ad.(*audio)
	msg.fm.mh.i++
	msg.fm.mh.amount++
	logs.InterfaceSaved("Audio")
}

// Added a message to the message you're building
func (m *Message) AddMessage(inf MSGInformation) {
	m.fm.inf = inf.(*information)
	logs.InterfaceSaved("Message")
}

// Added a chat to the message you're building
func (m *Message) AddChat(ch IChat) {
	m.fm.ch = ch.(*chat)
	logs.ObjectSaved("Chat")
}
