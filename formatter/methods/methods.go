package methods

const (
	Photo      string = "sendPhoto"
	Audio      string = "sendAudio"
	Document   string = "sendDocument"
	Video      string = "sendVideo"
	Animation  string = "sendAnimation"
	Voice      string = "sendVoice"
	VideoNote  string = "sendVideoNote"
	MediaGroup string = "sendMediaGroup"
	Message    string = "sendMessage"
)

var Media = []string{Photo, Audio, Document, Video, Animation, Voice, VideoNote, MediaGroup}
