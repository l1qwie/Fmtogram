package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/l1qwie/Fmtogram/logs"
)

const (
	photoMethod     string = "sendPhoto"
	audipMethod     string = "sendAudio"
	documentMethod  string = "sendDocument"
	videoMethod     string = "sendVideo"
	animationMethod string = "sendAnimation"
	voiceMethod     string = "sendVoice"
	videoNoteMethod string = "sendVideoNote"
)

type handleMedia interface {
	uniqueFields() (string, error)
}

type mediaFields struct {
	data   handleMedia
	method string
}

func (fm *Formatter) AddPhotoFromStorage(path string) {
	fm.Media.mediaFile.file.File = path
	fm.kindofmedia = []int{fromStorage}
	fm.mediatype = []string{"photo"}
	logs.DataWrittenSuccessfully("A Photo From The Storage")
}

func (fm *Formatter) AddPhotoFromTG(photoID string) {
	fm.Media.mediaFile.urlOrID = photoID
	fm.kindofmedia = []int{fromTelegram}
	fm.mediatype = []string{"photo"}
	logs.DataWrittenSuccessfully("A Photo Telegram ID")
}

func (fm *Formatter) AddPhotoFromInternet(URL string) {
	fm.Media.mediaFile.urlOrID = URL
	fm.kindofmedia = []int{fromInternet}
	fm.mediatype = []string{"photo"}
	logs.DataWrittenSuccessfully("A Photo URL From The Internet")
}
func (fm *Formatter) AddVideoFromStorage(path string) {
	fm.Media.mediaFile.file.File = path
	fm.kindofmedia = []int{fromStorage}
	fm.mediatype = []string{"video"}
	logs.DataWrittenSuccessfully("A Video From The Storage")
}

func (fm *Formatter) AddVideoFromTG(videoID string) {
	fm.Media.mediaFile.urlOrID = videoID
	fm.kindofmedia = []int{fromTelegram}
	fm.mediatype = []string{"video"}
	logs.DataWrittenSuccessfully("A Video Telegram ID")
}

func (fm *Formatter) AddVideoFromInternet(URL string) {
	fm.Media.mediaFile.urlOrID = URL
	fm.kindofmedia = []int{fromInternet}
	fm.mediatype = []string{"video"}
	logs.DataWrittenSuccessfully("A Video URL From The Internet")
}

func (fm *Formatter) AddAudioFromStorage(path string) {
	fm.Media.mediaFile.file.File = path
	fm.kindofmedia = []int{fromStorage}
	fm.mediatype = []string{"audio"}
	logs.DataWrittenSuccessfully("An Audio Path From The Storage")
}

func (fm *Formatter) AddAudioFromTG(audioID string) {
	fm.Media.mediaFile.urlOrID = audioID
	fm.kindofmedia = []int{fromTelegram}
	fm.mediatype = []string{"audio"}
	logs.DataWrittenSuccessfully("An Audio Telegram ID")
}

func (fm *Formatter) AddAudioFromInternet(URL string) {
	fm.Media.mediaFile.urlOrID = URL
	fm.kindofmedia = []int{fromInternet}
	fm.mediatype = []string{"audio"}
	logs.DataWrittenSuccessfully("An Audio URL From The Internet")
}

func (fm *Formatter) AddDocumentFromStorage(path string) {
	fm.Media.mediaFile.file.File = path
	fm.kindofmedia = []int{fromStorage}
	fm.mediatype = []string{"document"}
	logs.DataWrittenSuccessfully("A Document Path From The Storage")
}

func (fm *Formatter) AddDocumentFromTG(documentID string) {
	fm.Media.mediaFile.urlOrID = documentID
	fm.kindofmedia = []int{fromTelegram}
	fm.mediatype = []string{"document"}
	logs.DataWrittenSuccessfully("A Document Telegram ID")
}

func (fm *Formatter) AddDocumentFromInternet(URL string) {
	fm.Media.mediaFile.urlOrID = URL
	fm.kindofmedia = []int{fromInternet}
	fm.mediatype = []string{"document"}
	logs.DataWrittenSuccessfully("A Document URL From The Internet")
}

// func (fm *Formatter) AddMapOfMedia(arr []types.Media) {
// 	i := 0
// 	fm.Message.InputMedia = make([]types.Media, len(arr))
// 	for _, val := range arr {
// 		fm.Message.InputMedia[i].Type = val.Type
// 		fm.Message.InputMedia[i].Media = val.Media
// 		i++
// 	}
// }

// func (fm *Formatter) createMediaGroup(buf *bytes.Buffer) (string, error) {
// 	var (
// 		file    *os.File
// 		err     error
// 		part    io.Writer
// 		w, wrtr *multipart.Writer
// 	)
// 	bf := bytes.NewBuffer(nil)
// 	w = multipart.NewWriter(bf)
// 	for i := 0; i < len(fm.Message.InputMedia); i++ {
// 		file, err = os.Open(fm.Message.InputMedia[i].Media)
// 		if err == nil {
// 			err = w.WriteField("type", fm.Message.InputMedia[i].Type)
// 		}
// 		if err == nil {
// 			part, err = w.CreateFormFile("media", fm.Message.InputMedia[i].Media)
// 		}
// 		if err == nil {
// 			_, err = io.Copy(part, file)
// 		}
// 	}
// 	wrtr = multipart.NewWriter(buf)
// 	err = wrtr.WriteField("chat_id", fmt.Sprint(fm.Message.ChatID))
// 	if err == nil {
// 		fmt.Println(bf.String())
// 		fmt.Println()
// 		fmt.Println()
// 		fmt.Println()
// 		err = wrtr.WriteField("media", bf.String())
// 	}
// 	if err == nil {
// 		err = w.Close()
// 	}
// 	if err == nil {
// 		err = wrtr.Close()
// 	}

// 	return wrtr.FormDataContentType(), err
// }

func (fm *Formatter) checkTypeOfMedia() (*mediaFields, error) {
	var (
		ok  bool
		err error
	)

	acceptedTypes := make(map[string]*mediaFields, 7)
	acceptedTypes["photo"] = &mediaFields{data: &photo{fm}, method: photoMethod}
	acceptedTypes["audio"] = &mediaFields{data: &audio{fm}, method: audipMethod}
	acceptedTypes["document"] = &mediaFields{data: &document{fm}, method: documentMethod}
	acceptedTypes["video"] = &mediaFields{data: &video{fm}, method: videoMethod}
	acceptedTypes["animation"] = &mediaFields{data: &animation{fm}, method: animationMethod}
	acceptedTypes["voice"] = &mediaFields{data: &voice{fm}, method: voiceMethod}
	acceptedTypes["video-note"] = &mediaFields{data: &videoNote{fm}, method: videoNoteMethod}

	media, found := acceptedTypes[fm.mediatype[0]]

	if found {
		if fm.Media.mediaFile.urlOrID == "" {
			fm.path, ok = fm.Media.mediaFile.file.File.(string)
		} else if fm.mediatype[0] == "photo" {
			fm.Media.Photo = fm.Media.mediaFile.urlOrID
		} else if fm.mediatype[0] == "audio" {
			fm.Media.Audio = fm.Media.mediaFile.urlOrID
		} else if fm.mediatype[0] == "document" {
			fm.Media.Document = fm.Media.mediaFile.urlOrID
		} else if fm.mediatype[0] == "video" {
			fm.Media.Video = fm.Media.mediaFile.urlOrID
		} else if fm.mediatype[0] == "animation" {
			fm.Media.Animation = fm.Media.mediaFile.urlOrID
		} else if fm.mediatype[0] == "voice" {
			fm.Media.Voice = fm.Media.mediaFile.urlOrID
		} else if fm.mediatype[0] == "video-note" {
			fm.Media.VideoNote = fm.Media.mediaFile.urlOrID
		}
	}
	if !ok && fm.Media.mediaFile.urlOrID == "" {
		err = fmt.Errorf("the file path is supposted to be 'string' type, but there's something else. Please, make sure you give a 'string' type value")
	}
	return media, err
}

func writeToFile(file *os.File, writer *multipart.Writer, mediatype, path string) error {
	defer file.Close()
	part, err := writer.CreateFormFile(mediatype, path)
	if err == nil {
		_, err = io.Copy(part, file)
	}
	return err
}

func (fm *Formatter) durationWidthHeight(writer *multipart.Writer) error {
	var err error
	if fm.Char.Duration != 0 {
		err = writer.WriteField("duration", fmt.Sprintf("%d", fm.Char.Duration))
	}
	if err == nil && fm.Media.Width != 0 {
		err = writer.WriteField("width", fmt.Sprintf("%d", fm.Media.Width))
	}
	if err == nil && fm.Media.Height != 0 {
		err = writer.WriteField("height", fmt.Sprintf("%d", fm.Media.Height))
	}
	return err
}

func (fm *Formatter) showCapAbMediaHasSpoiler(writer *multipart.Writer) error {
	var err error
	if fm.Message.ShowCaptionAboveMedia {
		err = writer.WriteField("show_caption_above_media", "True")
	}
	if err == nil && fm.Char.HasSpoiler {
		err = writer.WriteField("has_spoiler", "True")
	}
	return err
}

func (fm *Formatter) thumbnail(file *os.File, writer *multipart.Writer) error {
	var (
		thumbpath string
		err       error
		part      io.Writer
	)
	ok := true
	if fm.Media.Thumbnail.urlOrID == "" {
		thumbpath, ok = fm.Media.Thumbnail.file.File.(string)
		log.Print(thumbpath, ok)
	} else {
		thumbpath = fm.Media.Thumbnail.urlOrID
	}
	if thumbpath != "" && ok {
		part, err = writer.CreateFormFile("thumbnail", thumbpath)
		if err == nil {
			_, err = io.Copy(part, file)
		}
	}
	return err
}

func (photo *photo) uniqueFields() (string, error) {
	err := writeToFile(photo.fm.file, photo.fm.writer, "photo", photo.fm.path)
	if err == nil {
		err = photo.fm.showCapAbMediaHasSpoiler(photo.fm.writer)
	}
	return photo.fm.writer.FormDataContentType(), err
}

func (audio *audio) uniqueFields() (string, error) {
	err := writeToFile(audio.fm.file, audio.fm.writer, "audio", audio.fm.path)
	if err == nil && audio.fm.Char.Duration != 0 {
		err = audio.fm.writer.WriteField("duration", fmt.Sprintf("%d", audio.fm.Char.Duration))
	}
	if err == nil && audio.fm.Char.Performer != "" {
		err = audio.fm.writer.WriteField("performer", audio.fm.Char.Performer)
	}
	if err == nil {
		err = audio.fm.thumbnail(audio.fm.file, audio.fm.writer)
	}
	return audio.fm.writer.FormDataContentType(), err
}

func (doc *document) uniqueFields() (string, error) {
	err := writeToFile(doc.fm.file, doc.fm.writer, "document", doc.fm.path)
	if err == nil {
		err = doc.fm.thumbnail(doc.fm.file, doc.fm.writer)
	}
	if err == nil && doc.fm.Media.DisableContentType {
		err = doc.fm.writer.WriteField("disable_content_type_detection", "True")
	}
	return doc.fm.writer.FormDataContentType(), err
}

func (video *video) uniqueFields() (string, error) {
	err := writeToFile(video.fm.file, video.fm.writer, "video", video.fm.path)
	if err == nil {
		err = video.fm.durationWidthHeight(video.fm.writer)
	}
	if err == nil {
		err = video.fm.thumbnail(video.fm.file, video.fm.writer)
	}
	if err == nil {
		err = video.fm.showCapAbMediaHasSpoiler(video.fm.writer)
	}
	if err == nil && video.fm.Media.SupportsStreaming {
		err = video.fm.writer.WriteField("supports_streaming", "True")
	}
	return video.fm.writer.FormDataContentType(), err
}

func (anim *animation) uniqueFields() (string, error) {
	err := writeToFile(anim.fm.file, anim.fm.writer, "animation", anim.fm.path)
	if err == nil {
		err = anim.fm.durationWidthHeight(anim.fm.writer)
	}
	if err == nil {
		err = anim.fm.thumbnail(anim.fm.file, anim.fm.writer)
	}
	if err == nil {
		err = anim.fm.showCapAbMediaHasSpoiler(anim.fm.writer)
	}
	return anim.fm.writer.FormDataContentType(), err
}

func (voice *voice) uniqueFields() (string, error) {
	err := writeToFile(voice.fm.file, voice.fm.writer, "voice", voice.fm.path)
	if err == nil && voice.fm.Char.Duration != 0 {
		err = voice.fm.writer.WriteField("duration", fmt.Sprintf("%d", voice.fm.Char.Duration))
	}
	return voice.fm.writer.FormDataContentType(), err
}

func (vn *videoNote) uniqueFields() (string, error) {
	err := writeToFile(vn.fm.file, vn.fm.writer, "video_note", vn.fm.path)
	if err == nil && vn.fm.Char.Duration != 0 {
		err = vn.fm.writer.WriteField("duration", fmt.Sprintf("%d", vn.fm.Char.Duration))
	}
	if err == nil && vn.fm.Media.Length != 0 {
		err = vn.fm.writer.WriteField("length", fmt.Sprintf("%d", vn.fm.Media.Length))
	}
	if err == nil {
		err = vn.fm.thumbnail(vn.fm.file, vn.fm.writer)
	}
	return vn.fm.writer.FormDataContentType(), err
}

func (fm *Formatter) commonMediaFields(writer *multipart.Writer) error {
	var err error
	if fm.Business.BusinessConnectionID != "" {
		err = writer.WriteField("business_connection_id", fm.Business.BusinessConnectionID)
	}
	if err == nil {
		err = writer.WriteField("chat_id", fmt.Sprintf("%d", fm.Chat.ChatID))
	}
	if err == nil && fm.Message.ThreadID != "" {
		err = writer.WriteField("message_thread_id", fm.Message.ThreadID)
	}
	if err == nil && fm.Message.Text != "" {
		err = writer.WriteField("caption", fm.Message.Text)
	}
	if err == nil && fm.Message.ParseMode != "" {
		err = writer.WriteField("parse_mode", fm.Message.ParseMode)
	}
	if err == nil && len(fm.Message.CaptionEntities) > 0 {
		body, err1 := json.Marshal(fm.Message.CaptionEntities)
		if err1 == nil {
			err = writer.WriteField("caption_entities", string(body))
		}
	}
	if err == nil && fm.Char.DisableNotification {
		err = writer.WriteField("disable_notification", "True")
	}
	if err == nil && fm.Char.ProtectContent {
		err = writer.WriteField("protect_content", "True")
	}
	if err == nil && fm.Message.EffectID != "" {
		err = writer.WriteField("message_effect_id", fm.Message.EffectID)
	}
	if err == nil && fm.Chat.ReplyParameters.MessageID != 0 {
		body, err1 := json.Marshal(fm.Chat.ReplyParameters)
		if err1 == nil {
			err = writer.WriteField("reply_parameters", string(body))
		}
	}
	if err == nil && fm.Message.ReplyMarkup != nil {
		body, err1 := json.Marshal(fm.Message.ReplyMarkup)
		if err1 == nil {
			err = writer.WriteField("reply_markup", string(body))
		}
	}
	return err
}

func (fm *Formatter) fromStorageMedia(buf *bytes.Buffer) (string, string, error) {
	fm.file = new(os.File)
	fm.writer = multipart.NewWriter(buf)
	media, err := fm.checkTypeOfMedia()
	if err == nil {
		fm.file, err = os.Open(fm.path)
	}
	if err == nil {
		fm.contenttype, err = media.data.uniqueFields()
	}
	if err == nil {
		err = fm.commonMediaFields(fm.writer)
	}
	if err == nil {
		defer fm.writer.Close()
	}
	return media.method, fm.writer.FormDataContentType(), err
}

func (fm *Formatter) tgOrURLMedia(buf *bytes.Buffer) (string, string, error) {
	var (
		body1, body2, body3, finalbody []byte
	)
	map1 := make(map[string]interface{})
	map2 := make(map[string]interface{})
	map3 := make(map[string]interface{})
	finalmap := make(map[string]interface{})
	media, err := fm.checkTypeOfMedia()
	if err == nil {
		if fm.Message.Text != "" {
			fm.Message.Caption = fm.Message.Text
		}
		body1, err = json.Marshal(fm.Message)
		body2, err = json.Marshal(fm.Chat)
		body3, err = json.Marshal(fm.Media)
		if err == nil {
			err = json.Unmarshal(body1, &map1)
			err = json.Unmarshal(body2, &map2)
			err = json.Unmarshal(body3, &map3)
			if err == nil {
				for k, v := range map1 {
					finalmap[k] = v
				}
				for k, v := range map2 {
					finalmap[k] = v
				}
				for k, v := range map3 {
					finalmap[k] = v
				}
				finalbody, err = json.Marshal(&finalmap)
				if err == nil {
					buf.Write(finalbody)
				}
			}
		}
		log.Print(string(finalbody))
	}
	return media.method, "application/json", err
}

// func (fm *Formatter) PrepareMediaForEdit(buf *bytes.Buffer) (string, error) {
// 	var (
// 		file   *os.File
// 		part   io.Writer
// 		writer *multipart.Writer
// 		err    error
// 		media  string
// 	)
// 	writer = multipart.NewWriter(buf)
// 	if fm.mediatype[0] == "photo" {
// 		media = fm.Message.Photo
// 	} else if fm.mediatype[0] == "video" {
// 		media = fm.Message.Video
// 	}

// 	file, err = os.Open(media)
// 	if err == nil {
// 		part, err = writer.CreateFormFile(fm.mediatype[0], media)
// 	}
// 	if err == nil {
// 		_, err = io.Copy(part, file)
// 	}
// 	if err == nil {
// 		err = writer.WriteField("caption", fm.Message.Text)
// 	}
// 	if err == nil {
// 		err = writer.WriteField("parse_mode", fm.Message.ParseMode)
// 	}
// 	if err == nil {
// 		err = writer.Close()
// 	}
// 	file.Close()
// 	return writer.FormDataContentType(), err
// }
