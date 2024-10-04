package formatter

import (
	"bytes"
	"mime/multipart"
	"os"

	"github.com/l1qwie/Fmtogram/formatter/chat"
	"github.com/l1qwie/Fmtogram/formatter/keyboard"
	"github.com/l1qwie/Fmtogram/formatter/media"
	"github.com/l1qwie/Fmtogram/formatter/message"
)

type mediaHolder struct {
	storage [10]media.Handler
	i       int
	amount  int
	evenone bool
}

type formatter struct {
	m           *message.Message
	ch          *chat.Chat
	kb          keyboard.Handler
	contenttype string
	file        *os.File
	writer      *multipart.Writer
	path        string
	mh          *mediaHolder
	method      string
	mediajson   []byte
	buf         *bytes.Buffer
}

type Message struct {
	fm *formatter
}
