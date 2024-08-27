package formatter

import "github.com/l1qwie/Fmtogram/types"

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
	InputMediaPhoto InputMediaPhoto `json:"InputMediaPhoto"`
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
	BusinessConnectionID string                    `json:"business_connection_id,omitempty"`
	ChatID               interface{}               `json:"chat_id"`
	MessageThreadID      int                       `json:"message_thread_id,omitempty"`
	Text                 string                    `json:"text"`
	ParseMode            string                    `json:"parse_mode,omitempty"`
	Entities             []*types.MessageEntity    `json:"entities,omitempty"`
	LinkPreviewOptions   *types.LinkPreviewOptions `json:"link_preview_options,omitempty"`
	DisableNotification  bool                      `json:"disable_notification,omitempty"`
	ProtectContent       bool                      `json:"protect_content,omitempty"`
	MessageEffectID      string                    `json:"message_effect_id,omitempty"`
	ReplyParameters      *types.ReplyParameters    `json:"reply_parameters,omitempty"`
	ReplyMarkup          interface{}               `json:"reply_markup,omitempty"`
}

type message struct {
	text                  string                 `json:"text"`
	messageID             interface{}            `json:"message_id"`
	parseMode             string                 `json:"parse_mode,omitempty"`
	protectContent        bool                   `json:"protect_content,omitempty"`
	replyParameters       *types.ReplyParameters `json:"reply_parameters,omitempty"`
	replyMarkup           interface{}            `json:"reply_markup,omitempty"`
	threadID              string                 `json:"message_thread_id,omitempty"`
	captionEntities       []*types.MessageEntity `json:"caption_entities,omitempty"`
	showCaptionAboveMedia bool                   `json:"show_caption_above_media,omitempty"`
	effectID              string                 `json:"message_effect_id,omitempty"`
	action                int
}

type user struct {
	// userID int `json:""`
}

type chat struct {
	chatID               interface{}            `json:"chat_id"`
	businessConnectionID string                 `json:"business_connection_id,omitempty"`
	replyParameters      *types.ReplyParameters `json:"reply_parameters,omitempty"`
}

type mediaType struct {
	file    *types.InputFile
	urlOrID string
}

type media struct {
	photo              *mediaType `json:"photo,omitempty"`
	audio              *mediaType `json:"audio,omitempty"`
	document           *mediaType `json:"document,omitempty"`
	video              *mediaType `json:"video,omitempty"`
	thumbnail          *mediaType `json:"thumbnail,omitempty"`
	disableContentType bool       `json:"disable_content_type_detection,omitempty"`
	width              int        `json:"width,omitempty"`
	height             int        `json:"height,omitempty"`
	supportsStreaming  bool       `json:"supports_streaming,omitempty"`
	length             int        `json:"length,omitempty"`
}

type business struct {
	businessConnectionID string `json:"business_connection_id,omitempty"`
}

type characteristics struct {
	hasSpoiler          bool   `json:"has_spoiler,omitempty"`
	disableNotification bool   `json:"disable_notification,omitempty"`
	protectContent      bool   `json:"protect_content,omitempty"`
	duration            int    `json:"duration,omitempty"`
	performer           string `json:"performer,omitempty"`
}

type Formatter struct {
	message *message
	user    *user
	chat    *chat
	media   *media
	char    *characteristics
	// Message        SendMessage
	Keyboard    InlineKeyboard
	Err         error
	contenttype string
	kindofmedia []int
	mediatype   []string
}
