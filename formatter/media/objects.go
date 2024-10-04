package media

import (
	"mime/multipart"

	"github.com/l1qwie/Fmtogram/types"
)

type Handler interface {
	MultipartFields(*multipart.Writer, *[10]interface{}, int, bool) error
	JsonFileds() ([]byte, error)
}

type Photo struct {
	Type  string `json:"type,omitempty"`
	Media string `json:"media,omitempty"`
	Photo string `json:"photo,omitempty"`

	// It's going to work only if you want to send more than 1 photo. Overwise use message object
	Caption string `json:"caption,omitempty"`
	// It's going to work only if you want to send more than 1 photo. Overwise use message object
	ParseMode string `json:"parse_mode,omitempty"`

	CaptionEntities       []*types.MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                   `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool                   `json:"has_spoiler,omitempty"`
	GottenFrom            int
}

type Video struct {
	Type      string `json:"type,omitempty"`
	Media     string `json:"media,omitempty"`
	Video     string `json:"video,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`

	// It's going to work only if you want to send more than 1 video. Overwise use message object
	Caption string `json:"caption,omitempty"`
	// It's going to work only if you want to send more than 1 video. Overwise use message object
	ParseMode string `json:"parse_mode,omitempty"`

	CaptionEntities       []*types.MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                   `json:"show_caption_above_media,omitempty"`
	Width                 int                    `json:"width,omitempty"`
	Height                int                    `json:"height,omitempty"`
	Duration              int                    `json:"duration,omitempty"`
	SupportsStreaming     bool                   `json:"supports_streaming,omitempty"`
	HasSpoiler            bool                   `json:"has_spoiler,omitempty"`
	VideoGottenFrom       int
	ThumbnailGottenFrom   int
}

type Audio struct {
	Type      string `json:"type,omitempty"`
	Media     string `json:"media,omitempty"`
	Audio     string `json:"audio,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`

	// It's going to work only if you want to send more than 1 audio. Overwise use message object
	Caption string `json:"caption,omitempty"`
	// It's going to work only if you want to send more than 1 audio. Overwise use message object
	ParseMode string `json:"parse_mode,omitempty"`

	CaptionEntities     []*types.MessageEntity `json:"caption_entities,omitempty"`
	Duration            int                    `json:"duration,omitempty"`
	Performer           string                 `json:"performer,omitempty"`
	Title               string                 `json:"title,omitempty"`
	AudioGottenFrom     int
	ThumbnailGottenFrom int
}

type Document struct {
	Type      string `json:"type,omitempty"`
	Media     string `json:"media,omitempty"`
	Document  string `json:"document,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`

	// It's going to work only if you want to send more than 1 document. Overwise use message object
	Caption string `json:"caption,omitempty"`
	// It's going to work only if you want to send more than 1 document. Overwise use message object
	ParseMode string `json:"parse_mode,omitempty"`

	CaptionEntities             []*types.MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                   `json:"disable_content_type_detection,omitempty"`
	DocumentGottenFrom          int
	ThumbnailGottenFrom         int
}

type MediaArray struct {
	Media []interface{} `json:"media,omitempty"`
}
