package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/l1qwie/Fmtogram/logs"
)

func (fm *Formatter) AddPhotoFromStorage(path string) {
	logs.DataWrittenSuccessfully("A Photo From The Storage")
	fm.media.photo.file.File = path
	fm.kindofmedia = []int{fromStorage}
	fm.mediatype = []string{"photo"}
}

func (fm *Formatter) AddPhotoFromTG(photoID string) {
	logs.DataWrittenSuccessfully("A Photo Telegram ID")
	fm.media.photo.urlOrID = photoID
	fm.kindofmedia = []int{fromTelegram}
	fm.mediatype = []string{"photo"}
}

func (fm *Formatter) AddPhotoFromInternet(URL string) {
	logs.DataWrittenSuccessfully("A Photo URL From The Internet")
	fm.media.photo.urlOrID = URL
	fm.kindofmedia = []int{fromInternet}
	fm.mediatype = []string{"photo"}
}
func (fm *Formatter) AddVideoFromStorage(path string) {
	logs.DataWrittenSuccessfully("A Video From The Storage")
	fm.media.video.file.File = path
	fm.kindofmedia = []int{fromStorage}
	fm.mediatype = []string{"video"}
}

func (fm *Formatter) AddVideoFromTG(videoID string) {
	logs.DataWrittenSuccessfully("A Video Telegram ID")
	fm.media.video.urlOrID = videoID
	fm.kindofmedia = []int{fromTelegram}
	fm.mediatype = []string{"video"}
}

func (fm *Formatter) AddVideoFromInternet(URL string) {
	logs.DataWrittenSuccessfully("A Video URL From The Internet")
	fm.media.video.urlOrID = URL
	fm.kindofmedia = []int{fromInternet}
	fm.mediatype = []string{"video"}
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

func (fm *Formatter) checkTypeOfMedia() (string, func(*os.File, *multipart.Writer, string) (string, string, error), error) {
	var (
		media string
		ok    bool
		err   error
	)
	acceptedTypes := make(map[string]func(*os.File, *multipart.Writer, string) (string, string, error), 7)
	acceptedTypes["photo"] = fm.fromStoragePhoto
	acceptedTypes["audio"] = fm.fromStorageAudio
	acceptedTypes["document"] = fm.fromStorageDocument
	acceptedTypes["video"] = fm.fromStorageVideo
	acceptedTypes["animation"] = fm.fromStorageAnimation
	acceptedTypes["voice"] = fm.fromStorageVoice
	acceptedTypes["video-note"] = fm.fromStorageVideoNote

	f, found := acceptedTypes[fm.mediatype[0]]

	if found {
		media, ok = fm.media.video.file.File.(string)
	}
	if !ok {
		err = fmt.Errorf("The file path is supposted to be 'string' type, but there's something else. Please, make sure you give a 'string' type value")
	}
	return media, f, err
}

func writeToFile(file *os.File, writer *multipart.Writer, mediatype, path string) error {
	part, err := writer.CreateFormFile(mediatype, path)
	if err == nil {
		_, err = io.Copy(part, file)
	}
	return err
}

func (fm *Formatter) durationWidthHeight(writer *multipart.Writer) error {
	var err error
	if fm.char.duration != 0 {
		err = writer.WriteField("duration", fmt.Sprintf("%d", fm.char.duration))
	}
	if err == nil && fm.media.width != 0 {
		err = writer.WriteField("width", fmt.Sprintf("%d", fm.media.width))
	}
	if err == nil && fm.media.height != 0 {
		err = writer.WriteField("height", fmt.Sprintf("%d", fm.media.height))
	}
	return err
}

func (fm *Formatter) showCapAbMediaHasSpoiler(writer *multipart.Writer) error {
	var err error
	if fm.message.showCaptionAboveMedia {
		err = writer.WriteField("show_caption_above_media", "True")
	}
	if err == nil && fm.char.hasSpoiler {
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
	if fm.media.thumbnail.urlOrID == "" {
		thumbpath = fm.media.thumbnail.file.File.(string)
	} else {
		thumbpath = fm.media.thumbnail.urlOrID
	}
	if thumbpath != "" {
		part, err = writer.CreateFormFile("thumbnail", thumbpath)
		if err == nil {
			_, err = io.Copy(part, file)
		}
	}
	return err
}

func (fm *Formatter) fromStoragePhoto(file *os.File, writer *multipart.Writer, path string) (string, string, error) {
	err := writeToFile(file, writer, "photo", path)
	if err == nil {
		err = fm.showCapAbMediaHasSpoiler(writer)
	}
	return writer.FormDataContentType(), "sendPhoto", err
}

func (fm *Formatter) fromStorageAudio(file *os.File, writer *multipart.Writer, path string) (string, string, error) {
	err := writeToFile(file, writer, "audio", path)
	if err == nil && fm.char.duration != 0 {
		err = writer.WriteField("duration", fmt.Sprintf("%d", fm.char.duration))
	}
	if err == nil && fm.char.performer != "" {
		err = writer.WriteField("performer", fm.char.performer)
	}
	if err == nil {
		err = fm.thumbnail(file, writer)
	}
	return writer.FormDataContentType(), "sendAudio", err
}

func (fm *Formatter) fromStorageDocument(file *os.File, writer *multipart.Writer, path string) (string, string, error) {
	err := writeToFile(file, writer, "document", path)
	if err == nil {
		err = fm.thumbnail(file, writer)
	}
	if err == nil && fm.media.disableContentType {
		err = writer.WriteField("disable_content_type_detection", "True")
	}
	return writer.FormDataContentType(), "sendDocument", err
}

func (fm *Formatter) fromStorageVideo(file *os.File, writer *multipart.Writer, path string) (string, string, error) {
	err := writeToFile(file, writer, "video", path)
	if err == nil {
		err = fm.durationWidthHeight(writer)
	}
	if err == nil {
		err = fm.thumbnail(file, writer)
	}
	if err == nil {
		err = fm.showCapAbMediaHasSpoiler(writer)
	}
	if err == nil && fm.media.supportsStreaming {
		err = writer.WriteField("supports_streaming", "True")
	}
	return writer.FormDataContentType(), "sendVideo", err
}

func (fm *Formatter) fromStorageAnimation(file *os.File, writer *multipart.Writer, path string) (string, string, error) {
	err := writeToFile(file, writer, "animation", path)
	if err == nil {
		err = fm.durationWidthHeight(writer)
	}
	if err == nil {
		err = fm.thumbnail(file, writer)
	}
	if err == nil {
		err = fm.showCapAbMediaHasSpoiler(writer)
	}
	return writer.FormDataContentType(), "sendAnimation", err
}

func (fm *Formatter) fromStorageVoice(file *os.File, writer *multipart.Writer, path string) (string, string, error) {
	err := writeToFile(file, writer, "voice", path)
	if err == nil && fm.char.duration != 0 {
		err = writer.WriteField("duration", fmt.Sprintf("%d", fm.char.duration))
	}
	return writer.FormDataContentType(), "sendVoice", err
}

func (fm *Formatter) fromStorageVideoNote(file *os.File, writer *multipart.Writer, path string) (string, string, error) {
	err := writeToFile(file, writer, "video_note", path)
	if err == nil && fm.char.duration != 0 {
		err = writer.WriteField("duration", fmt.Sprintf("%d", fm.char.duration))
	}
	if err == nil && fm.media.length != 0 {
		err = writer.WriteField("length", fmt.Sprintf("%d", fm.media.length))
	}
	if err == nil {
		err = fm.thumbnail(file, writer)
	}
	return writer.FormDataContentType(), "sendVideoNote", err
}

func (fm *Formatter) commonMediaFields(file *os.File, writer *multipart.Writer) error {
	var err error
	// file, err := os.Open(media)
	// if err == nil {
	// }
	if err == nil && fm.chat.businessConnectionID != "" {
		err = writer.WriteField("business_connection_id", fmt.Sprintf("%d", fm.chat.businessConnectionID))
	}
	if err == nil {
		err = writer.WriteField("chat_id", fmt.Sprintf("%d", fm.chat.chatID))
	}
	if err == nil && fm.message.threadID != "" {
		err = writer.WriteField("message_thread_id", fmt.Sprintf("%d", fm.message.threadID))
	}
	if err == nil && fm.message.text != "" {
		err = writer.WriteField("caption", fm.message.text)
	}
	if err == nil && fm.message.parseMode != "" {
		err = writer.WriteField("parse_mode", fm.message.parseMode)
	}
	if err == nil && len(fm.message.captionEntities) > 0 {
		body, err1 := json.Marshal(fm.message.captionEntities)
		if err1 == nil {
			err = writer.WriteField("caption_entities", string(body))
		}
	}
	if err == nil && fm.char.disableNotification {
		err = writer.WriteField("disable_notification", "True")
	}
	if err == nil && fm.char.protectContent {
		err = writer.WriteField("protect_content", "True")
	}
	if err == nil && fm.message.effectID != "" {
		err = writer.WriteField("message_effect_id", fm.message.effectID)
	}
	if err == nil && fm.chat.replyParameters != nil {
		body, err1 := json.Marshal(fm.chat.replyParameters)
		if err1 == nil {
			err = writer.WriteField("reply_parameters", string(body))
		}
	}
	if err == nil && fm.message.replyMarkup != nil {
		body, err1 := json.Marshal(fm.message.replyMarkup)
		if err1 == nil {
			err = writer.WriteField("reply_markup", string(body))
		}
	}
	return err
}

func (fm *Formatter) PrepareMedia(buf *bytes.Buffer) (string, error) {
	var tgMethod string
	file := new(os.File)
	writer := new(multipart.Writer)
	media, f, err := fm.checkTypeOfMedia()
	if err == nil {
		file, err = os.Open(media)
	}
	if err == nil {
		fm.contenttype, tgMethod, err = f(file, writer, media)
	}
	if err == nil {
		err = fm.commonMediaFields(file, writer)
	}
	return tgMethod, err
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
