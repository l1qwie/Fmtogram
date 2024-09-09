package formatter

import (
	"mime/multipart"
	"os"

	"github.com/l1qwie/Fmtogram/formatter/media"
	"github.com/l1qwie/Fmtogram/formatter/message"
	"github.com/l1qwie/Fmtogram/types"
)

const (
	EDIT   int = 1
	DELETE int = 2
)

type InlineKeyboard struct {
	Keyboard [][]btn `json:"inline_keyboard"`
	x        int
	y        int
}

type InputMediaPhoto struct {
	Type      string `json:"type"`
	Media     string `json:"media"`
	Caption   string `json:"caption"`
	ParseMode string `json:"parse_mode"`
}

type InputMedia struct {
	Caption               string                `json:"caption,omitempty"`
	Type                  string                `json:"type,omitempty"`
	Media                 string                `json:"media,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool                  `json:"has_spoiler,omitempty"`
	CaptionEntities       []types.MessageEntity `json:"caption_entities,omitempty"`
}

type btnKind int

const (
	bCmd btnKind = 1
	bUrl btnKind = 2

	fromStorage  int = 0
	fromTelegram int = 1
	fromInternet int = 2
)

type btn struct {
	Label string `json:"text"`
	what  btnKind
	Cmd   string `json:"callback_data"`
	Url   string `json:"url"`
}

type SendMessage struct {
	BusinessConnectionID string                   `json:"business_connection_id,omitempty"`
	ChatID               interface{}              `json:"chat_id"`
	MessageThreadID      int                      `json:"message_thread_id,omitempty"`
	Text                 string                   `json:"text"`
	ParseMode            string                   `json:"parse_mode,omitempty"`
	Entities             []types.MessageEntity    `json:"entities,omitempty"`
	LinkPreviewOptions   types.LinkPreviewOptions `json:"link_preview_options,omitempty"`
	DisableNotification  bool                     `json:"disable_notification,omitempty"`
	ProtectContent       bool                     `json:"protect_content,omitempty"`
	MessageEffectID      string                   `json:"message_effect_id,omitempty"`
	ReplyParameters      types.ReplyParameters    `json:"reply_parameters,omitempty"`
	ReplyMarkup          interface{}              `json:"reply_markup,omitempty"`
}

type photo struct {
	fm *Formatter
}

type audio struct {
	fm *Formatter
}

type document struct {
	fm *Formatter
}

type video struct {
	fm *Formatter
}

type animation struct {
	fm *Formatter
}

type voice struct {
	fm *Formatter
}

type videoNote struct {
	fm *Formatter
}

type message2 struct {
	Text                  string                `json:"text,omitempty"`
	Caption               string                `json:"caption,omitempty"`
	MessageID             interface{}           `json:"message_id,omitempty"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	ProtectContent        bool                  `json:"protect_content,omitempty"`
	ReplyParameters       types.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup           interface{}           `json:"reply_markup,omitempty"`
	ThreadID              string                `json:"message_thread_id,omitempty"`
	CaptionEntities       []types.MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                  `json:"show_caption_above_media,omitempty"`
	EffectID              string                `json:"message_effect_id,omitempty"`
	action                int
}

type user struct {
	// userID int `json:""`
}

type chat struct {
	ChatID          interface{}           `json:"chat_id,omitempty"`
	ReplyParameters types.ReplyParameters `json:"reply_parameters,omitempty"`
}

type mediaStorage struct {
	mediaData string
}

type media2 struct {
	Group              []InputMedia `json:"media,omitempty"`
	Photo              string       `json:"photo,omitempty"`
	Audio              string       `json:"audio,omitempty"`
	Document           string       `json:"document,omitempty"`
	Video              string       `json:"video,omitempty"`
	Animation          string       `json:"animation,omitempty"`
	Voice              string       `json:"voice,omitempty"`
	VideoNote          string       `json:"videoNote,omitempty"`
	Thumbnail          mediaStorage `json:"thumbnail,omitempty"`
	DisableContentType bool         `json:"disable_content_type_detection,omitempty"`
	Width              int          `json:"width,omitempty"`
	Height             int          `json:"height,omitempty"`
	SupportsStreaming  bool         `json:"supports_streaming,omitempty"`
	Length             int          `json:"length,omitempty"`
	// group              []mediaStorage
	mediaFile mediaStorage
}

type business struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
}

type characteristics struct {
	HasSpoiler          bool   `json:"has_spoiler,omitempty"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ProtectContent      bool   `json:"protect_content,omitempty"`
	Duration            int    `json:"duration,omitempty"`
	Performer           string `json:"performer,omitempty"`
}

type Formatter struct {
	Message  message2
	Business business
	User     user
	Chat     chat
	Media    media2
	Char     characteristics
	m        *message.Message
	// Photo    media
	// Message        SendMessage
	// Keyboard    InlineKeyboard
	// Err         error
	contenttype string
	// kindofmedia  []int
	// mediatype    []string
	file         *os.File
	writer       *multipart.Writer
	path         string
	mediaStorage []media.Handler
	evenone      bool
}
