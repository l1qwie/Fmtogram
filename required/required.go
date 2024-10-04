package required

import "github.com/l1qwie/Fmtogram/formatter/methods"

const (
	ChatID     string = "chat_id"
	Text       string = "text"
	Photo      string = "photo"
	Video      string = "video"
	Audio      string = "audio"
	Document   string = "document"
	Animation  string = "animation"
	Voice      string = "voice"
	VideoNote  string = "video_note"
	MediaGroup string = "media"
)

var Fields map[string][7]string

func MakeRequiredFields() {
	Fields[methods.Message] = [7]string{ChatID, Text}
	Fields[methods.Photo] = [7]string{ChatID, Photo}
	Fields[methods.Video] = [7]string{ChatID, Video}
	Fields[methods.Audio] = [7]string{ChatID, Audio}
	Fields[methods.Document] = [7]string{ChatID, Document}
	Fields[methods.Animation] = [7]string{ChatID, Animation}
	Fields[methods.Voice] = [7]string{ChatID, Voice}
	Fields[methods.VideoNote] = [7]string{ChatID, VideoNote}
	Fields[methods.MediaGroup] = [7]string{ChatID, MediaGroup}
}
